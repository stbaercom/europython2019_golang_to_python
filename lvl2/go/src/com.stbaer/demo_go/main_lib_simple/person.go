package main_lib_simple

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

func Add(v, v2 int) int {
	return v + v2
}

func GetMostBeFriendedReport(persons []*Person) string {

	type stringCount struct {
		s string
		c uint
	}

	counter := map[string]uint{}
	for _, person := range persons {
		for _, friend := range person.GetFriends() {
			person_text := friend.String()
			counter[person_text] += 1
		}
	}

	scs := []stringCount{}

	for k, v := range counter {
		scs = append(scs, stringCount{k, v})
	}

	sort.Slice(scs, func(i, j int) bool { return scs[i].c > scs[j].c })

	var sb strings.Builder

	for _, sc := range scs[:10] {
		sb.WriteString(fmt.Sprintf("%v : %v\n ", sc.c, sc.s))
	}
	sb.WriteString("\n")
	return sb.String()

}

type MyStrings []string

func (self MyStrings) Contains(search_str string) bool {
	for _, s := range self {
		if strings.Contains(s, search_str) {
			return true
		}
	}
	return false
}

type Person struct {
	firstname, lastname string
	age                 uint
	friends             []*Person
}

func NewPerson(firstname string, lastname string, age uint) *Person {
	return &Person{firstname: firstname, lastname: lastname, age: age}
}

func (p *Person) String() string {
	return fmt.Sprintf("%v %v (%v)", p.firstname, p.lastname, p.age)
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

func (p *Person) AddFriend(friend *Person) (uint, error) {

	if friend == p {
		return uint(len(p.friends)), errors.New("You cannot be your own friend")
	}

	if friend == nil {
		return uint(len(p.friends)), errors.New("Friends must not be Nil")
	}

	for _, old_friend := range p.friends {
		if old_friend == friend {
			return uint(len(p.friends)), errors.New("Cannot have the same friend twice")
		}
	}

	p.friends = append(p.friends, friend)
	return uint(len(p.friends)), nil

}


func (p *Person) GetFriends() []*Person {
	result := append([]*Person(nil), p.friends...)
	return result
}


func (p *Person) GetFriendFirstNames() []string {
	result := []string{}
	seen := make(map[string]bool)

	for _, friend := range p.friends {
		first_name := friend.firstname
		if !seen[first_name] {
			result = append(result, first_name)
		}
	}
	return result
}


func (p *Person) GetFriendCountByAge() map[uint]uint {

	result := make(map[uint]uint)

	for _, friend := range p.friends {
		age := friend.age
		result[age] += 1
	}

	return result

}


//LVL3
func (p *Person) GetFriendsFiltered(fun func(*Person) bool) []*Person {
	result := []*Person{}
	for _, friend := range p.friends {
		if fun(friend) {
			result = append(result, friend)
		}
	}
	return result
}

//LVL3
func (p *Person) GetFriendsFilteredByAge(fun func(uint) bool) []*Person {
	result := []*Person{}
	for _, friend := range p.friends {
		if fun(friend.age) {
			result = append(result, friend)
		}
	}
	return result
}

//LVL4
func (p *Person) GetFriendsFilteredByAge_2(fun func(uint, interface{}) bool, data interface{}) []*Person {
	result := []*Person{}
	for _, friend := range p.friends {
		if fun(friend.age, data) {
			result = append(result, friend)
		}
	}
	return result
}
