#coding: utf-8
#cython: language_level=3


from libc.stdlib cimport free

ctypedef void* OPAQUE_OID_LIST
ctypedef void* OPAQUE_STRING_LIST
ctypedef void* OPAQUE_UINT2UINT_MAP

cdef extern from "cgo_lib/cgo_lib_lvl2.h":
    long unsigned int cgo_NewPerson(char* p0, char* p1, unsigned int p2)
    void cgo_DeletePerson(long unsigned int p0)
    char* cgo_Person_Firstname(long unsigned int p0)
    char* cgo_Person_Lastname(long unsigned int p0)
    char* cgo_Person_String(long unsigned int p0)
    unsigned int cgo_Person_Age(long unsigned int p0)
    unsigned int cgo_Person_AddFriend(long unsigned int p0, long unsigned int p1, char** p2)
    OPAQUE_OID_LIST cgo_Person_GetFriends(long unsigned int p0)
    OPAQUE_STRING_LIST cgo_Person_GetFriendFirstNames(long unsigned int p0)
    OPAQUE_UINT2UINT_MAP cgo_Person_GetFriendCountByAge(long unsigned int p0)


from libcpp.vector cimport vector
from libcpp.string cimport string
from libcpp.unordered_map cimport unordered_map


cdef class Person:
    KNOWN = {}


    @classmethod
    def get_person(cls,oid):
        print("Getting Person %s. In Map: %s" %(oid,oid in cls.KNOWN))
        if oid in cls.KNOWN:
            result = cls.KNOWN[oid]
        else:
            result = Person(oid)
        print("Returning Person %s" % str(result))
        return result



    cdef unsigned long oid

    def __cinit__(self, str firstname = None , str lastname = None, unsigned int age = 0,  long oid = -1):
        cdef char* cfirstname
        cdef char* clastname

        if (firstname is not None) and (lastname is not None) and (oid == -1):
            py_byte_string = firstname.encode("UTF-8")
            cfirstname = py_byte_string

            py_byte_string = lastname.encode("UTF-8")
            clastname  = py_byte_string

            self.oid = cgo_NewPerson(cfirstname,clastname,age)
            print("Alloc New (%s) %s" % (self.oid,self.string()))
        else:
            self.oid = oid
            print("Alloc Existing (%s) %s" % (self.oid,self.string()))

        Person.KNOWN[self.oid] = self


    def __dealloc__(self):
        print("Dealloc (%s) %s" % (self.oid,self.string()))
        cgo_DeletePerson(self.oid)

    cpdef str firstname(self):
        cdef char* cresult =  cgo_Person_Firstname(self.oid)
        cdef str result = tounicode_with_free(cresult)
        return result

    cpdef str lastname(self):
        cdef char* cresult =  cgo_Person_Lastname(self.oid)
        cdef str result = tounicode_with_free(cresult)
        return result

    cpdef str string(self):
        cdef char* cresult =  cgo_Person_String(self.oid)
        cdef str result = tounicode_with_free(cresult)
        return result

    def __str__(self):
        return self.string()

    cpdef unsigned int age(self):
        return cgo_Person_Age(self.oid)

    def add_friend(self, Person friend):
        cdef unsigned long oid_friend = friend.oid
        cdef unsigned long oid = self.oid
        cdef char * error_char_pointer = NULL
        cdef char ** error_char_pointer_pointer = &error_char_pointer
        cdef str error_string
        cdef int friend_count = cgo_Person_AddFriend(oid,oid_friend,error_char_pointer_pointer)
        if error_char_pointer != NULL:
            error_string = tounicode_with_free(error_char_pointer_pointer[0])
            raise Exception(error_string)
        return friend_count

    cpdef list get_friends(self):
        cdef OPAQUE_OID_LIST opaque_p_id_vector = cgo_Person_GetFriends(self.oid)
        cdef vector[unsigned long]* p_id_vector = <vector[unsigned long]*> opaque_p_id_vector
        cdef list result = []
        for oid_friend in p_id_vector[0]:
            friend = Person.get_person(oid_friend)
            result.append(friend)
        del p_id_vector
        return result

    cpdef list get_friends_first_names(self):
        cdef OPAQUE_STRING_LIST opaque_p_string_vector =  cgo_Person_GetFriendFirstNames(self.oid)
        cdef vector[string]* p_id_vector = <vector[string]*> opaque_p_string_vector
        cdef list result = [n.decode("UTF-8") for n in p_id_vector[0]]
        del p_id_vector
        return result

    cpdef dict get_friends_count_by_age(self):
        cdef OPAQUE_UINT2UINT_MAP qpaque_p_map = cgo_Person_GetFriendCountByAge(self.oid)
        cdef unordered_map[unsigned int, unsigned int]* p_map = <unordered_map[unsigned int, unsigned int]*> qpaque_p_map
        cdef dict result = p_map[0]
        del p_map
        return result



cdef str tounicode_with_free(char* s):
    try:
        return s.decode('UTF-8', 'strict')
    finally:
        free(s)