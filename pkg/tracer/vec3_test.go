package tracer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVecNegation(t *testing.T) {
	v := Vec3{1, -2, 3}
	n := v.Neg()

	assert.Equal(t, n, Vec3{-1, 2, -3})
}

func TestVecAddition(t *testing.T) {
	v := Vec3{1, 2, 3}
	w := Vec3{2, 3, 4}

	assert.Equal(t, v.Add(w), Vec3{3, 5, 7})
}

func TestVecSubtraction(t *testing.T) {
	v := Vec3{1, 2, 3}
	w := Vec3{3, 2, 1}

	assert.Equal(t, v.Sub(w), Vec3{-2, 0, 2})
}

func TestVecMultiplication(t *testing.T) {
	v := Vec3{1, 2, 3}

	t.Run("by a scalar", func(t *testing.T) {
		assert.Equal(t, v.Mul(2), Vec3{2, 4, 6})
		assert.Equal(t, v.Mul(0.5), Vec3{0.5, 1, 1.5})
	})

	t.Run("by another Vec", func(t *testing.T) {
		w := Vec3{0.5, 3, -7}

		assert.Equal(t, v.Prod(w), Vec3{0.5, 6, -21})
	})
}

func TestVecDivision(t *testing.T) {
	v := Vec3{1, 2, 3}

	assert.Equal(t, v.Div(2), Vec3{0.5, 1, 1.5})
}

func TestDotProduct(t *testing.T) {
	// a := NewVector(1, 2, 3)
	// b := NewVector(2, 3, 4)
	a := Vec3{1, 2, 3}
	b := Vec3{2, 3, 4}

	assert.Equal(t, a.Dot(b), 20.0)
}

func TestCrossProduct(t *testing.T) {
	a := Vec3{1, 2, 3}
	b := Vec3{2, 3, 4}

	assert.Equal(t, a.Cross(b), Vec3{-1, 2, -1})
	assert.Equal(t, b.Cross(a), Vec3{1, -2, 1})
}
