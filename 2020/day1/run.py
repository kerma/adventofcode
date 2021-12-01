
def first(e):
    for n in e:
        f = 2020 - n
        if f in e:
            break
    return n * f  # 485739
    

def second(e):
    le = len(e)
    a, b, c = 0, 1, 2

    while True:
        if c == le:
            a += 1
            b = a + 1
            c = b + 1

        s = e[a] + e[b] + e[c]

        if s == 2020:
            return(e[a] * e[b] * e[c])  # 161109702

        if s < 2020:
            c += 1

        if s > 2020:
            b += 1 
            c = b + 1


with open('input') as f:
    expenses = [int(x) for x in f.readlines()]

expenses.sort()

a = first(expenses)
print(f"Part 1: prod: {a}")
b = second(expenses)
print(f"Part 2: prod {b}")

