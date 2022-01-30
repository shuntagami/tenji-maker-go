package tenjibuilder_test

import (
	"testing"

	"github.com/shuntagami/tenji-maker-go/tenjibuilder"
	"github.com/stretchr/testify/assert"
)

func Test_A_HI_RU(t *testing.T) {
	actual := tenjibuilder.Build("A HI RU")
	expected := "o- o- oo\n-- o- -o\n-- oo --"
	assert.Equal(t, actual, expected)
}

func Test_KI_RI_N(t *testing.T) {
	actual := tenjibuilder.Build("KI RI N")
	expected := "o- o- --\no- oo -o\n-o -- oo"
	assert.Equal(t, actual, expected)
}

func Test_SI_U_MA_U_MA(t *testing.T) {
	actual := tenjibuilder.Build("SI MA U MA")
	expected := "o- o- oo o-\noo -o -- -o\n-o oo -- oo"
	assert.Equal(t, actual, expected)
}

func Test_NI_WA_TO_RI(t *testing.T) {
	actual := tenjibuilder.Build("NI WA TO RI")
	expected := "o- -- -o o-\no- -- oo oo\no- o- o- --"
	assert.Equal(t, actual, expected)
}

func Test_HI_YO_KO(t *testing.T) {
	actual := tenjibuilder.Build("HI YO KO")
	expected := "o- -o -o\no- -o o-\noo o- -o"
	assert.Equal(t, actual, expected)
}

func Test_KI_TU_NE(t *testing.T) {
	actual := tenjibuilder.Build("KI TU NE")
	expected := "o- oo oo\no- -o o-\n-o o- o-"
	assert.Equal(t, actual, expected)
}
