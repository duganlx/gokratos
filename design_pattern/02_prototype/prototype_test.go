package prototype

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var manager *PrototypeManager

type Type1 struct {
	name string
}

func (t *Type1) Clone() Cloneable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Cloneable {
	tc := *t
	return &tc
}

func init() {
	manager = NewPrototypeManager()

	t1 := &Type1{
		name: "type1",
	}

	t2 := &Type2{
		name: "type1",
	}
	manager.Set("t1", t1)
	manager.Set("t2", t2)
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")
	t2 := t1.Clone()
	t3 := manager.Get("t2")

	assert.NotSame(t, t1, t2)
	assert.Equal(t, t1, t2)
	assert.NotEqual(t, t1, t3)
}
