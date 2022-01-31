package tenjibuilder

import (
	"strings"

	"github.com/shuntagami/tenji-maker-go/helper"
	"github.com/shuntagami/tenji-maker-go/tenji"
)

// 受け取った複数のかな文字を点字に変換します
func Build(text string) string {
	var matrix [][]string
	var sl []string

	// textをスペースで分割
	chars := strings.Split(text, " ")
	for _, char := range chars {
		matrix = append(matrix, tenji.CreateSingleUnit(char))
	}
	for _, s := range helper.Transpose(matrix) {
		sl = append(sl, strings.Join(s, " "))
	}
	return strings.Join(sl, "\n")
}
