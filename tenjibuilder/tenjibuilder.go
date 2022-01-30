package tenjibuilder

import (
	"strings"

	"github.com/shuntagami/tenji-maker-go/helper"
	"github.com/shuntagami/tenji-maker-go/tenji"
)

// 受け取った複数のかな文字を点字に変換します
func Build(text string) string {
	var matrix [][]string
	var array []string

	// textをスペースで分割
	arr := strings.Split(text, " ")
	for _, s := range arr {
		matrix = append(matrix, tenji.CreateSingleUnit(s))
	}
	for _, arr := range helper.Transpose(matrix) {
		array = append(array, strings.Join(arr, " "))
	}
	return strings.Join(array, "\n")
}
