import math
from collections import deque
from itertools import count

def parse(data, ignore_x=True):
    ts, bus_ids  = data.split("\n")
    ts = int(ts)
    if ignore_x:
        bus_ids = [int(x) for x in bus_ids.split(",") if x != 'x']
    else:
        bus_ids = [0 if x == 'x' else int(x) for x in bus_ids.split(",")]

    return ts, bus_ids

def one(data):
    ts, bus_ids = parse(data)
    times = []
    for i in bus_ids:
        x = math.ceil(ts / int(i)) * i
        times.append((i, x-ts))

    times.sort(key=lambda x: x[1]) 

    return math.prod(times[0])

def by_force(bus_ids):
    start = max(bus_ids)*2
    counter = count(start)
    while i := next(counter):
        ids = iter(bus_ids)
        prev = next(ids)
        checksum = 1
        timestamp = prev*i
        for bus in ids:
            found = False
            if bus == 0:
                checksum += 1
                continue
            
            if (current := math.ceil(timestamp / bus) * bus) - timestamp == checksum:
                found = True
                checksum = 1
                timestamp = current
                prev = bus
            else:
                break

        if found:
            break

    return i * bus_ids[0]


def two(data):
    _, bus_ids = parse(data, ignore_x=False)
    return by_force(bus_ids)

tests = {
    3417:       '1\n17,x,13,19',
    754018:     '1\n67,7,59,61',
    779210:     '1\n67,x,7,59,61',
    1068781:    '1\n7,13,x,x,59,x,31,19',
    1261476:    '1\n67,7,x,59,61',
    1202161486: '1\n1789,37,47,1889'
}

for expect, data in tests.items():
    assert two(data) == expect

with open('input') as f:
    data = f.read().strip("\n")

assert one(data) == 370
