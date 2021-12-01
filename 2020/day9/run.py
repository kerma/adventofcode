from collections import deque


def find_invalid(numbers: list, preamble: int) -> int:
    val  = 0
    q = deque(numbers[preamble:])
    for i in range(len(q)):
        val = q.popleft()
        prev = numbers[i:i+preamble]
        prev.sort()
        
        valid = False
        for n in prev:
            s = val - n
            if s in prev:
                valid = True
                break

        if not valid:
            break

    return val


def check_a(search: list, expect: int) -> list:
    for i in range(1, len(search)): 
        total = sum(found := search[0:i])
        if total < expect:
            continue
        if total == expect:
            return found
        break
        
    return []

def check(search: list, expect: int) -> list:
    total = search[0]
    for i in range(1, len(search)): 
        total += search[i]
        if total < expect:
            continue
        if total == expect:
            return search[0:i+1]
        break
        
    return []


def one(numbers: list, preamble: int = 25) -> int:
    return find_invalid(numbers, preamble)


def two(numbers: list, preamble: int = 25) -> int:
    invalid = find_invalid(numbers, preamble)
    for i, _ in enumerate(numbers):
        if (found := check(numbers[i:], invalid)):
            break
            
    found.sort()
    return found[0] + found[-1]  # 13549369


test = [int(x) for x in '''35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576'''.split("\n")]

assert one(test, 5) == 127
assert two(test, 5) == 62

with open('input') as f:
    numbers = [int(x) for x in f.readlines()]

assert one(numbers) == 88311122
assert two(numbers) == 13549369
