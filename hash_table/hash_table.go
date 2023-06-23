package hash_table

import (
	"fmt"
	"math"
)

// 拉链法: 数组 + 链表

type HashTable interface {
	Get(int) (interface{}, bool)
	Set(int, interface{})
	Delete(int)
	Traverse()
}

// Node kv节点
type Node struct {
	Key   int
	Value interface{}
	Next  *Node
}

type SimpleHash struct {
	BasicArr []*Node //底层数组
	Size     int     //元素个数
	Len      int     //数组长度
}

// 确保实现了所有接口方法
var _ HashTable = &SimpleHash{}

// NewSimpleHash 构造函数
func NewSimpleHash(len int) *SimpleHash {
	return &SimpleHash{
		BasicArr: make([]*Node, len),
		Size:     0,
		Len:      len,
	}
}

func (h *SimpleHash) get(key int) (*Node, bool) {
	exist := false
	index := key % h.Len
	head := h.BasicArr[index]
	//查找
	for head != nil {
		if head.Key == key {
			exist = true
			break
		}
		head = head.Next
	}
	return head, exist
}

func (h *SimpleHash) Get(key int) (value interface{}, isExist bool) {
	node, isExist := h.get(key)
	if isExist {
		value = node.Value
	}
	return
}

// Set 设置kv键值对
func (h *SimpleHash) Set(key int, value interface{}) {
	node, isExist := h.get(key)
	if isExist {
		node.Value = value
		return
	}
	//不存在时新建
	h.Size++
	index := key % h.Len
	head := h.BasicArr[index]
	newNode := &Node{Key: key, Value: value}
	if head != nil {
		newNode.Next = head // 头插法
	}
	h.BasicArr[index] = newNode
}

func (h *SimpleHash) Delete(key int) {
	_, isExist := h.get(key)
	//不存在时直接返回
	if !isExist {
		return
	}
	h.Size--
	index := key % h.Len
	head := h.BasicArr[index]
	if head.Key == key {
		h.BasicArr[index] = head.Next
		return
	}
	for head.Next != nil {
		next := head.Next
		if next.Key == key {
			head.Next = next.Next
			return
		}
		head = head.Next
	}
}

func (h *SimpleHash) Traverse() {
	for _, head := range h.BasicArr {
		traverList(head)
	}
}

func traverList(head *Node) {
	if head == nil {
		return
	}
	for head != nil {
		if head.Next == nil {
			fmt.Printf("%v-->%v\n", head.Key, head.Value)
			return
		}
		fmt.Printf("%v-->%v ", head.Key, head.Value)
		head = head.Next
	}
}

// 开发地址法

const N = 200003
const null = math.MaxInt

// 底层数组
var buf [N]int

func init() {
	for i := 0; i < N; i++ {
		buf[i] = null
	}
}

// 查找位置
func find(key int) int {
	k := (key%N + N) % N
	for buf[k] != null && buf[k] != key { //说明该位置被占用
		k++         //开放寻址法向后一位移动
		if k == N { //当移动到最后一位时，循环到第一个位置
			k = 0
		}
	}
	// 返回下标，如果x在哈希数组h中，就返回x在h中的位置，
	// 如果不在，就返回应该存储的位置
	return k
}

func Set(key int) {
	k := find(key)
	buf[k] = key
}

func Get(key int) (value int, isExist bool) {
	i := find(key)
	if buf[i] != null {
		return buf[i], true
	} else {
		return null, false
	}
}
