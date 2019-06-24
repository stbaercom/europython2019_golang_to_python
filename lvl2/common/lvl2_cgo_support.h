


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


OPAQUE_UINT2UINT_MAP get_OPAQUE_UINT2UINT_MAP();
void insert_into_UINT2UINT_MAP(OPAQUE_UINT2UINT_MAP obj,unsigned int age, unsigned int count);


#ifdef __cplusplus
}
#endif