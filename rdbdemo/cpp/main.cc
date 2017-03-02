/*
RDB=$HOME/rocksdb-4.11.2
DIR=$HOME/rocksdb_demo/data

g++ main.cc $RDB/librocksdb.a -std=c++11 -O2 -lz -lbz2 -I$RDB/include

rm -Rf $DIR
mkdir -p $DIR
./a.out $DIR
*/

#include <rocksdb/db.h>
#include <rocksdb/options.h>
#include <rocksdb/slice.h>
#include <rocksdb/status.h>

#include <cassert>
#include <cstdio>
#include <cstring>
#include <iostream>
#include <memory>
#include <string>

using namespace rocksdb;  // Lazy here.

void put_demo(DB* db) {
	char key[20];
	char val[5000 * 5 + 20];
	memset(val, 'v', sizeof(val));
	const int key_len = 3 + 9;
	const int val_len = 5000 * 5 + 9;

	WriteOptions w_opt;
	w_opt.sync = false;

	for (int i = 0; i < 100000000; ++i) {
		if ((i % 100000) == 0) {
			std::cout << "Added " << i << " keys\n";
		}
		sprintf(key, "key%09x", i % 10000);
		sprintf(val + 5000 * 5, "%09x", i);
		Status status = db->Put(w_opt, Slice(key, key_len), Slice(val, val_len));
		assert(status.ok());
	}
}

// Doesn't seem to have any memory leak. Htop mem usage remains constant at 2.3%.
void writebatch_demo(DB* db) {
	char key[20];
	char val[5000 * 5 + 20];
	memset(val, 'v', sizeof(val));
	const int key_len = 3 + 9;
	const int val_len = 5000 * 5 + 9;

	WriteOptions w_opt;
	w_opt.sync = false;

	WriteBatch wb;

	for (int i = 0; i < 100000000; ++i) {
		if ((i % 100000) == 0) {
			std::cout << "Added " << i << " keys\n";
		}
		if ((i % 1000) == 0) {
			Status status = db->Write(w_opt, &wb);
			assert(status.ok());
			wb.Clear();
		}
		sprintf(key, "key%09x", i % 10000);
		sprintf(val + 5000 * 5, "%09x", i);

		wb.Put(Slice(key, key_len), Slice(val, val_len));
		//db->Put(w_opt, Slice(key, key_len), Slice(val, val_len));
	}
}

int main(int argc, char* argv[]) {
	assert(argc == 2);
	std::unique_ptr<DB> db;
	{
		DB* raw;
		Options opt;
		opt.create_if_missing = true;
		Status s = DB::Open(opt, argv[1], &raw);
		assert(s.ok());
		db.reset(raw);
	}
	// put_demo(db.get());
	writebatch_demo(db.get());
	return 0;
}
