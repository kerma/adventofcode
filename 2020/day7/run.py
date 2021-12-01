import re

def first(rules):
    def search(rules, bag):
        bags = [x for x in rules if bag in x and not x.startswith(bag)]
        return [" ".join(x.split(" ")[0:2]) for x in bags]

    bags = ['shiny gold']
    for b in bags:
        bags.extend(search(rules, b))

    return len(set(bags)) - 1  # 296
        

def second(rules):

    d = {}
    def count(bag):
        return 1 + sum(n * count(b) for b, n in d[bag].items())

    p = re.compile(r" (\d+) (\w+ \w+) bags?")
    for r in rules:
        bag = r.split(" bags contain")[0]
        contents = re.findall(p, r)
        d[bag] = { x[1]: int(x[0]) for x in contents }
    
    return count("shiny gold") - 1  # 9339


with open('input') as f:
    rules = [x.strip("\n") for x in f.readlines()]

print(first(rules))
print(second(rules))

test_rules = '''light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.'''.split("\n")

test_rules2 = '''shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.'''.split("\n")

print(second(test_rules))
print(second(test_rules2))
