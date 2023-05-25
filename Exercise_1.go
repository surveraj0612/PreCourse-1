



package main

import (
	"fmt"
)

type StackArray struct {
  s []int
}

func (sa *StackArray) isEmpty() bool {
  length := len(sa.s)
  return length == 0
}

func (sa *StackArray) Push(value int) {
  sa.s = append(sa.s,value)
}

func (sa *StackArray) Pop() int {
  length := len(sa.s)
  rem := sa.s[length-1]
  sa.s = sa.s[:length-1]
  return rem
}

func (sa *StackArray) Peek() int {
  length := len(sa.s)
  tp := sa.s[length-1]
  return tp
}

func (sa *StackArray) Size() int {
  length := len(sa.s)
  return length
}

func (sa *StackArray) Show() {
  length := len(sa.s)
  for i := 0; i < length; i++ {
    fmt.Print(sa.s[i], "")
  }
}




func main() {

  var stack StackArray
  stack.Push(1)
  stack.Push(2)
  stack.Push(3)

  fmt.Println("Printng the stack:")    
  stack.Show() // Print the stack
  fmt.Println("Checking size of stack:")
  fmt.Println(stack.Size()) // Get size
  fmt.Println("Removing an Item:")
  stack.Pop() // Remove an item
  stack.Show() 
  fmt.Println("Getting last Added Item in Stack")
  fmt.Println(stack.Peek()) // Get last item
}

