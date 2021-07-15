package scrapeit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Tests that IsFinalState returns false for an empty string
func TestIsFinalStateEmptyString(t *testing.T) {
	assert.Equal(t, false, isFinalState(""))
}
