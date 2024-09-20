package main

import (
	"reflect"
	"testing"
)

// Write a program that prints out the numbers 1 to 100 (inclusive). If the number is divisible by 3, print Crackle instead of the number. If it's divisible by 5, print Pop. If it's divisible by both 3 and 5, print CracklePop. You can use any language.

func TestGenerateCracklePopSequence(t *testing.T) {
	expected := []string{
		"1", "2", "Crackle", "4", "Pop", "Crackle", "7", "8", "Crackle", "Pop",
		"11", "Crackle", "13", "14", "CracklePop", "16", "17", "Crackle", "19", "Pop",
		"Crackle", "22", "23", "Crackle", "Pop", "26", "Crackle", "28", "29", "CracklePop",
		"31", "32", "Crackle", "34", "Pop", "Crackle", "37", "38", "Crackle", "Pop",
		"41", "Crackle", "43", "44", "CracklePop", "46", "47", "Crackle", "49", "Pop",
		"Crackle", "52", "53", "Crackle", "Pop", "56", "Crackle", "58", "59", "CracklePop",
		"61", "62", "Crackle", "64", "Pop", "Crackle", "67", "68", "Crackle", "Pop",
		"71", "Crackle", "73", "74", "CracklePop", "76", "77", "Crackle", "79", "Pop",
		"Crackle", "82", "83", "Crackle", "Pop", "86", "Crackle", "88", "89", "CracklePop",
		"91", "92", "Crackle", "94", "Pop", "Crackle", "97", "98", "Crackle", "Pop",
	}

	result := GenerateCracklePopSequence(1, 100)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("GenerateCracklePopSequence() = %v; want %v", result, expected)
	}
}
