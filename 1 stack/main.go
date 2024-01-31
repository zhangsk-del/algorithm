package main

import "fmt"

// 栈(stack)，是计算机科学中一种特殊的串列形式的抽象数据类型，其也通常使用链表或者数据来实现
// 与队列不用，栈的性质是后进先出。也就是只能总栈的顶部插入元素与取出元素。
// 特点 先进后出

type Stack struct {
	items []interface{}
}

func New() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push 入栈
func (s *Stack) Push(value interface{}) {
	s.items = append(s.items, value)
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	if len(s.items) == 0 {
		return nil
	}
	value := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return value
}

func main() {
	stack := New()

	stack.Push(1)
	stack.Push(3)
	stack.Push(5)

	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
