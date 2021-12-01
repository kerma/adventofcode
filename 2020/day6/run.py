from functools import reduce

def first(a):
    ans = [set(x.replace("\n", " ").replace(" ", "").strip(" ")) for x in a]
    return sum([len(x) for x in ans])  # 6259


def second(a):
    ans = [x.replace(" ", "").strip("\n").split("\n") for x in a]
    return sum([
        len(
            reduce(lambda x,y: set(x)&set(y), group)
        ) for group in ans
    ])  # 3178

with open('input') as f:
    answers = f.read().split("\n\n")


a = first(answers)
print(f"Part 1: Sum of the counts: {a}")

b = second(answers)
print(f"Part 2: Sum of the counts: {b}")





def second_a(a):
    ans = [x.replace(" ", "").strip("\n").split("\n") for x in a]
    counts = []
    for group in ans:
        inter = set(group[0])
        for g in group[1:]:
            inter = inter.intersection(set(g))
        counts.append(len(inter))
    return sum(counts)  # 3178 



