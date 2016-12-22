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
