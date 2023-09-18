// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

// Собственным множеством будет слайс строк
func ownSet(arr []string) []string {
	// Для начала создаём карту с пустыми значениями
	mp := make(map[string]struct{})
	// И переменную для результата
	res := make([]string, 0)

	// Каждое слово в исходной последовательности
	// отправляется в карту
	for _, word := range arr {
		if _, ok := mp[word]; ok {
			continue // Если слово уже есть, пропускаем
		}
		mp[word] = struct{}{}
		res = append(res, word)
	}
	return res
}

func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Собственное множество: ", ownSet(set))
}
