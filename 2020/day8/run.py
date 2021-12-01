from collections import deque
from dataclasses import dataclass


@dataclass
class Result:
    completed: bool
    acc: int
    visited: list


def execute(instructions):
    index = 0
    r = Result(False, 0, [])

    while index not in r.visited:
        if index == len(instructions):
            r.completed = True
            break

        r.visited.append(index)

        ins, n = instructions[index].split(" ")
        r.acc += int(n) if ins == "acc" else 0
        index += int(n) if ins == "jmp" else 1

    return r
    

def flip(val):
    return val.replace('jmp', 'nop') if 'jmp' in val else val.replace('nop', 'jmp')


def one(instructions):
    return execute(instructions).acc  # 1859


def two_a(instructions):
    def update(instructions, idx, skipped):
        if "acc" in instructions[idx]:
            return instructions, idx+1, skipped+1

        if idx > 0:
            p = idx - skipped - 1
            instructions[p] = flip(instructions[p])
        instructions[idx] = flip(instructions[idx])

        return instructions, idx+1, 0

    index, skipped = 0, 0
    while (e := execute(instructions)).completed != True:
        instructions, index, skipped = update(instructions, index, skipped)
        
    return e.acc  # 1235


def two_b(instructions):
    q = deque()
    def update(instructions, last=0):
        while (index := q.pop()):
            if "acc" in instructions[index]:
                continue

            if last != 0:
                instructions[last] = flip(instructions[last])
            instructions[index] = flip(instructions[index])
            break

        return instructions, index

    last = 0
    while (e := execute(instructions)).completed != True:
        q = q or deque(e.visited)
        instructions, last = update(instructions, last)
        
    return e.acc  # 1235


with open('input') as f:
    ins = [x.strip("\n") for x in f.readlines()]

assert one(ins) == 1859
assert two_a(ins) == 1235
assert two_b(ins) == 1235

