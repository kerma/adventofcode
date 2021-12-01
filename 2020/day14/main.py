import re
from itertools import combinations_with_replacement, permutations

pattern = re.compile('([01X]{36})|\[(\d+)\] = (\d+)')


def apply_mask(value, mask: str, replace=('0', '1')) -> int:
    val = bin(int(value))[2:].zfill(len(mask))
    for i, m in enumerate(mask):
        if m in replace:
            val = val[0:i] + m + val[i+1:]
    return val


def get_addr(val) -> list:
    count = val.count("X")
    combinations = set(combinations_with_replacement([0, 1], count)).union(
                   set(combinations_with_replacement([1, 0], count)))
    perm = set()
    for c in combinations:
        for p in permutations(c, count):
            perm.add(p)

    addresses = []
    for p in perm:
        mask = val
        for x in p:
            mask = mask.replace('X', str(x), 1)
        addresses.append(int(mask, base=2))

    return addresses


def one(input_raw):
    memory = {}
    for cur, *rest in pattern.findall(input_raw):
        if cur != '':
            mask = cur
        else:
            addr, val = rest
            memory[addr] = int(apply_mask(val, mask), base=2)

    return sum(memory.values())


def two(input_raw):
    memory = {}
    for cur, *rest in pattern.findall(input_raw):
        if cur != '':
            mask = cur
        else:
            addr, val = rest
            addr = apply_mask(addr, mask, replace=('X', '1'))

            for a in get_addr(addr):
                memory[a] = int(val)

    return sum(memory.values())


test='''mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0'''
assert one(test) == 165

test2='''mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1'''
assert two(test2) == 208


with open('input') as f:
    data = f.read()

assert one(data) == 11884151942312
assert two(data) == 2625449018811
