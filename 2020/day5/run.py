
def get_row(val):
    code = "".join([x.replace("F", "1").replace("B", "0") for x in val[:6]])

    rows = range(0, 128)
    for c in val[:6]:
        if c == 'F':
            mid = int(len(rows) / 2)
            rows = rows[:mid]
        else:
            mid = int(len(rows) / 2)
            rows = rows[mid:]
    row = rows[0] if val[6] == 'F' else rows[1]
    r2 =int(code, 2) + 1
    print(f"{row} vs {r2}")
    return r2

def get_seat(val):
    code = "".join([x.replace("L", "0").replace("R", "1") for x in val[7:]])
    return int(code, 2) + 1

def decode(val):
    rows = range(0, 128)
    for c in val[:6]:
        if c == 'F':
            mid = int(len(rows) / 2)
            rows = rows[:mid]
        else:
            mid = int(len(rows) / 2)
            rows = rows[mid:]
    row = rows[0] if val[6] == 'F' else rows[1]

    seats = range(0, 8)
    for c in val[7:]:
        if c == 'L':
            mid = int(len(seats) / 2)
            seats = seats[:mid]
        else:
            mid = int(len(seats) / 2)
            seats = seats[mid:]
    seat = seats[0]
    return row * 8 + seat


def first(passes):
    ids = [get_row(x) * 8 + get_seat(x) for x in passes]
    ids.sort()
    return ids[-1]

    #
    #ids = [decode(x) for x in passes]
    #ids.sort()
    #return ids[-1]  # 987

def second(passes):
    ids = [decode(x) for x in passes]
    ids.sort()
    seat = ids[0] + 1
    while True:
        if seat in ids:
            seat += 1
            continue

        n = seat + 1
        if n not in ids:
            seat += 1
            continue
        
        p = seat - 1
        if n not in ids:
            seat += 1
            continue
        
        break
    return seat  # 603
        

with open('input') as f:
    passes = f.readlines()

a = first(passes)
print(f"Part 1: highest seat ID: {a}")

#b = second(passes)
#print(f"Part 2: My seat ID: {b}")

#for p in passes:
#    print(get_seat(p))

