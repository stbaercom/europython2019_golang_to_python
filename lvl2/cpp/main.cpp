


#include "cgo_lib/cgo_lib_lvl2.h"
#include <iostream>

int main() {
    char first_name[] = "Stefan";
    char last_name[] = "Baerisch";
    unsigned int age = 42;
    long unsigned int oid = cgo_NewPerson(first_name, last_name, age);
    char *person_string = cgo_Person_String(oid);
    std::cout << "Person Name " << person_string;
    cgo_DeletePerson(oid);
}