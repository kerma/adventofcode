import re

pattern = re.compile(r'(\w)(\d+)')

DIRECTIONS = (
    NORTH   := 'N', 
    EAST    := 'E', 
    SOUTH   := 'S', 
    WEST    := 'W'
)

FORWARD = 'F'
LEFT    = 'L'
RIGHT   = 'R'

def move_x(direction, x, steps):
    steps = int(steps)
    x += steps if direction == EAST else -steps
    return x


def move_y(direction, y, steps):
    steps = int(steps)
    y += steps if direction == SOUTH else -steps
    return y


def one(input_raw):

    def move_forward(d, x, y, z):
        z = int(z)
        if d == EAST:
            x += z
        elif d == WEST:
            x -= z
        elif d == NORTH:
            y -= z
        elif d == SOUTH:
            y += z
        return x, y

    def turn(direction, action, steps):
        start = DIRECTIONS.index(direction)
        rotate = int(int(steps) / 90)

        if action == RIGHT:
            new = start + rotate
            new -= 4 if new >= 4 else 0
        else:
            new = start - rotate

        return DIRECTIONS[new]

    def move(x, y, direction, instruction):
        action, steps = instruction

        if action == FORWARD:
            x, y = move_forward(direction, x, y, steps)
        elif action in (LEFT, RIGHT):
            direction = turn(direction, action, steps)
        elif action in (EAST, WEST):
            x = move_x(action, x, steps)
        elif action in (NORTH, SOUTH):
            y = move_y(action, y, steps)

        return x, y, direction

    x, y, direction = 0, 0, EAST
    for instruction in pattern.findall(input_raw):
        x, y, direction = move(x, y, direction, instruction)

    return abs(x + y)


def two(input_raw):

    def move_ship(x, y, wx, wy):
        return x  + wx*int(steps), y + wy*int(steps)

    def turn_waypoint(action, steps, x, y):
        for _ in range(0, int(int(steps) / 90)):
            x, y = y, x
            x, y = (-1*x, y) if action == RIGHT else (x, -1*y)
        return x, y

    def move_waypoint(x, y, action, steps):
        if action in (LEFT, RIGHT):
            x, y = turn_waypoint(action, steps, x, y)
        elif action in (EAST, WEST):
            x = move_x(action, x, steps)
        elif action in (NORTH, SOUTH):
            y = move_y(action, y, steps)

        return x, y

    x, y, wx, wy = 0, 0, 10, -1 
    for action, steps in pattern.findall(input_raw):
        if action == FORWARD:
            x, y = move_ship(x, y, wx, wy)
        else:
            wx, wy = move_waypoint(wx, wy, action, steps)

    return abs(x + y)


with open('input') as f:
    data = f.read()

assert one(data) == 1177
assert two(data) == 46530

