package main


import "C"
import 	"com.stbaer/demo_go/main_lib_simple"
import "com.stbaer/demo_cgo/cgo_lib_simple"

//export cgo_NewPerson
func cgo_NewPerson(cgo_firstname *C.char, cgo_lastname *C.char, cgo_age C.uint) C.ulong {
	// Convert Parameters
	firstname := C.GoString(cgo_firstname)
	lastname := C.GoString(cgo_lastname)
	age := uint(cgo_age)

	// Call Go Functions
	person := main_lib_simple.NewPerson(firstname,lastname,age)

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
	var cgo_result *C.char = C.CString(result)
	return cgo_result
}

//export cgo_Person_Lastname
func cgo_Person_Lastname(cgo_oid C.ulong) *C.char {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var result string = person.Lastname()
	var cgo_result *C.char = C.CString(result)
	return cgo_result
}

//export cgo_Person_String
func cgo_Person_String(cgo_oid C.ulong) *C.char {
	var oid uint64 = uint64(cgo_oid)
	var person *main_lib_simple.Person = cgo_lib_simple.ObjectForPersonId(oid)
	var result string = person.String()
	var cgo_result *C.char = C.CString(result)
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


func main() {}


