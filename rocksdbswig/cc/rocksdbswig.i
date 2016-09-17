%module(directors="1") rocksdbswig
//%module rocksdbswig

%{
#include <string>
#include <rocksdb/slice.h>
#include <rocksdb/status.h>
#include <rocksdb/options.h>
#include <rocksdb/db.h>

#include "extra.h"


using namespace rocksdb;
%}


// %rename only works for function names. We use a C++ macro to rename "range"
// as myrange so that the Go code can compile.
#define range myrange

//%feature("director");

%include <typemaps.i>

// Disable mapping between C++ and Go strings to prevent extra copying.
//%include "std_string.i"
// %include "cpointer.i"

/*%typemap(in) rocksdb::DB ** (rocksdb::DB *temp) {
  $1 = &temp;
}

%typemap(argout) rocksdb::DB ** {
  %set_output(SWIG_NewPointerObj(SWIG_as_voidptr(*$1), $*1_descriptor, SWIG_POINTER_OWN));
}*/

%apply DB **OUTPUT { DB **dbptr };

%include "/usr/include/c++/5.4.0/string"

%include <rocksdb/slice.h>
%include <rocksdb/status.h>
%include <rocksdb/options.h>
%include <rocksdb/db.h>

%include "extra.h"

//%pointer_functions(DB, DB_p)