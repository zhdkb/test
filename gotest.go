package main

import (
	"fmt"
	"sort"
)

var m map[int]int

type MinStack struct {
	nums []int
}


func Constructor() MinStack {

	minStack := MinStack{nums: []int{}}
	m = map[int]int{}
	return minStack

}

func (this *MinStack) Push(val int) {

	if len(this.nums) == 0 {
		m[0] = val
	} else {
		m[len(this.nums)] = min(m[len(this.nums)-1], val)
	}
	this.nums = append(this.nums, val)

}

func (this *MinStack) Pop() {

	this.nums = this.nums[:len(this.nums)-1]

}

func (this *MinStack) Top() int {

	return this.nums[len(this.nums)-1]

}

func (this *MinStack) GetMin() int {

	return m[len(this.nums)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {

	var s1 string
	s1 = "fefer"
	s1 += "ewf"
	fmt.Println(s1[0:2])

	fmt.Println((8 + 1) / 2.0)

	anagramMap := make(map[string][]string)

	// 示例数据
	anagramMap["eat"] = append(anagramMap["eat"], "tea")
	anagramMap["eat"] = append(anagramMap["eat"], "ate")

	// 输出 map 的内容
	for key, value := range anagramMap {
		fmt.Printf("键: %s, 值: %v\n", key, value)
	}

	var age int = 9
	var num = 10
	fmt.Println("age =", age)
	fmt.Println("age =", num)

	a := "eevwfqf"
	s := []rune(a)
	sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
	fmt.Println(s)

	for i := 1; i <= 5; i++ {
		num += i
		fmt.Println(num)
	}

}
