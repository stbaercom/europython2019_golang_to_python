#coding: utf-8
#cython: language_level=3


from libc.stdlib cimport free


cdef extern from "cgo_lib/main.h":
    long unsigned int cgo_NewPerson(char* p0, char* p1, unsigned int p2)
    void cgo_DeletePerson(long unsigned int p0)
    char* cgo_Person_Firstname(long unsigned int p0)
    char* cgo_Person_Lastname(long unsigned int p0)
    char* cgo_Person_String(long unsigned int p0)
    unsigned int cgo_Person_Age(long unsigned int p0)


cdef class Person:
    cdef unsigned long oid

    def __cinit__(self, str firstname, str lastname, unsigned int age):
        py_byte_string = firstname.encode("UTF-8")
        cdef char* cfirstname = py_byte_string

        py_byte_string = lastname.encode("UTF-8")
        cdef char* clastname  = py_byte_string

        self.oid = cgo_NewPerson(cfirstname,clastname,age)


    def __dealloc__(self):
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

    cpdef unsigned int age(self):
        return cgo_Person_Age(self.oid)

cdef str tounicode_with_free(char* s):
    try:
        return s.decode('UTF-8', 'strict')
    finally:
        free(s)