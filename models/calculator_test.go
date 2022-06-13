package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should return calculator with 0 value", func(t *testing.T) {
		calculator := NewCalculator()
		assert.Equal(t, 0.0, calculator.value)
	})

	t.Run("should implement Calculator", func(t *testing.T) {
		calc := NewCalculator()
		assert.Implements(t, new(Calculator), calc)
	})
}
