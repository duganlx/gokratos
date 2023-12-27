package decorator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExampleDecorator(t *testing.T) {
	var c Component = &ConcreteComponent{}
	var res int

	c = WrapAddDecorator(c, 10)
	res = c.Calc()
	assert.Equal(t, res, 10)

	c = WrapMulDecorator(c, 8)
	res = c.Calc()
	assert.Equal(t, res, 80)
}
