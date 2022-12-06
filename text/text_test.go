package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/width"
)

// var (
// 	// Contain Hiragana, Katakana, Kanji, spaces, alphabet, numeric, dot, and punctuation marks characters
// 	halfWidthStr      = "1, Vietnam.｢Betonamu｣､ベﾄﾅﾑ､越南｡"
// 	fullWidthStr      = "１，　Ｖｉｅｔｎａｍ．「Ｂｅｔｏｎａｍｕ」、ベトナム、越南。"
// 	canonicalWidthStr = "1, Vietnam.「Betonamu」、ベトナム、越南。"
// )

func TestTextPkg(t *testing.T) {
	assert.Equal(t, halfWidthStr, width.Narrow.String(fullWidthStr))
	assert.Equal(t, halfWidthStr, width.Narrow.String(canonicalWidthStr))
}
