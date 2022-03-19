package main

import (
	"container/list"
	"fmt"
)

type Stack struct {
	stack *list.List
}

func (c *Stack) Push(value interface{}) {
	c.stack.PushFront(value)
}

func (c *Stack) Pop() error {
	if c.stack.Len() > 0 {
		ele := c.stack.Front()
		c.stack.Remove(ele)
	}
	return fmt.Errorf("Pop Error: Stack is empty")
}

func (c *Stack) Front() (interface{}, error) {
	if c.stack.Len() > 0 {
		if val, ok := c.stack.Front().Value.(interface{}); ok {
			return val, nil
		}
		return nil, fmt.Errorf("Peep Error: Stack Datatype is incorrect")
	}
	return nil, fmt.Errorf("Peep Error: Stack is empty")
}

func (c *Stack) Size() int {
	return c.stack.Len()
}

func (c *Stack) Empty() bool {
	return c.stack.Len() == 0
}
