@0xd44fb1ce6b5a1003;
using Go = import "go.capnp";
$Go.package("plist");
$Go.import("testpkg");

struct PostingList {
	ids @0: List(UInt64);
	title @1: Text;
}
