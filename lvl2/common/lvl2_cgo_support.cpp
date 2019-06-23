#include "lvl2_cgo_support.h"
#include <vector>
#include <string>
#include <unordered_map>
#include <cstdlib>

using namespace std;




OPAQUE_STRING_LIST get_OPAQUE_STRING_LIST(size_t size) {
    auto list = new vector<string>(size);
    return static_cast<OPAQUE_STRING_LIST>(list);
}

void append_to_OPAQUE_STRING_LIST(OPAQUE_STRING_LIST obj, char *val) {
    auto list = static_cast<vector<string>*>(obj);
    auto str = string(val);
    free(val);
    list->push_back(str);
}

OPAQUE_OID_LIST get_OPAQUE_OID_LIST(size_t size) {
    auto list = new vector<unsigned long>(size);
    return static_cast<OPAQUE_OID_LIST>(list);
}

void append_to_OPAQUE_OID_LIST(OPAQUE_OID_LIST obj, unsigned long oid) {
    auto list = static_cast<vector<unsigned long>*>(obj);
    list->push_back(oid);
}

OPAQUE_UINT2UINT_MAP get_OPAQUE_UINT2UINT_MAP() {
    auto map = new unordered_map<unsigned int, unsigned int>();
    return static_cast<OPAQUE_UINT2UINT_MAP>(map);
}

void insert_into_UINT2UINT_MAP(OPAQUE_UINT2UINT_MAP obj, unsigned int age, unsigned int count) {
    auto map = static_cast<unordered_map<unsigned int, unsigned int>*>(obj);
    map->insert({age,count});
}
