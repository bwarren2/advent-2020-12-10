package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"time"
)

// NumsFromFile gets the nums
func NumsFromFile(filename string) (returnlist []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		text := scanner.Text()
		input, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			panic(err)
		}
		returnlist = append(returnlist, int(input))
	}
	// spew.Dump([]string)
	return
}

func MinMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

// Part1 answers part 1
func Part1(filename string) (count int) {
	nums := NumsFromFile(filename)
	_, max := MinMax(nums)
	nums = append(nums, max+3)
	sort.Ints(nums)

	ones := make([]int, 0)
	threes := make([]int, 0)

	var lower int
	for idx, value := range nums {
		if idx == 0 {
			lower = 0
		} else {
			lower = nums[idx-1]
		}
		if value-lower == 1 {
			ones = append(ones, 1)
		}
		if value-lower == 3 {
			threes = append(threes, 1)
		}
	}
	// fmt.Println(len(ones), len(threes))
	return len(ones) * len(threes)
}

var Memo = make(map[int]int)

// ConnectSum does work
func ConnectSum(target int, adapters []int) int {
	// fmt.Println(target, adapters)
	if val, ok := Memo[target]; ok {
		return val
	}
	if len(adapters) == 0 {
		return 1
	}
	for idx, value := range adapters {
		if value > target+3 {
			break
		} else {
			if idx+1 <= len(adapters) {
				Memo[target] += ConnectSum(value, adapters[idx+1:])
			}
		}
	}
	return Memo[target]
}

// Part2 answers part 2
func Part2(filename string) (count int) {
	nums := NumsFromFile(filename)
	_, max := MinMax(nums)
	nums = append(nums, max+3)
	sort.Ints(nums)

	return ConnectSum(0, nums)

}

func main() {
	start := time.Now()
	fmt.Println("part 2:", Part2("input.txt"))
	end := time.Now()
	fmt.Println(end.Sub(start))
}
