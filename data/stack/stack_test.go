package stack

import (
	"fmt"
	"testing"
)

func TestLinkStack(t *testing.T) {
	var stacks = NewLinkStack()
	stacks.PushData("1")
	stacks.PushData("2")
	stacks.PushData("3")
	stacks.PushData("4")
	stacks.PushData("5")

	for data, l, _ := stacks.PullData(); data != nil; data, l, _ = stacks.PullData() {
		fmt.Println(l, data)
	}

	stacks.PushData("1")
	stacks.PushData("2")
	data, _, _ := stacks.PullData()
	fmt.Println(data)
	data, _, _ = stacks.PullData()
	fmt.Println(data)
	stacks.PushData("3")
	data, _, _ = stacks.PullData()
	fmt.Println(data)
}

func TestArrayStack(t *testing.T) {
	var stacks = NewArrayStack(5)
	stacks.PushData("1")
	stacks.PushData("2")
	stacks.PushData("3")
	stacks.PushData("4")
	stacks.PushData("5")

	for data, l, _ := stacks.PullData(); data != nil; data, l, _ = stacks.PullData() {
		fmt.Println(l, data)
	}

	stacks.PushData("1")
	stacks.PushData("2")
	data, _, _ := stacks.PullData()
	fmt.Println(data)
	data, _, _ = stacks.PullData()
	fmt.Println(data)
	stacks.PushData("3")
	data, _, _ = stacks.PullData()
	fmt.Println(data)
}
