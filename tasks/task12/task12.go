// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

// Собственным множеством будет слайс строк
func ownSet(arr []string) []string {
	// Для начала создаём мапу с пустыми значениями
	mp := make(map[string]struct{})
	// И переменную для результата
	res := make([]string, 0)

	// Каждое слово в исходной последовательности
	// отправляется в карту
	for _, word := range arr {
		mp[word] = struct{}{}
	}
	// Затем проходим по ключам карты
	// и добавляем их в результат, таким образом
	// исключая повторы
	for key := range mp {
		res = append(res, key)
	}
	return res
}

func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("Собственное множество: ", ownSet(set))
}
