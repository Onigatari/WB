package main

import "fmt"

// Task 12

// Имеется последовательность строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.

// Time: O(n)
// Space: O(n)
func createStringSet(strings []string) []string {
	strMap := make(map[string]bool)

	for _, val := range strings {
		strMap[val] = true
	}

	res := make([]string, 0, len(strMap))
	for key := range strMap {
		res = append(res, key)
	}

	return res
}

func main() {
	var slice = []string{"cat", "name", "cat", "dog", "cat", "tree", "name"}

	fmt.Println(createStringSet(slice))
}
