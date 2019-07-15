from itertools import count
import random
import collections
import time

import cgo_cython as cgo_mod
import person as py_mod

def get_persons(n_pers, n_friends, mod, seed):
    names = "joe jack jane anne anna aug paul paula raoul eve eva erwin hans hanna daniel dan don donna".split()
    all_persons = []

    if seed is not None:
        random.seed(seed)


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


def main(number_of_persons,number_of_friends,seed):
    py_person = get_persons(number_of_persons,number_of_friends,py_mod,seed)
    cgo_persons = get_persons(number_of_persons,number_of_friends,cgo_mod,seed)
    mod_options = [(py_person,"Python"),(cgo_persons,"CGO")]
    use_str_options = [True,False]

    for persons in mod_options:
        for use_str in use_str_options:
            print(f"Persons: {persons[1]} UseString: {use_str}")
            s = time.perf_counter()
            most_common_report = get_most_befriended_report(persons[0], use_str = use_str)
            t = time.perf_counter() - s
            print(f"Time used {t}")
            print(most_common_report)

    print(f"Golang Implementation")
    s = time.perf_counter()
    most_common_report = cgo_mod.get_most_befriended_report(cgo_persons)
    t = time.perf_counter() - s
    print(f"Time used {t}")
    print(most_common_report)

if __name__ == '__main__':
    number_of_persons = 5000
    number_of_friends = 50
    rnd_seed = 42
    main(number_of_persons, number_of_friends, rnd_seed)