//
// Created by Stefan Baerisch on 2019-06-28.
//

#ifndef CGO_DEMO_2_PERSON_H
#define CGO_DEMO_2_PERSON_H

class Person;

#include <vector>
#include <string>
#include <memory>

using sperson = std::shared_ptr<Person>;

class Person {
    std::string first_name;
    std::string last_name;
    unsigned int age;
    std::vector<sperson> friends;
public:
    Person(const std::string firstName, const std::string lastName, unsigned int age);

    const std::string getFirstName() const;

    const std::string getLastName() const;

    const std::string getString() const;

    unsigned int getAge() const;

    const std::vector<sperson> getFriends() const;

    unsigned int addFriend(sperson);
};

#endif //CGO_DEMO_2_PERSON_H
