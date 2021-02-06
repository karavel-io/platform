package predicate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringTrue(t *testing.T) {
	assert.True(t, StringTrue("whatever"))
}

func TestStringFalse(t *testing.T) {
	assert.False(t, StringFalse("whatever"))
}

func TestStringOr(t *testing.T) {
	// true || true
	p := StringOr(StringTrue, StringTrue)
	assert.True(t, p("whatever"))

	// true || false
	p = StringOr(StringTrue, StringFalse)
	assert.True(t, p("whatever"))

	// false || true
	p = StringOr(StringFalse, StringTrue)
	assert.True(t, p("whatever"))

	// false || false
	p = StringOr(StringFalse, StringFalse)
	assert.False(t, p("whatever"))
}

func TestIsStringInSlice(t *testing.T) {
	haystack := []string{"a", "b", "c"}

	assert.True(t, IsStringInSlice(haystack)("a"))
	assert.True(t, IsStringInSlice(haystack)("b"))
	assert.True(t, IsStringInSlice(haystack)("c"))
	assert.False(t, IsStringInSlice(haystack)("d"))
}

func TestStringHasPrefix(t *testing.T) {
	p := StringHasPrefix("test/")

	assert.True(t, p("test/test"))
	assert.False(t, p("test-test"))
	assert.False(t, p("whatever"))
}
