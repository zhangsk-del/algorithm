package main

import "fmt"

// 整形
// 一个字节代表8个bit

// int8 -2^7~2^7-1  int16 int32 int64

// 浮点型
// float32 float64

// 布尔型
// bool 一个字节

// 字节
// byte

// array struct string

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func main() {
	/* 初始化链表 1 -> 3 -> 2 -> 5 -> 4 */
	// 初始化各个节点
	n0 := NewListNode(1)
	n1 := NewListNode(3)
	n2 := NewListNode(2)
	n3 := NewListNode(5)
	n4 := NewListNode(4)
	// 构建节点之间的引用
	n0.Next = n1
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	insertNode(n0, NewListNode(0))
	remove(n0)
	node := access(n0, 2)
	fmt.Println("ACCESS:", node.Val)

	for n0 != nil {
		fmt.Println(n0.Val)
		n0 = n0.Next
	}

}

/* 在链表的节点 n0 之后插入节点 P */
func insertNode(head *ListNode, p *ListNode) {
	if head == nil || p == nil {
		return
	}
	next := head.Next
	head.Next = p
	p.Next = next
}

/* 删除链表的节点 n0 之后的首个节点 */
func remove(n0 *ListNode) {
	if n0.Next == nil {
		return
	}
	p := n0.Next
	n1 := p.Next
	n0.Next = n1
}

/* 访问链表中索引为 index 的节点 */
func access(head *ListNode, index int) *ListNode {
	for i := 0; i < index; i++ {
		if head == nil {
			return nil
		}
		head = head.Next
	}
	return head
}

/* 在链表中查找值为 target 的首个节点 */
// 查找其中值为 target 的节点，输出该节点在链表中的索引。此过程也属于线性查找
func findNode(head *ListNode, target int) int {
	index := 0
	for head != nil {
		if head.Val == target {
			return index
		}
		head = head.Next
		index++
	}
	return -1
}
