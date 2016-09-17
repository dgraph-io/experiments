#ifndef __WSTRING__
#define __WSTRING__

#ifdef __cplusplus
extern "C" {
#endif

typedef struct wstring_t wstring_t;
wstring_t* wstring_new();
void wstring_destroy(wstring_t* s);
char* wstring_get(wstring_t* s, size_t* len);
void wstring_set(wstring_t* s, char* data, size_t len);
size_t wstring_len(wstring_t* s);

#ifdef __cplusplus
}  /* end extern "C" */
#endif

#endif  // __WSTRING__
