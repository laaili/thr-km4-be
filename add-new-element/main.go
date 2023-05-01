package main

import "fmt"

func AddElement(data []int, newData int, position string) []int {
	if position == "up" {
		return append([]int{newData}, data...)
	} else if position == "down" {
		return append(data, newData)
	}
	return nil
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	new := 6
	post := "up"
	result := AddElement(data, new, post)
	fmt.Println(result)

	data2 := []int{1, 2, 3, 4, 5}
	new2 := 6
	post2 := "down"
	result2 := AddElement(data2, new2, post2)
	fmt.Println(result2)
}
