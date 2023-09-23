package main

import (
	"fmt"
	"sort"
)

func main() {
	//栈
	stack := make([]int, 0)
	stack = append(stack, 10)
	v := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	if len(stack) == 0 {

	}

	//队列
	queue := make([]int, 0)
	queue = append(queue, 10)
	q := queue[0]
	queue = queue[1:]
	fmt.Printf("%d", q)
	if len(queue) == 0 {

	}

	//map
	m := make(map[string]int)
	m["hello"] = 1
	delete(m, "hello")
	for k, v := range m {
		print(k, v)
	}
	//map 键需要可比较，不能为 slice、map、function
	//map 值都有默认值，可以直接操作默认值，如：m[age]++ 值由 0 变为 1
	//比较两个 map 需要遍历，其中的 kv 是否相同，因为有默认值关系，所以需要检查 val 和 ok 两个值
	fmt.Printf("%d%d", v, stack)

	s := []int{1, 2, 3, 4, 5, 6, 9, 4}
	sort.Ints(s)
	println(s)
}
