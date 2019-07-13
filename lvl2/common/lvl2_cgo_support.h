


#ifdef __cplusplus
extern "C" {
#endif

typedef void* OPAQUE_STRING_LIST;
typedef void* OPAQUE_OID_LIST;
typedef void* OPAQUE_UINT2UINT_MAP;

OPAQUE_STRING_LIST get_OPAQUE_STRING_LIST();
void append_to_OPAQUE_STRING_LIST(OPAQUE_STRING_LIST obj,char * val);

OPAQUE_OID_LIST get_OPAQUE_OID_LIST();
void append_to_OPAQUE_OID_LIST(OPAQUE_OID_LIST obj, unsigned long oid);
unsigned int get_size_OPAQUE_OID_LIST(OPAQUE_OID_LIST obj);
unsigned long get_at_OPAQUE_OID_LIST(OPAQUE_OID_LIST obj, unsigned int index);

OPAQUE_UINT2UINT_MAP get_OPAQUE_UINT2UINT_MAP();
void insert_into_UINT2UINT_MAP(OPAQUE_UINT2UINT_MAP obj,unsigned int age,
    unsigned int count);

#ifdef __cplusplus
}
#endif


#ifdef __cplusplus
extern "C" {
#endif

//LVL3
typedef int (*CALLBACK_AGE)(unsigned int);
int call_CALLBACK_AGE(CALLBACK_AGE cgo_fun,unsigned int age);

typedef int (*CALLBACK_OID_PERSON)(unsigned long);
int call_CALLBACK_OID_PERSON(CALLBACK_OID_PERSON cgo_fun, unsigned long oid_friend);


//LVL4
typedef int (*CALLBACK_AGE_2)(unsigned int, void * data);
int call_CALLBACK_AGE_2(CALLBACK_AGE_2 cgo_fun, unsigned int age, void *data);

#ifdef __cplusplus
}
#endif