#coding: utf-8
#cython: language_level=3

import weakref

from libc.stdlib cimport free



ctypedef void*OPAQUE_OID_LIST
ctypedef void*OPAQUE_STRING_LIST
ctypedef void*OPAQUE_UINT2UINT_MAP
ctypedef int (*CALLBACK_AGE)(unsigned int)
ctypedef int (*CALLBACK_OID_PERSON)(unsigned long)
ctypedef int (*CALLBACK_AGE_2)(unsigned int, void *data)

ctypedef unsigned long ulong

cdef extern from "cgo_lib/cgo_lib.h":
    int cgo_Add(int p0, int p1)

    #LVL1
cdef extern from "cgo_lib/cgo_lib.h":
    long unsigned int cgo_NewPerson(char*p0, char*p1, unsigned int p2)
    void cgo_DeletePerson(long unsigned int p0)
    char*cgo_Person_Firstname(long unsigned int p0)
    char*cgo_Person_Lastname(long unsigned int p0)
    char*cgo_Person_String(long unsigned int p0)
    unsigned int cgo_Person_Age(long unsigned int p0)

    #LVL2
cdef extern from "cgo_lib/cgo_lib.h":
    unsigned int cgo_Person_AddFriend(long unsigned int p0, long unsigned int p1, char** p2)
    OPAQUE_OID_LIST cgo_Person_GetFriends(long unsigned int p0)
    OPAQUE_STRING_LIST cgo_Person_GetFriendFirstNames(long unsigned int p0)
    OPAQUE_UINT2UINT_MAP cgo_Person_GetFriendCountByAge(long unsigned int p0)

cdef extern from "cgo_lib/cgo_lib.h":
    OPAQUE_OID_LIST cgo_Person_GetFriendsFiltered(long unsigned int p0, CALLBACK_OID_PERSON p1)
    OPAQUE_OID_LIST cgo_Person_GetFriendsFilteredByAge(long unsigned int p0, CALLBACK_AGE p1)

    #LVL4
    OPAQUE_OID_LIST cgo_Person_GetFriendsFilteredByAge_2(long unsigned int p0, CALLBACK_AGE_2 p1, void*p2)

    #LVL5 - Benchmark
    char*cgo_GetMostBeFriendedReport(OPAQUE_OID_LIST p0)

from libcpp.vector cimport vector
from libcpp.string cimport string
from libcpp.unordered_map cimport unordered_map

cdef object _PersonFilterFun = None

cpdef int add(int v1, int v2):
    return cgo_Add(v1, v2)

cpdef str get_most_befriended_report(list persons):
    cdef vector[ulong]id_vector = vector[ulong]()
    for person in persons:
        id_vector.push_back(person._oid())

    cdef char*cresult = cgo_GetMostBeFriendedReport(<OPAQUE_OID_LIST> &id_vector)

    cdef str result = tounicode_with_free(cresult)
    return result

cdef class Person:
    _KNOWN = {}
    cdef unsigned long oid
    cdef object __weakref__

    def __cinit__(self, str firstname = None, str lastname = None,
                  unsigned int age = 0, long oid = -1):
        cdef char*cfirstname
        cdef char*clastname
        if (firstname is not None) and (lastname is not None) and (oid == -1):
            py_byte_string = firstname.encode("UTF-8")
            cfirstname = py_byte_string
            py_byte_string = lastname.encode("UTF-8")
            clastname = py_byte_string
            self.oid = cgo_NewPerson(cfirstname, clastname, age)
        else:
            self.oid = oid
        Person._KNOWN[self.oid] = weakref.ref(self)

    def __dealloc__(self):
        cgo_DeletePerson(self.oid)

    cpdef ulong _oid(self):
        return self.oid

    cpdef str firstname(self):
        cdef char*cresult = cgo_Person_Firstname(self.oid)
        cdef str result = tounicode_with_free(cresult)
        return result

    cpdef str lastname(self):
        cdef char*cresult = cgo_Person_Lastname(self.oid)
        cdef str result = tounicode_with_free(cresult)
        return result

    cpdef str string(self):
        cdef char*cresult = cgo_Person_String(self.oid)
        cdef str result = tounicode_with_free(cresult)
        return result

    def __str__(self):
        return self.string()

    @classmethod
    def get_person(cls, oid):
        if oid in cls._KNOWN:
            person_wref = cls._KNOWN[oid]
            person = person_wref()
            if person is not None:
                cls._KNOWN[oid] = weakref.ref(person)
                return person
        result = Person(oid)
        return result

    cpdef unsigned int age(self):
        return cgo_Person_Age(self.oid)

    def add_friend(self, Person friend):
        cdef unsigned long oid_friend = friend.oid
        cdef unsigned long oid = self.oid
        cdef char *error_char_pointer = NULL
        cdef char ** error_char_pointer_pointer = &error_char_pointer
        cdef str error_string
        cdef int friend_count = cgo_Person_AddFriend(oid, oid_friend, error_char_pointer_pointer)
        if error_char_pointer != NULL:
            error_string = tounicode_with_free(error_char_pointer_pointer[0])
            raise Exception(error_string)
        return friend_count

    cpdef list get_friends(self):
        cdef OPAQUE_OID_LIST opaque_p_id_vector = cgo_Person_GetFriends(self.oid)
        cdef vector[unsigned long]*p_id_vector = <vector[unsigned long]*> opaque_p_id_vector
        cdef list result = []
        for oid_friend in p_id_vector[0]:
            friend = Person.get_person(oid_friend)
            result.append(friend)
        del p_id_vector
        return result

    cpdef list get_friends_first_names(self):
        cdef OPAQUE_STRING_LIST opaque_p_string_vector = cgo_Person_GetFriendFirstNames(self.oid)
        cdef vector[string]*p_id_vector = <vector[string]*> opaque_p_string_vector
        cdef list result = [n.decode("UTF-8") for n in p_id_vector[0]]
        del p_id_vector
        return result

    cpdef dict get_friends_count_by_age(self):
        cdef OPAQUE_UINT2UINT_MAP qpaque_p_map = cgo_Person_GetFriendCountByAge(self.oid)
        cdef unordered_map[unsigned int, unsigned int]*p_map = <unordered_map[unsigned int, unsigned int]*> qpaque_p_map
        cdef dict result = p_map[0]
        del p_map
        return result

    #LVL3

    @staticmethod
    cdef list _build_result_list(OPAQUE_OID_LIST opaque_p_id_vector):
        cdef vector[unsigned long]*p_id_vector = <vector[unsigned long]*> opaque_p_id_vector
        cdef list result = []
        for oid_friend in p_id_vector[0]:
            friend = Person.get_person(oid_friend)
            result.append(friend)
        del p_id_vector
        return result

    cpdef list get_friends_filtered(self, object fun):
        global _PersonFilterFun
        _PersonFilterFun = fun
        cdef OPAQUE_OID_LIST opaque_p_id_vector = cgo_Person_GetFriendsFiltered(self.oid, filter_friends)
        result = Person._build_result_list(opaque_p_id_vector)
        _PersonFilterFun = None
        return result

    cpdef list get_friends_filter_by_age(self, object fun):
        global _PersonFilterFun
        _PersonFilterFun = fun

        cdef OPAQUE_OID_LIST opaque_p_id_vector = \
            cgo_Person_GetFriendsFilteredByAge(self.oid, filter_by_age)

        result = Person._build_result_list(opaque_p_id_vector)

        _PersonFilterFun = None
        return result

    cpdef list get_friends_filter_by_age_2(self, object fun):
        cdef void*p_fun = <void*> fun

        cdef OPAQUE_OID_LIST opaque_p_id_vector = cgo_Person_GetFriendsFilteredByAge_2(self.oid, filter_by_age_2, p_fun)
        result = Person._build_result_list(opaque_p_id_vector)
        return result

cdef int filter_friends(unsigned long oid):
    fun = _PersonFilterFun
    person = Person.get_person(oid)
    return fun(person)

cdef int filter_by_age(unsigned int age):
    fun = _PersonFilterFun
    return fun(age)

cdef int filter_by_age_2(unsigned int age, void *data):
    fun = <object> data
    return fun(age)

cdef str tounicode_with_free(char*s):
    try:
        return s.decode('UTF-8', 'strict')
    finally:
        free(s)
