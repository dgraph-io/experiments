namespace rocksdb {

typedef struct {
	Status status;
	DB* db;
} StatusDBPair;

StatusDBPair MyDBOpen(const Options& options, const std::string& name) {
	StatusDBPair out;
	out.status = DB::Open(options, name, &out.db);
	return out;
}

typedef struct {
	Status status;
	std::string value;
} StatusStringPair;

StatusStringPair MyDBGet(DB* db, const ReadOptions& options, const Slice& key) {
	StatusStringPair out;
	out.status = db->Get(options, key, &out.value);
	return out;
}

}  // namespace rocksdb