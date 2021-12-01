import math

def first(rows):
    return walk(rows, 3, 1)  # 284

def second(rows):
    trees = []
    for right, down in [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]:
        trees.append(walk(rows, right, down))
    return math.prod(trees)  # 3510149120

def walk(rows, right, down):
    pos = right
    line = down
    tree_count = 0

    while line < len(rows):
        row = rows[line]

        if pos >= len(row):
            pos = pos - len(row)

        tree_count += 1 if row[pos] == "#" else 0

        pos += right
        line += down

    return tree_count


with open('input') as f:
    inp = [x.strip('\n') for x in f.readlines()]


a = first(inp)
print(f"Part 1: Encoutered {a} trees")
b = second(inp)
print(f"Part 2: {b}")
