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
	input := 1000000.255
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "1,000,000.26 บาท", output)
}

func Test_ThaiCurrencyFormat_minus1000000254_return_minus100000025(t *testing.T) {
	input := -1000000.254
	output := ThaiCurrencyFormat(input)
	assert.Equal(t, "-1,000,000.25 บาท", output)
}
