package main

// #cgo CXXFLAGS: -std=c++17
// #cgo LDFLAGS: -lstdc++
// #include "/Users/imhiro/AllFiles/103_prog/europython_2019/lvl2/go/src/com.stbaer/demo_cgo/lvl2_cgo_support.h"
import "C"
import (
	"com.stbaer/demo_go/main_lib_simple"
	"unsafe"
)
import "com.stbaer/demo_cgo/cgo_lib_simple"

//export cgo_Add
func cgo_Add(cgo_v1, cgo_v2 C.int) C.int {
	var v1 int = int(cgo_v1)
	var v2 int = int(cgo_v2)
	var result int = main_lib_simple.Add(v1,v2)
	var cgo_result C.int = C.int(result)
	return cgo_result
}

//export cgo_GetMostBeFriendedReport
func cgo_GetMostBeFriendedReport(cgo_oids C.OPAQUE_OID_LIST) *C.char {
	var elements uint = uint(C.get_size_OPAQUE_OID_LIST(cgo_oids))
	var persons []*main_lib_simple.Person
	for i := 0; i < int(elements); i++ {
		var oid uint64 = uint64(C.get_at_OPAQUE_OID_LIST(cgo_oids,C.uint(i)))
		var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
		persons = append(persons,person)
	}

	result := main_lib_simple.GetMostBeFriendedReport(persons)
	var cgo_result *C.char = C.CString(result) // allocates, needs free() by caller
	return cgo_result
}


//export cgo_NewPerson
func cgo_NewPerson(cgo_firstname *C.char, cgo_lastname *C.char, cgo_age C.uint) C.ulong {
	// Convert Parameters
	firstname := C.GoString(cgo_firstname)
	lastname := C.GoString(cgo_lastname)
	age := uint(cgo_age)

	// Call Go Functions
	person := main_lib_simple.NewPerson(firstname, lastname, age)

	//Convert results, here, Golang Pointer to Reference
	var oid uint64 = cgo_lib_simple.PersonIdForObject(person)
	var cgo_oid C.ulong = C.ulong(oid)
	return cgo_oid
}

//export cgo_DeletePerson
func cgo_DeletePerson(cgo_oid C.ulong) {
	var oid uint64 = uint64(cgo_oid)
	cgo_lib_simple.RemovePersonObjectId(oid)
}

//export cgo_Person_Firstname
func cgo_Person_Firstname(cgo_oid C.ulong) *C.char {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var result string = person.Firstname()
	var cgo_result *C.char = C.CString(result) // allocates, needs free() by caller
	return cgo_result
}

//export cgo_Person_Lastname
func cgo_Person_Lastname(cgo_oid C.ulong) *C.char {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var result string = person.Lastname()
	var cgo_result *C.char = C.CString(result) // allocates, needs free() by caller
	return cgo_result
}

//export cgo_Person_String
func cgo_Person_String(cgo_oid C.ulong) *C.char {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var result string = person.String()
	var cgo_result *C.char = C.CString(result) // allocates, needs free() by caller
	return cgo_result
}

//export cgo_Person_Age
func cgo_Person_Age(cgo_oid C.ulong) C.uint {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var result uint = person.Age()
	var cgo_result C.uint = C.uint(result)
	return cgo_result
}

// New Functions to Handle Friends

//export cgo_Person_AddFriend
func cgo_Person_AddFriend(cgo_oid C.ulong, cgo_oid_friend C.ulong, cgo_error **C.char) C.uint {

	if *cgo_error != nil {
		panic("Outout Error Parameter Destination must point to Null")
	}

	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)

	var oid_friend uint64 = uint64(cgo_oid_friend)
	var friend *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid_friend)

	friend_count, err := person.AddFriend(friend)

	if (err != nil) {
		err_msg := err.Error()
		var cgo_err_msg *C.char = C.CString(err_msg) // allocates, needs free() by caller
		*cgo_error = cgo_err_msg;
	}

	var cgo_result C.uint = C.uint(friend_count)

	return cgo_result
}

//export cgo_Person_GetFriends
func cgo_Person_GetFriends(cgo_oid C.ulong) C.OPAQUE_OID_LIST {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var friends []*main_lib_simple.Person = person.GetFriends()


	var cgo_result C.OPAQUE_OID_LIST = C.get_OPAQUE_OID_LIST()

	for _, friend := range friends {
		var oid_friend uint64 = cgo_lib_simple.PersonIdForObject(friend)
		var cgo_oid_friend C.ulong = C.ulong(oid_friend)
		C.append_to_OPAQUE_OID_LIST(cgo_result, cgo_oid_friend)
	}
	return cgo_result
}

//export cgo_Person_GetFriendFirstNames
func cgo_Person_GetFriendFirstNames(cgo_oid C.ulong) C.OPAQUE_STRING_LIST {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)

	var first_names []string = person.GetFriendFirstNames()


	var cgo_result C.OPAQUE_STRING_LIST = C.get_OPAQUE_STRING_LIST()

	for _, first_name := range first_names {
		var cgo_first_name *C.char = C.CString(first_name)         // allocates, needs free() by caller
		C.append_to_OPAQUE_STRING_LIST(cgo_result, cgo_first_name) // takes ownership of *char
	}

	return cgo_result
}

//export cgo_Person_GetFriendCountByAge
func cgo_Person_GetFriendCountByAge(cgo_oid C.ulong) C.OPAQUE_UINT2UINT_MAP {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var cgo_result C.OPAQUE_UINT2UINT_MAP = C.get_OPAQUE_UINT2UINT_MAP()

	for count, age := range person.GetFriendCountByAge() {
		var cgo_count C.uint = C.uint(count)
		var cgo_age C.uint = C.uint(age)
		C.insert_into_UINT2UINT_MAP(cgo_result, cgo_count, cgo_age)
	}
	return cgo_result
}

//LVL3
//export cgo_Person_GetFriendsFiltered
func cgo_Person_GetFriendsFiltered(cgo_oid C.ulong,cgo_fun C.CALLBACK_OID_PERSON) C.OPAQUE_OID_LIST {

	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)

	fun := func(friend *main_lib_simple.Person) bool {
		var oid_friend uint64 = cgo_lib_simple.PersonIdForObject(friend)
		return C.call_CALLBACK_OID_PERSON(cgo_fun,C.ulong(oid_friend)) != 0
	}
	var friends []*main_lib_simple.Person = person.GetFriendsFiltered(fun)
	return build_OPAQUE_OID_LIST(friends)
}

//LVL3
//export cgo_Person_GetFriendsFilteredByAge
func cgo_Person_GetFriendsFilteredByAge(cgo_oid C.ulong, cgo_fun C.CALLBACK_AGE) C.OPAQUE_OID_LIST {

	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)

	fun := func(friend_age uint) bool {
		return C.call_CALLBACK_AGE(cgo_fun,C.uint(friend_age)) != 0
	}
	var friends []*main_lib_simple.Person = person.GetFriendsFilteredByAge(fun)


	return build_OPAQUE_OID_LIST(friends)

}

func build_OPAQUE_OID_LIST(persons []*main_lib_simple.Person)  C.OPAQUE_OID_LIST {
	var cgo_result C.OPAQUE_OID_LIST = C.get_OPAQUE_OID_LIST()

	for _, friend := range persons {
		var oid_friend uint64 = cgo_lib_simple.PersonIdForObject(friend)
		var cgo_oid_friend C.ulong = C.ulong(oid_friend)
		C.append_to_OPAQUE_OID_LIST(cgo_result, cgo_oid_friend)
	}
	return cgo_result
}

//LVL4
//export cgo_Person_GetFriendsFilteredByAge_2
func cgo_Person_GetFriendsFilteredByAge_2(cgo_oid C.ulong, cgo_fun C.CALLBACK_AGE_2, pointer unsafe.Pointer) C.OPAQUE_OID_LIST {

	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)

	fun := func(friend_age uint,data interface{}) bool {
		p := data.(unsafe.Pointer)
		return C.call_CALLBACK_AGE_2(cgo_fun,C.uint(friend_age),p) != 0
	}
	var friends []*main_lib_simple.Person = person.GetFriendsFilteredByAge_2(fun,pointer)


	return build_OPAQUE_OID_LIST(friends)

}


func main() {}
