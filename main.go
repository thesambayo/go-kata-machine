package main

import (
	"flag"
	"fmt"
	"go-kata-machine/scripts"
)

type Config struct {
	clear    bool
	generate bool
	day      bool
}

func main() {
	var config Config

	flag.BoolVar(&config.clear, "clear", false, "clear all day directories")
	flag.BoolVar(&config.generate, "generate", false, "generate new working day directory")
	flag.BoolVar(&config.day, "day", false, "get most recently created day directory")
	flag.Parse()

	// fmt.Println(config)
	if config.clear {
		scripts.ClearAllDayDirectories()
		return
	}

	if config.generate {
		scripts.GenerateNewDayDirectory()
		return
	}

	if config.day {
		fmt.Println(scripts.GetMostRecentlyCreatedDay())
		return
	}

	//foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	//fmt.Println(binarysearch.BinarySearch(foo, 99))
	////
	//bull := []bool{false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, false, true, true, true, true, true, true}
	//binarysearch.LinearSearch(bull) // 33
	////testData := make([]bool, 100)
	////brokenFloor := binarysearch.TwoCrystalBall(testData)
	////fmt.Println(brokenFloor)
	//
	//bar := []int{71, 81, 90, 69420, 420, 1, 3, 99, 4, 69, 1337}
	//sorted := bubbleSort.BubbleSort(bar)
	//fmt.Println(sorted)

	//testQueue := queue.Queue[int]{}
	//fmt.Println(testQueue)
	//testQueue.Enqueue(5)
	//fmt.Println(testQueue, "head and tail pointers are same")
	//testQueue.Enqueue(7)
	//fmt.Println(testQueue, "head and tail pointers are diff")
	//fmt.Println(testQueue.Deque()) // 5
	//fmt.Println(testQueue, "head and tail pointers are same again")
	//fmt.Println(testQueue.Deque()) // 7
	//fmt.Println(testQueue, "head and tail pointers are nil")

	//testQueue.Enqueue(5)
	//testQueue.Enqueue(7)
	//testQueue.Enqueue(9)
	//
	//fmt.Println(testQueue.Deque()) // 5
	//fmt.Println(testQueue.Length)  // 2
	//
	//testQueue.Enqueue(11)
	//fmt.Println(testQueue.Deque()) // 7
	//fmt.Println(testQueue.Deque()) // 9
	//fmt.Println(testQueue.Peek())  // 11
	//fmt.Println(testQueue.Deque()) // 11
	//fmt.Println(testQueue)
	//fmt.Println(testQueue.Deque()) // 0
	//fmt.Println(testQueue.Length)  // 0
	//
	//testQueue.Enqueue(69)
	//fmt.Println(testQueue.Peek()) // 69
	//fmt.Println(testQueue.Length) // 1

	//stackList := stack.Stack[int]{}
	//fmt.Println(stackList)
	//stackList.Push(5)
	//fmt.Println(stackList, "head is not nil")
	//stackList.Push(7)
	//fmt.Println(stackList.Length) //2
	//fmt.Println(stackList.Pop())  // 7
	//fmt.Println(stackList, "head and tail pointers are same again")
	//fmt.Println(stackList.Pop()) // 5
	//
	//fmt.Println(stackList.Length) //0
}
