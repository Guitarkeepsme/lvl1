// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func ownSet(arr []string) []string {
	mp := make(map[string]struct{})
	res := make([]string, 0)

	for _, word := range arr {
		mp[word] = struct{}{}
	}

	for key := range mp {
		res = append(res, key)
	}
	return res
}

func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Собственное множество: ", ownSet(set))
}
