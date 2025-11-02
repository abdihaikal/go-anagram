// main.go
package main

import (
	"fmt"

	"github.com/abdihaikal/go-anagram/anagram"
)

func main() {
	fmt.Println(anagram.IsAnagram("listen", "silent"))         // true
	fmt.Println(anagram.IsAnagram("Debit Card", "Bad Credit")) // true
	fmt.Println(anagram.IsAnagram("hello", "world"))           // false
}
