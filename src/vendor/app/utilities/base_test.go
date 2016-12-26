package utilities

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func Test_PadLeft_12345678_return_0012345678(t *testing.T) {
	input := "12345678"
	output := PadLeft(input, "0", 10)
	assert.Equal(t,"0012345678", output)
}

func Test_PadLeft_123456789_return_0123456789(t *testing.T) {
	input := "123456789"
	output := PadLeft(input, "0", 10)
	assert.Equal(t, "0123456789", output)
}

func Test_PadLeft_1234567890_return_1234567890(t *testing.T) {
	input := "1234567890"
	output := PadLeft(input, "0", 10)
	assert.Equal(t, "1234567890", output)
}

func Test_ThaiCurrencyFormat_025_return_025(t *testing.T) {
	input := 0.25
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "0.25 บาท", output)
}

func Test_ThaiCurrencyFormat_1000_return_1000(t *testing.T) {
	input := 1000.00
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "1,000.00 บาท", output)
}

func Test_ThaiCurrencyFormat_1000000_return_1000000(t *testing.T) {
	input := 1000000.00
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "1,000,000.00 บาท", output)
}

func Test_ThaiCurrencyFormat_100000025_return_100000025(t *testing.T) {
	input := 1000000.25
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "1,000,000.25 บาท", output)
}

func Test_ThaiCurrencyFormat_1000000255_return_100000026(t *testing.T) {
	input := 1000000.255
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "1,000,000.26 บาท", output)
}

func Test_ThaiCurrencyFormat_1000000254_return_100000025(t *testing.T) {
	input := 1000000.254
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "1,000,000.25 บาท", output)
}

func Test_ThaiCurrencyFormat_minus1000000254_return_minus100000025(t *testing.T) {
	input := -1000000.254
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "-1,000,000.25 บาท", output)
}

func Test_Substring_start0_LengthIs1_return_firstChar(t *testing.T) {
	param1 := "ทดสอบตัดสตริง"
	param2 := 0
	param3 := 1
	output := Substring(param1, param2, param3)
	assert.Equal(t,"ท", output)
}

func Test_Substring_start0_LengthInNegative_return_allString(t *testing.T) {
	param1 := "ทดสอบตัดสตริง"
	param2 := 0
	param3 := -1
	output := Substring(param1, param2, param3)
	assert.Equal(t,"ทดสอบตัดสตริง", output)
}

func Test_Substring_start0_LengthIs5_return_first5Char(t *testing.T) {
	param1 := "ทดสอบตัดสตริง"
	param2 := 0
	param3 := 5
	output := Substring(param1, param2, param3)
	assert.Equal(t,"ทดสอบ", output)
}

func Test_Substring_start0_LengthIsGreaterThanString_return_allString(t *testing.T) {
	param1 := "ทดสอบตัดสตริง"
	param2 := 0
	param3 := len([]rune(param1)) + 3
	output := Substring(param1, param2, param3)
	assert.Equal(t,"ทดสอบตัดสตริง", output)
}

func Test_Substring_startInNegative_LengthIs2_return_first2Char(t *testing.T) {
	param1 := "ทดสอบตัดสตริง"
	param2 := -3
	param3 := 2
	output := Substring(param1, param2, param3)
	assert.Equal(t,"ทด", output)
}