package cgo_lib_simple

import (
	"com.stbaer/demo_go/main_lib_simple"
	"log"
	"sync"
)

var objectStorePersonInstance = NewObjectStorePerson()

func PersonIdForObject(object *main_lib_simple.Person) uint64 {
	return objectStorePersonInstance.IdForObject(object)
}

func ObjectForPersonId(oid uint64) *main_lib_simple.Person {
	return objectStorePersonInstance.ObjectForId(oid)
}

func RemovePersonObjectId(id uint64) {
	objectStorePersonInstance.RemoveObject(id)
}

type ObjectStorePerson struct {
	obj2int map[*main_lib_simple.Person]uint64
	int2obj map[uint64]*main_lib_simple.Person
	counter uint64
	mutex   *sync.Mutex
}

func NewObjectStorePerson() *ObjectStorePerson {
	return &ObjectStorePerson{
		obj2int: make(map[*main_lib_simple.Person]uint64),
		int2obj: make(map[uint64]*main_lib_simple.Person),
		mutex:   &sync.Mutex{},
	}
}

func (p *ObjectStorePerson) NewId() uint64 {
	p.counter += 1
	return p.counter
}

func (p *ObjectStorePerson) RemoveObject(id uint64) {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if obj, id_prs := p.int2obj[id]; id_prs {

		delete(p.obj2int, obj)
		delete(p.int2obj, id)

	} else {
		log.Fatalf("uint64 in %T was not present %v", p, id)
	}
}

func (p *ObjectStorePerson) IdForObject(obj *main_lib_simple.Person) uint64 {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if int_id, prs := p.obj2int[obj]; prs {
		return int_id
	}

	new_id := p.NewId()
	p.obj2int[obj] = new_id
	p.int2obj[new_id] = obj

	return new_id

}

func (p *ObjectStorePerson) ObjectForId(oid uint64) *main_lib_simple.Person {
	p.mutex.Lock()
	defer p.mutex.Unlock()
	if obj, prs := p.int2obj[oid]; prs {
		return obj
	} else {
		log.Fatalf("Get Object in %T was not present %v", p, oid)
	}
	return nil
}
