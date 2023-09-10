/* Наиболее эффективным способом конкатенации строк считается strings.Builder -- тип, не так давно добавленный в Go.
Это аналог bytes.Buffer, но он отличается тем, что при вызове метода String()
не происходит повторого выделения памяти и данные не копируются.

Ниже представлено применение этих двух типов.
*/

// НАДО ПРОПИСАТЬ БЕНЧМАРК И ПОКАЗАТЬ РЕЗУЛЬТАТЫ ИСПОЛЬЗОВАНИЯ ПАМЯТИ
// https://golangnote.com/golang/golang-stringsbuilder-vs-bytesbuffer //

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	// Использование более эффективного strings.Builder
	var strBuilder strings.Builder
	strBuilder.WriteString("Hello")
	strBuilder.WriteString(" ")
	strBuilder.WriteString("world!")
	fmt.Println(strBuilder.String()) // Hello world!

	// Использование bytes.Butter
	var buffer bytes.Buffer
	buffer.WriteString("Hello")
	buffer.WriteString(" ")
	buffer.WriteString("world!")
	fmt.Println(buffer.String()) // Hello world!
}
