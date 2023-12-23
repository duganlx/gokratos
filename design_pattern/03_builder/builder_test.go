package builder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuilder(t *testing.T) {
	var (
		builder  Builder
		director *Director
	)

	builder = &Builder1{}
	director = NewDirector(builder)
	director.Construct()
	res1 := builder.(*Builder1).GetResult()
	assert.Equal(t, "123", res1)

	builder = &Builder2{}
	director = NewDirector(builder)
	director.Construct()
	res2 := builder.(*Builder2).GetResult()
	assert.Equal(t, 6, res2)
}
