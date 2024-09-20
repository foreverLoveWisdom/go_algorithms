package main

import (
	"log"
	"strconv"
)

func GenerateCracklePopSequence(start, end int) []string {
	const (
		divisorThree   = 3
		divisorFive    = 5
		divisorFifteen = 15
	)

	sequence := make([]string, 0, end-start+1)

	divisibleByThree := divisibleBy(divisorThree)
	divisibleByFive := divisibleBy(divisorFive)
	divisibleByFifteen := divisibleBy(divisorFifteen)

	for i := start; i <= end; i++ {
		switch {
		case divisibleByFifteen(i):
			sequence = append(sequence, "CracklePop")
		case divisibleByThree(i):
			sequence = append(sequence, "Crackle")
		case divisibleByFive(i):
			sequence = append(sequence, "Pop")
		default:
			sequence = append(sequence, strconv.Itoa(i))
		}
	}

	return sequence
}

func divisibleBy(divisor int) func(int) bool {
	return func(n int) bool {
		return n%divisor == 0
	}
}

func main() {
	const start = 1
	const end = 100

	sequence := GenerateCracklePopSequence(start, end)

	for _, item := range sequence {
		log.Println(item)
	}
}
