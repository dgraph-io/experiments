# Motivation

We have been using GoRocksDB. Some disadvantages are:

* GoRocksDB is not always compatible with the latest version of RocksDB.
* GoRocksDB uses the C API of RocksDB. To get a new function, you may have to modify both the C API and GoRocksDB.
* I believe there is some unnecessary string copying with Get calls in the C API. See [here](https://github.com/facebook/rocksdb/blob/master/db/c.cc#L722).
Notice the C++ code writes to a temporary string `std::string tmp`. Then we call `CopyString` which does a `malloc` and `memcpy`. Then GoRocksDB receives it as a `[]byte`. And I suppose Go's GC will take care of it.

# Difficulties

We wish that SWIG can do a better job, but it is a difficult tool to master. The main problem I encountered is `DB**`. `DB::Open` calls takes in a `DB**` argument and `DB` itself is an abstract class. SWIG has difficulty interpreting it.

We have tried different ways to fix this. One way is to use a '''typemap'''.

```
%typemap(in) rocksdb::DB ** (rocksdb::DB *temp) {
  $1 = &temp;
}

%typemap(argout) rocksdb::DB ** {
  %set_output(SWIG_NewPointerObj(SWIG_as_voidptr(*$1), $*1_descriptor, SWIG_POINTER_OWN));
}
```

This doesn't seem to work. This creates some DB ptr swig object which is undefined because DB is an abstract class.

To get around the abstract class problem, we tried using directors. But that seems to require that `DBImpl` is included as well. But the latter is an internal RocksDB object not found in the installed includes. I tried including more and more of these internal headers until I encountered some problem with the file `port/port.h`. Not sure why. For future reference, we tried including the following:

```cpp
#include "/home/jchiu/rocksdb-4.6.1/port/port.h"
#include "/home/jchiu/rocksdb-4.6.1/util/coding.h"
#include "/home/jchiu/rocksdb-4.6.1/db/dbformat.h"
#include "/home/jchiu/rocksdb-4.6.1/db/memtable_list.h"
#include "/home/jchiu/rocksdb-4.6.1/db/column_family.h"
#include "/home/jchiu/rocksdb-4.6.1/db/db_impl.h"
```

We have also tried playing around with `cpointers.i` and `pointer_functions`. But I failed to massage it to work.

In the end, I decided to just add some functions for `DB::Open` and `DB::Get`. You can see these in `extra.h`. I am not sure if they work completely fine (I need a key that is in the database) and I am not sure if there is some unnecessary string copying due to SWIG.

# How everything else works

We do not use `.swigcxx` as we want more control over the building and linking. See `build.sh` first. We run `swig` with a custom include folder. It generates C++ files in a subdirectory and a Go interface file in the current directory.

We then run a script `add_cgo_flays.py` to insert a `#cgo LDFLAGS` line into the Go file. One alternative is to set `CGO_LDFLAGS` in shell before you `go build` but this is not nice because any binary that uses that Go lib will need to do the same. The best way is to insert these flags right into the Go file. However, we couldn't find a way to do this in Swig. Most of the code insertion mechanisms in SWIG pertains to the `.cxx` file, not the Go file.
