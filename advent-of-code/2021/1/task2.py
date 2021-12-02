def main():
    # open input.txt and read lines
    with open('input.txt', 'r') as f:
        lines = f.readlines()

    # convert lines to ints
    lines = [int(line) for line in lines]

    # compute 3-sliding window sums
    sums = [sum(lines[i:i+3]) for i in range(len(lines)-2)]

    # iterate over sums and compare to previous; add to count if previous is smaller
    count = 0
    for i in range(len(sums)):
        if i > 0:
            if sums[i] > sums[i-1]:
                count += 1

    # print result
    print(count)


if __name__ == "__main__":
    main()
