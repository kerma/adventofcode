import re 
import math

required_keys = {'byr', 'iyr', 'eyr', 'hgt', 'hcl', 'ecl', 'pid'}

def first(passports):
    valid = 0
    for p in passports:
        keys = set(x.split(":")[0] for x in p.split(" "))
        if required_keys.issubset(keys):
            valid += 1

    return valid

def second(passports):
    valid = 0
    for p in passports:
        check = dict(x.split(":") for x in p.split(" "))
        if "cid" in check.keys():
            del check["cid"]  # no need to check the optional value
        validated = [1 if is_valid(k,v) else 0 for k,v in check.items()]
        
        if len(validated) >= 7:
            valid += math.prod(validated)

    return valid

validators = {
    'hgt': re.compile(r'^(\d{2,3})(cm|in)$'),
    'hcl': re.compile(r'^#[a-f\d]{6}'),
    'ecl': re.compile(r'^(amb|blu|brn|gry|grn|hzl|oth)$'),
    'pid': re.compile(r'^\d{9}$')
}

def is_valid(key, val):

    if key == 'byr':
        return 1920 <= int(val) <= 2002

    if key == 'iyr':
        return 2010 <= int(val) <= 2020

    if key == 'eyr':
        return 2020 <= int(val) <= 2030

    if key == 'hgt':
        if (m := validators[key].match(val)) is None:
            return False

        v, t = m.groups()
        if t == 'cm':
            return 150 <= int(v) <= 193
        elif t == 'in':
            return 59 <= int(v) <= 76
        else:
            return False

    elif key in validators.keys():
        return validators[key].match(val)

    return False


with open('input') as f:
    passports = [x.replace("\n", " ").strip(" ") for x in f.read().split("\n\n")]

a = first(passports)
print(f"Part 1: Valid passports: {a}")

b = second(passports)
print(f"Part 2: Valid passports: {b}")

