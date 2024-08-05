package main

import (
	"fmt"

	"github.com/tanukiops/leetcode/easy"
)

func main() {
	result := easy.KthDistinct([]string{"a", "a", "aaa", "b", "b", "bb"}, 1)
	fmt.Println(result)
}
