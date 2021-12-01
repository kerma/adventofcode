import time
from collections import deque

def one(starting, to=2020):
    starting.reverse()
    init = len(starting) + 1
    d = deque(starting)
    s = time.time()
    for i in range(init, to+1):
        if i % 10000 == 0:
            e = time.time()
            print(e-s)
            s = e
        last = d.popleft()
        if last not in d:
            d.insert(0, last)
            d.insert(0, 0)
            continue
        last_seen = len(d) - d.index(last)
        new_number = i - 1 - last_seen 
        d.insert(0, last)
        d.insert(0, new_number)

    return d.popleft()


class Memory:

    def __init__(self):
        self._last = {}
        self._numbers = {}

    def has(self, val):
        return val in self._numbers

    def add(self, pos, val):
        if self._last:
            self._numbers.update(self._last)
            if val in self._last
        self._last[val] = deque([pos], maxlen=2)

    def get(self, val):
        return self._numbers[val]

    def __str__(self):
        return str(self._numbers)
           

def two(numbers, to=2020):
    #print(f"{numbers}")
    init = len(numbers) + 1

    positions = { x : deque([i+1], maxlen=2) for i, x in enumerate(numbers) }
    current, numbers = numbers[-1], set(numbers[:-1])

    for i in range(init, to+1):
        if current not in numbers:
            numbers.add(0)
            if d := positions.get(0):
                d.append(i)
            else:
                positions[0] = deque([i], maxlen=2)
            current = 0
            continue

        pos = positions[current]
        print(pos)
        last_pos = pos.pop()
        print(last_pos)
        new_nr = i - last_pos
        print(new_nr)


    #print(memory)
    print(new_number)
    return new_number
        


numbers = [int(x) for x in '17,1,3,16,19,0'.split(',')]
test = [int(x) for x in '0,3,6'.split(',')]
assert one(test) == 436
assert one(numbers) == 694

test = [int(x) for x in '0,3,6'.split(',')]
assert two(test, 10) == 436

more_tests = {
        175594: [int(x) for x in '0,3,6'.split(',')],
        #2578: [int(x) for x in '0,3,6'.split(',')],
        #3544142: [int(x) for x in '0,3,6'.split(',')],
        #261214: [int(x) for x in '0,3,6'.split(',')],
        #6895259: [int(x) for x in '0,3,6'.split(',')],
        #18: [int(x) for x in '0,3,6'.split(',')],
        #362: [int(x) for x in '0,3,6'.split(',')],
    }
for expect, numbers in more_tests.items():
    try:
        assert (res := one(numbers, to=30000000)) == expect
    except:
        print(f"{expect} != {res}")
        raise
one_tests = {
        436: [int(x) for x in '0,3,6'.split(',')],
        1: [int(x) for x in '1,3,2'.split(',')],
        10: [int(x) for x in '2,1,3'.split(',')],
        27: [int(x) for x in '1,2,3'.split(',')],
        78: [int(x) for x in '2,3,1'.split(',')],
        438: [int(x) for x in '3,2,1'.split(',')],
        1836: [int(x) for x in '3,1,2'.split(',')],
    }
for expect, numbers in one_tests.items():
    try:
        assert (res := one(numbers, to=2020)) == expect
    except:
        print(f"{expect} != {res}")
        raise
