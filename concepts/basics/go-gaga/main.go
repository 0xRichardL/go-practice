package main

import (
	"github.com/y-bash/go-gaga"
)

var (
	// Contain Hiragana, Katakana, Kanji, spaces, alphabet, numeric, dot, and punctuation marks characters
	halfWidthStr      = "1, Vietnam.｢Betonamu｣､ﾍﾞﾄﾅﾑ､越南｡"
	fullWidthStr      = "１，　Ｖｉｅｔｎａｍ．「Ｂｅｔｏｎａｍｕ」、ベトナム、越南。"
	canonicalWidthStr = "1, Vietnam.「Betonamu」、ベトナム、越南。"
)

func main() {
	n, _ := gaga.Norm(gaga.Fold)

	println(n.String(halfWidthStr) == n.String(canonicalWidthStr))
	println(n.String(canonicalWidthStr) == n.String(fullWidthStr))
}
