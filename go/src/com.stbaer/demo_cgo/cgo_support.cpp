#include "cgo_support.h"
#include <vector>
#include <string>
#include <unordered_map>
#include <cstdlib>

using namespace std;


OPAQUE_STRING_LIST get_OPAQUE_STRING_LIST() {
    auto list = new vector<string>();
    return static_cast<OPAQUE_STRING_LIST>(list);
}

void append_to_OPAQUE_STRING_LIST(OPAQUE_STRING_LIST obj, char *val) {
    auto list = static_cast<vector<string>*>(obj);
    auto str = string(val);
    free(val);
    list->push_back(str);
}

OPAQUE_OID_LIST get_OPAQUE_OID_LIST() {
    auto list = new vector<unsigned long>();
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

//LVL3

int call_CALLBACK_OID_PERSON(CALLBACK_OID_PERSON cgo_fun, unsigned long oid_friend) {
    return cgo_fun(oid_friend);
}

int call_CALLBACK_AGE(CALLBACK_AGE cgo_fun,unsigned int age) {
    return cgo_fun(age);
}


//LVL4
int call_CALLBACK_AGE_2(CALLBACK_AGE_2 cgo_fun, unsigned int age, void *data) {
    return cgo_fun(age,data);
}

unsigned int get_size_OPAQUE_OID_LIST(OPAQUE_OID_LIST obj) {
    auto list = static_cast<vector<unsigned long>*>(obj);
    return list->size();
}

unsigned long get_at_OPAQUE_OID_LIST(OPAQUE_OID_LIST obj, unsigned int index) {
    auto list = static_cast<vector<unsigned long>*>(obj);
    return list->at(index);
}
