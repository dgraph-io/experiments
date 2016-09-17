#include <cstdio>
#include <cstdint>
#include <string>

#include "wstring.h"

struct wstring_t { std::string rep; };

wstring_t* wstring_new() {
	printf("~~Creating wstring\n");
	return new wstring_t;
}

void wstring_destroy(wstring_t* s) {
	printf("~~Destroying wstring\n");
	delete s;
}

char* wstring_get(wstring_t* s, size_t* len) {
	*len = s->rep.length();
	return (char*)(s->rep.data());
}

void wstring_set(wstring_t* s, char* data, size_t len) {
	s->rep.assign(data, len);
}

size_t wstring_len(wstring_t* s) {
	return s->rep.length();
}