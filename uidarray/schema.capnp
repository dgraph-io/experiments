@0x802751b28ad31f4c;
using Go = import "go.capnp";
$Go.package("main");
$Go.import("testpkg");


struct UidArrayCapn { 
   uids  @0:   List(UInt64); 
} 

##compile with:

##
##
##   capnp compile -ogo ./schema.capnp

