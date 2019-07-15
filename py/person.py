class Person:
    def __init__(self, firstname, lastname, age):
        self._firstname = firstname
        self._lastname = lastname
        self._age = age
        self._friends = []

    def firstname(self):
        return self._firstname

    def lastname(self):
        return self._lastname

    def age(self):
        return self._age

    def string(self):
        return "%s %s %s" % (self._firstname,self._lastname,self._age)

    def add_friend(self,friend):
        self._friends.append(friend)

    def get_friends(self):
        return list(self._friends)