package main

import "fmt"

//队列(queue)， 是一种先进先出的线性表。通常用数据或者链表来实现队列。 队列只允许在后端插入，前端删除操作。
//性质：
//先进先出

type Queue struct {
	queue []interface{}
}

func newQueue() *Queue {
	return &Queue{
		queue: make([]interface{}, 0),
	}
}

// Push 入队列
func (q *Queue) Push(value interface{}) {
	q.queue = append(q.queue, value)
}

// Pop 出队列
func (q *Queue) Pop() interface{} {
	if len(q.queue) == 0 {
		return nil
	}
	value := q.queue[0]
	q.queue = q.queue[1:len(q.queue)]

	return value
}

// Size 队列大小
func (q *Queue) Size() int {
	return len(q.queue)
}

func main() {
	queue := newQueue()

	queue.Push(1)
	queue.Push(3)
	queue.Push(5)

	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
}
