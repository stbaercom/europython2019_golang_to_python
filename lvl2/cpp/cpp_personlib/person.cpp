//
// Created by Stefan Baerisch on 2019-06-28.
//

#include "person.h"

#include <algorithm>
#include <stdexcept>

using namespace std;



Person::Person(const string &firstName, const string &lastName, unsigned int age) : first_name(firstName),
                                                                                    last_name(lastName), age(age) {}

const string &Person::getFirstName() const {
    return first_name;
}

const string &Person::getLastName() const {
    return last_name;
}

unsigned int Person::getAge() const {
    return age;
}

const vector<sperson> &Person::getFriends() const {
    vector<sperson> result;
    copy(friends.cbegin(),friends.cend(),result.begin());
    return result;
}

unsigned int Person::addFriend(sperson p) {


    if(!p) {
        throw invalid_argument("Friends must not be Nil");
    }

    if(p.get()==this) {
        throw invalid_argument("You cannot be your own friend");
    }

    for (const auto& f : friends) {
        if (f == p) {
            throw invalid_argument("")
        }
    }



    friends.push_back(p);
}
