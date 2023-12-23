package factorymethod

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestOperator(t *testing.T) {
	var (
		factory OperatorFactory
		res     int
	)

	factory = PlusOperatorFactory{}
	res = compute(factory, 1, 2)
	assert.Equal(t, 3, res)

	factory = MinusOperatorFactory{}
	res = compute(factory, 4, 2)
	assert.Equal(t, 2, res)
}
