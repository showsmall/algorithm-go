package structure

import "fmt"

type ArrayStack struct {
	data []interface{} // 数据
	top  int           // 栈顶指针
}

func NewArrayStack(capacity uint) *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, capacity, capacity),
		top:  -1,
	}
}

func (stack *ArrayStack) IsEmpty() bool {
	if stack.top < 0 {
		return true
	}
	return false
}

func (stack *ArrayStack) Push(value interface{}) error {
	if cap(stack.data) == stack.top+1 {
		return ErrorFullArray
	}
	stack.top++
	stack.data[stack.top] = value
	return nil
}

func (stack *ArrayStack) Pop() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	value := stack.data[stack.top]
	stack.top--
	return value
}

func (stack *ArrayStack) Top() interface{} {
	if stack.IsEmpty() {
		return nil
	}
	return stack.data[stack.top]
}

func (stack *ArrayStack) Print() {
	if stack.IsEmpty() {
		fmt.Println([]interface{}{})
	} else {
		fmt.Println(stack.data[0:stack.top])
	}
}
