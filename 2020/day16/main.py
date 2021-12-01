import re
import math

p = re.compile(r'(\w+ ?\w+?): (\d+)-(\d+) or (\d+)-(\d+)')
n = re.compile(r'(\d+)')
t = re.compile(r'your ticket:\n(.*)\n')


def validate(row, valid_ranges):
    invalid = []
    for i in n.findall(row):
        valid = False
        for r in valid_ranges.values():
            i = int(i) 
            a, b = r
            if i in range(a[0], a[1]) or i in range(b[0], b[1]):
                valid = True
                break
        if not valid:
            invalid.append(i)
    return invalid


def get_ranges(input_raw) -> list:
    r = {}
    for x in p.findall(input_raw):
        k, a, b, c, d = x
        r[k] = (
            (int(a), int(b)+1),
            (int(c), int(d)+1)
        )
    return r

def one(input_raw):
    ranges = get_ranges(input_raw)
    nearby = input_raw.split("nearby tickets:\n")[1]
    invalid = []
    for row in nearby.split("\n"):
        invalid.extend(validate(row, ranges))
    return sum(invalid)

def classify(row: str, valid_ranges: dict):
    def get_keys(n, valid_ranges):
        keys = []
        valid = False
        for k, r in valid_ranges.items():
            i = int(n) 
            a, b = r
            if i in range(a[0], a[1]) or i in range(b[0], b[1]):
                keys.append(k)
        return keys

    positions = {}
    for i, num in enumerate(n.findall(row)):
        k = get_keys(num, valid_ranges)
        positions[i] = k
    return positions

def two(input_raw):
    ranges = get_ranges(input_raw)
    nearby = input_raw.split("nearby tickets:\n")[1]
    valid = []
    for row in nearby.split("\n"):
        if not validate(row, ranges):
            valid.append(row)

    sets = {}
    for row in valid:
        for position, classes in classify(row, ranges).items():
            if position not in sets:
                sets[position] = set(classes)
            else:
                sets[position] = sets[position].intersection(set(classes))


    def remove_known(sets, removed_positions=None):
        removed_positions = removed_positions if removed_positions else {}
        new_sets = {}
        
        for k, v in sets.items():
            if len(v) == 1:
                v = v.pop()
                removed_positions[v] = k
            else:
                new_sets[k] = v

        return removed_positions, new_sets

    def remove_by_value(sets, val):
        new_sets = {}
        for k, v in sets.items():
            try:
                v.remove(val)
                new_sets[k] = v
            except KeyError:
                new_sets[k] = v
        return new_sets

    removed = {}
    l = len(sets)
    while len(removed) < l:
        r, sets = remove_known(sets)
        removed.update(r)

        for k in removed:
            sets = remove_by_value(sets, k)

    positions = [int(x) for x in t.findall(input_raw)[0].split(',')]
    values = [positions[v] for k, v in removed.items() if k.startswith('departure')]

    return math.prod(values)


test = '''class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
'''
assert one(test) == 71

with open('input') as f:
    d = f.read()

assert one(d) == 18142
assert two(d) == 1069784384303

test2 = '''departure class: 0-1 or 4-19
departure row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
'''
assert two(test2) == 132
