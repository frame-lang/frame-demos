package main

import "fmt"

func (m *fizzBuzzStruct) print(msg string) {
	fmt.Print(msg)
}
func (m *fizzBuzzStruct) gt_100(i int) bool {
	return i > 100
}
func (m *fizzBuzzStruct) mod3_eq0(i int) bool {
	return i%3 == 0
}
func (m *fizzBuzzStruct) mod5_eq0(i int) bool {
	return i%5 == 0
}
func (m *fizzBuzzStruct) plus_1(i int) int {
	return i + 1
}
