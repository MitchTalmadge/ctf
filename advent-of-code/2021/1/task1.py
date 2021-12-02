def main():
    # open input.txt and read lines
    with open('input.txt', 'r') as f:
        lines = f.readlines()

    # convert lines to ints
    lines = [int(line) for line in lines]

    # count lines where the previous is smaller
    count = 0
    for i in range(len(lines)):
        if i > 0:
            if lines[i] > lines[i-1]:
                count += 1

    # print result
    print(count)


if __name__ == "__main__":
    main()
