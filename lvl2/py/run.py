from itertools import count

import cgo_cython as cgo_mod
import random
import collections
import person as py_mod



def get_persons(n_pers, n_friends, mod=cgo_mod):
    names = "joe jack jane anne anna aug paul paula raoul eve eva erwin hans hanna daniel dan don donna".split()
    all_persons = []

    for _ in range(n_pers):
        first_name, *last_names = random.sample(names,3)
        age = random.randint(10,60)
        p = mod.Person(first_name, "-".join(last_names), age)
        all_persons.append(p)


    for person in all_persons:
        rnd_friends= max(0,int(random.normalvariate(n_friends, n_friends / 3.0)))
        friends = random.sample(all_persons,rnd_friends)
        for friend in friends:
            if friend is not person:
                person.add_friend(friend)

    return all_persons


def get_stefan():
    p1 = cgo_mod.Person("Jake", "Ant", 41)
    p2 = cgo_mod.Person("Joe", "Bee", 12)
    p3 = cgo_mod.Person("Jane", "Cat", 11)
    p4 = cgo_mod.Person("Jill", "Dog", 12)
    p1.add_friend(p2)
    p1.add_friend(p3)
    p1.add_friend(p4)
    f = p1.get_friends_filter_by_age(lambda x: x > 10 and x % 2 == 1)[0]
    f.string()

    return p1


def run_stefan():
    p1 = get_stefan()

    def print_persons(ps):
        for i, p in enumerate(ps):
            print(i, str(p))

    print("\nPrinting Friend Objects")
    print_persons(p1.get_friends())
    print("\nPrinting Friend Names")
    for i, s in enumerate(p1.get_friends_first_names()):
        print("Name %s %s" % (i, s))
    print("\nPrint Friend Age count")
    for age, count in p1.get_friends_count_by_age().items():
        print("Age %s : Count %s" % (age, count))

    def filter_ente_12(person):
        return (person.lastname() == "Ente") and (person.age() == 12)

    print("\nFriends Filtered")
    print_persons(p1.get_friends_filtered(filter_ente_12))

    def filter_12(age):
        if age == 12:
            return True
        return False

    print("\nFriends Filtered by Age")
    print_persons(p1.get_friends_filter_by_age(filter_12))
    print("\nFriends Filtered by Age 2")
    print_persons(p1.get_friends_filter_by_age_2(filter_12))

def run_persons_py(n):
    persons = get_persons(n,50,mod=py_mod)
    print_most_befriended_str(persons)
    print_most_befriended_getter(persons)

def run_persons(n):
    persons = get_persons(n,50)
    print_most_befriended_str(persons)
    print_most_befriended_getter(persons)
    print_most_befriended_go(persons)


def print_most_befriended_go(persons, mod=cgo_mod):
    most_common_report = mod.get_most_befriended_report(persons)
    print(most_common_report)

def print_most_befriended_str(persons):
    most_common_report = get_most_befriended_report(persons, use_str= True)
    print(most_common_report)

def print_most_befriended_getter(persons):
    most_common_report = get_most_befriended_report(persons, use_str = False)
    print(most_common_report)



def get_most_befriended_report(persons, use_str):
    friend_count = collections.Counter()
    for person in persons:
        for friend in person.get_friends():
            if use_str:
                person_txt = friend.string()
            else:
                person_txt = "%s %s, age %s " % (friend.firstname(),
                                                 friend.lastname(),
                                                 friend.age())
            friend_count[person_txt] +=1
    most_common = friend_count.most_common(10)
    return "\n".join("%5s : %s" % (v, k) for k, v in most_common) + "\n"



#run_stefan()
#run_persons(5000)
run_persons_py(5000)