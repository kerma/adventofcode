
def first(passwords):
    invalid = 0
    for line in passwords:
        min, max, char, val = split(line)
        char_count = val.count(char)
        if char_count < min or char_count > max:
            invalid += 1

    return len(passwords) - invalid  # 515


def second(passwords):
    valid = 0
    for line in passwords:
        p1, p2, char, val = split(line)
        p1 -= 1
        p2 -= 1

        if val[p1] == char or val[p2] == char:
            if not (val[p1] == char and val[p2] == char):
                valid += 1

    return valid  # 711


def split(line):
    rnge, char, val = line.split(' ')
    char = char.strip(":")
    a, b = [int(x) for x in rnge.split('-')]
    
    return a, b, char, val


with open('input') as f:
    passwords = [x.strip('\n') for x in f.readlines()]

a = first(passwords)
print(f"Part 1, valid password count: {a}")
b = second(passwords)
print(f"Part 2, valid password count: {b}")

