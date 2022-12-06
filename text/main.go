package main

import (
	"fmt"

	"golang.org/x/text/width"
)

var (
	// Contain Hiragana, Katakana, Kanji, spaces, alphabet, numeric, dot, and punctuation marks characters
	halfWidthStr      = "1, Vietnam.｢Betonamu｣､ベﾄﾅﾑ､越南｡"
	fullWidthStr      = "１，　Ｖｉｅｔｎａｍ．「Ｂｅｔｏｎａｍｕ」、ベトナム、越南。"
	canonicalWidthStr = "1, Vietnam.「Betonamu」、ベトナム、越南。"
)

func main() {
	fmt.Println(width.Widen.String("ｶﾞ"))
	fmt.Println(width.Narrow.String("ガ"))
}
