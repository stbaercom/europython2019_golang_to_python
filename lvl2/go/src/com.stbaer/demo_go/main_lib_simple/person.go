package main_lib_simple

import (
	"errors"
	"fmt"
)

type Person struct {
	firstname, lastname string
	age                 uint
	friends []*Person
}

func (p *Person) AddFriend(friend *Person) (uint,error) {

	if friend == nil {
		return uint(len(p.friends)), errors.New("Friends must not be Nil")
	}

	for _, old_friend := range p.friends {
		if old_friend == friend {
			return uint(len(p.friends)), errors.New("Cannot have the same friend twice")
		}
	}

	p.friends = append(p.friends,friend)
	return uint(len(p.friends)),nil

}

func (p *Person) GetFriends() []*Person {
	result := append([]*Person(nil), p.friends...)
	return result
}

func (p *Person) GetFriendFirstNames() []string {
	result := []string{}
	seen := make(map[string]bool)

	for _,friend := range p.friends {
		first_name  := friend.firstname
		if !seen[first_name] {
			result = append(result,first_name)
		}
	}
	return result
}

func (p *Person) GetFriendCountByAge() map[uint]uint {

	result := make(map[uint]uint)

	for _,friend := range p.friends {
		age := friend.age
		result[age] +=1
	}

	return result

}



func (p *Person) String() string {
	return fmt.Sprintf("%v %v (%v)" , p.firstname, p.lastname, p.age)
}

func (p *Person) Age() uint {
	return p.age
}

func (p *Person) Firstname() string {
	return p.firstname
}

func (p *Person) Lastname() string {
	return p.lastname
}

func NewPerson(firstname string, lastname string, age uint) *Person {
	return &Person{firstname: firstname, lastname: lastname, age: age}
}

