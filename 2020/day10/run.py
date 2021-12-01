from pprint import pprint
from collections import deque
from copy import copy

def one(jolts):
    jolts.sort()
    diff = {
        1: 1,
        2: 1,
        3: 1,
    }
    for i, j in enumerate(jolts):
        if i == 0:
            prev = j
            continue
        d = j - prev
        diff[d] += 1
        prev = j
        
    return diff[1] * diff[3]


def two(jolts):
    jolts += 0, max(jolts)+3  # add 0 and built in joltage adapter value to list
    jolts.sort()
    jolts = iter(jolts)

    counts = deque(maxlen=3)
    counts.append(
        (next(jolts), 1)
    )  # add a first jolt (max+3) with count 1

    for j in jolts:
        counts.append((
            j, 
            s := sum(k for jolt, k in counts if j - jolt <= 3)
        ))
    return s


with open('input') as f:
    jolts = [int(x) for x in f.readlines()]


assert one(jolts) == 1980
assert two(jolts) == 4628074479616
