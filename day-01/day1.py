def part1_initial(line: str) -> int:
    for char in line:
        if char.isdigit():
            digit1 = char
    
    for char in line[::-1]:
        if char.isdigit():
            digit2 = char
    
    return int(digit1 + digit2)


def part1_optimized(line: str) -> int:
    digit1 = None
    digit2 = None

    for char in line:
        if char.isdigit():
            if not digit1:
                digit1 = char
            digit2 = char
    
    return int(digit1 + digit2)


digits = [
    "0", "zero",
    "1", "one",
    "2", "two",
    "3", "three",
    "4", "four",
    "5", "five",
    "6", "six",
    "7", "seven",
    "8", "eight",
    "9", "nine",
]


digits_dict = {
    "zero": "0",
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9",
}


def process_line_part2(line: str) -> int:
    digit1 = None
    digit2 = None

    index_digit1 = len(line)
    index_digit2 = -1

    for digit in digits:
        index_left = line.find(digit)
        index_right = line.rfind(digit)

        if index_left != -1 and index_left < index_digit1:
            index_digit1 = index_left
            digit1 = digit

        if index_right != -1 and index_right > index_digit2:
            index_digit2 = index_right
            digit2 = digit

    if not digit1.isdigit():
        digit1 = digits_dict[digit1]

    if not digit2.isdigit():
        digit2 = digits_dict[digit2]

    return int(digit1 + digit2)


def part2():
    with open("input.txt") as fp:
        lines = fp.readlines()

    _sum = 0

    for line in lines:
        _sum += process_line_part2(line)

    print(_sum)


if __name__ == "__main__":
    part2()
