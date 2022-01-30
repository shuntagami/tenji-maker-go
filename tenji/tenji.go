// tenji package はかな一文字から点字を生成できます
package tenji

import (
	"fmt"
	"strings"
)

/*
Tenji はかな一文字を生成するのに必要な情報を保持します
vowel: 母音
consonant: 子音
unit: 点字のマス目。点字は6個の点で表されます
*/
type Tenji struct {
	vowel     string
	consonant string
	unit      []string
}

// 与えられた1かな文字から点字を生成します
func CreateSingleUnit(text string) []string {
	tenji := &Tenji{}
	switch {
	case len(text) == 1:
		if strings.ContainsAny(text[0:1], "AIUEO") {
			tenji.vowel = text[0:1]
		}
		if text == "N" {
			tenji.consonant = "N"
		}
	case len(text) == 2:
		tenji.consonant = text[0:1]
		tenji.vowel = text[1:2]
	}
	replaced := tenji.replaceRomanToTenji()
	tenji.unit = append(tenji.unit, replaced[0:2], replaced[2:4], replaced[4:6])
	return tenji.unit
}

// ローマ字を6bitの2進数に変換した後,0と1をそれぞれ点字の"-"と"o"に変換します
func (c *Tenji) replaceRomanToTenji() string {
	return strings.ReplaceAll(
		strings.ReplaceAll(fmt.Sprintf("%06b", c.toByte()), "0", "-"),
		"1", "o",
	)
}

// 母音と子音の対応する点を6bitの2進数に変換し両者の論理和を返します
func (c *Tenji) toByte() byte {
	// 「ん」は特にルールがないのでただの固定値とみなす
	if c.consonant == "N" && c.vowel == "" {
		return byte(0b000111)
	}
	if c.consonant == "Y" && (c.vowel == "A" || c.vowel == "U") {
		return ConsonantMap[c.consonant] | VowelMap[c.vowel]>>4
	}
	if c.consonant == "Y" && c.vowel == "O" {
		return ConsonantMap[c.consonant] | VowelMap[c.vowel]>>2
	}
	if c.consonant == "W" && c.vowel == "A" {
		return VowelMap[c.vowel] >> 4
	}
	return ConsonantMap[c.consonant] | VowelMap[c.vowel]
}
