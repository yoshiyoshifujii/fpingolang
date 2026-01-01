package gettingstarted

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_formatAbs(t *testing.T) {
	// given
	in := -42

	// when
	actual := formatAbs(in)

	// then
	assert.Equal(t, "The absolute value of -42 is 42", actual)
}
