package main

import "fmt"

// Task 11

// Реализовать пересечение двух неупорядоченных множеств.

// Time: O(n)
// Space: O(n)
func intersect(a, b *[]int) []int {
	cnt := make(map[int]int)

	for _, val := range *a {
		cnt[val]++
	}

	var res []int

	for _, val := range *b {
		if cnt[val] > 0 {
			res = append(res, val)
			cnt[val]--
		}
	}

	return res
}

func main() {
	a := []int{7, 6, 2, 4, 3, 8}
	b := []int{3, 5, 6, 7, 9, 1}

	fmt.Println(intersect(&a, &b))
}
