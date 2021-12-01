from pprint import pprint

EMPTY = "L"
OCCUPIED = "#"


def flip_if(seat: str, a: list, tolerate=4) -> str:
    if seat == EMPTY and OCCUPIED not in a:
        return OCCUPIED
    if seat == OCCUPIED and sum(1 for x in a if x == OCCUPIED) >= tolerate:
        return EMPTY
    return seat


def one(seats):

    LAST_IDX = len(seats)
    ROW_LEN = len(seats[0])

    def get_seats(seat_n, row_n: int) -> list:
        seats = [x for x in (seat_n-1, seat_n, seat_n+1) if x >= 0 and x < ROW_LEN]
        positions = [(row_n-1, x) for x in seats if row_n > 0]

        for x in seats:
            if row_n < LAST_IDX-1:
                positions.append((row_n+1, x))
        for x in (seat_n-1, seat_n+1):
            if x >= 0 and x < ROW_LEN:
                positions.append((row_n, x))
        return positions

    def get_adjacent(seat_n, row_n: int, seats: list) -> list:
        adjacent = []
        for rn, sn in get_seats(seat_n, row_n):
            for x in seats[rn][sn]:
                adjacent.append(x) 
        return adjacent

    def transform(seats: list) -> list:
        transformed = []
        for rn, row in enumerate(seats):
            new_row = ""
            for sn, seat in enumerate(row):
                new_row += flip_if(seat, get_adjacent(sn, rn, seats))
            transformed.append(new_row)
        return transformed

    prev = seats
    while (seats := transform(seats)) != prev:
        prev = seats

    return sum(1 for x in "".join(seats) if x == OCCUPIED)


def two(seats):

    return 0


test_input='''L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL'''.split("\n")

assert one(test_input) == 37
#assert two(test_input) == 26


asd1 = '''.......#.
...#.....
.#.......
.........
..#L....#
....#....
.........
#........
...#.....'''.split("\n")

two(asd1)

#with open('input') as f:
#    seats = [x.strip("\n") for x in f.readlines()]


#import time

#s = time.time()
#assert one(seats) == 2310
#print(time.time() - s)
