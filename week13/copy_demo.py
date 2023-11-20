import copy

a = [3, [5, 1], 17]
b = a
c = a[:]
d = list(a)
e = a.copy
f = copy.deepcopy(a)

print(a, b, c, d, e, f)
a[1][0] = 100
print(a, b, c, d, e, f)