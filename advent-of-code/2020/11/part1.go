package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	seats := [][]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := []int{}
		for _, c := range line {
			if c == 'L' {
				row = append(row, 0)
			} else if c == '#' {
				row = append(row, 1)
			} else {
				row = append(row, -1)
			}
		}
		seats = append(seats, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for {
		changed, updated_seats := shuffle(seats)
		if !changed {
			break
		}

		seats = updated_seats

		// for _, row := range seats {
		// 	for _, seat := range row {
		// 		if seat == -1 {
		// 			fmt.Printf(".")
		// 		} else if seat == 0 {
		// 			fmt.Printf("L")
		// 		} else {
		// 			fmt.Printf("#")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
	}

	occ := 0
	for _, row := range seats {
		for _, seat := range row {
			if seat == 1 {
				occ += 1
			}
		}
	}
	fmt.Println(occ)
}

func shuffle(seats [][]int) (changed bool, updated_seats [][]int) {
	updated_seats = make([][]int, len(seats))
	for i := range seats {
		updated_seats[i] = make([]int, len(seats[i]))
		copy(updated_seats[i], seats[i])
	}

	changed = false
	for y, row := range seats {
		for x, seat := range row {
			new_seat := transform(x, y, seats)
			if seat != new_seat {
				changed = true
				updated_seats[y][x] = new_seat
			}
		}
	}
	return
}

func transform(x, y int, seats [][]int) (seat int) {
	seat = seats[y][x]
	if seat == -1 {
		return
	} else if seat == 0 {
		if num_adjacent(x, y, seats) == 0 {
			seat = 1
		}
	} else {
		if num_adjacent(x, y, seats) >= 4 {
			seat = 0
		}
	}
	return
}

func num_adjacent(x, y int, seats [][]int) int {
	count := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if x+i < 0 || x+i >= len(seats[y]) {
				continue
			}
			if y+j < 0 || y+j >= len(seats) {
				continue
			}
			if seats[y+j][x+i] == 1 {
				count++
			}
		}
	}
	return count
}
