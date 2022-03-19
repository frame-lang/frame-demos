package main

import "fmt"

func (m *historyStateContextStruct) print(msg string) {
	fmt.Println(msg)
}

func (m *historyBasicStruct) print(msg string) {
	fmt.Println(msg)
}
