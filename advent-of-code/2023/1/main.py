import os
import re

def main():
  print(part1("input.txt"))
  print(part2("input.txt"))

def part1(input_file_name):
  file_path = os.path.dirname(os.path.realpath(__file__)) + "/" + input_file_name
  with open(file_path, "r") as input_file:
    input = input_file.read().splitlines()
  
  total = 0
  for line in input:
    numbers = re.sub(r"[a-z]", "", line)
    total += int(numbers[0] + numbers[-1])
  return total

numbers = {
  "one": 1,
  "two": 2,
  "three": 3,
  "four": 4,
  "five": 5,
  "six": 6,
  "seven": 7,
  "eight": 8,
  "nine": 9
}

def part2(input_file_name):
  file_path = os.path.dirname(os.path.realpath(__file__)) + "/" + input_file_name
  with open(file_path, "r") as input_file:
    input = input_file.read().splitlines()
  
  total = 0
  for line in input:
    valid = ""
    i = 0
    while i < len(line):
      if re.match(r"[0-9]", line[i]):
        valid += line[i]
      else: 
        for number in numbers:
          if line[i:i+len(number)] == number:
            valid += str(numbers[number])
            break
      i += 1
    total += int(valid[0] + valid[-1])
  return total

if __name__ == "__main__":
  main()
