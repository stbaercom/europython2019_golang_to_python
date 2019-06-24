import cgo_cython_lvl2 as m

p1 = m.Person("Stefan","Baerisch",41)
p2 = m.Person("Jochen", "Ente", 12)
p3 = m.Person("Claudia", "Ente", 11)
p4 = m.Person("Matthew", "Ente", 12)

p1.add_friend(p2)
p1.add_friend(p3)
p1.add_friend(p4)

print("Printing Friend Objects")
for i, p in enumerate(p1.get_friends()):
    print(i,str(p))

print("Printing Friend Names")
for i, s in enumerate(p1.get_friends_first_names()):
    print("Name %s %s" % (i,s))

print("Print Friend Agecout")
for age,count in p1.get_friends_count_by_age().items():
    print ("Age %s : Count %s" % (age,count))