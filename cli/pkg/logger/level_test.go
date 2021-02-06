package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsLevelActive(t *testing.T) {
	current := LvlDebug
	assert.True(t, IsLevelActive(current, LvlDebug))
	assert.True(t, IsLevelActive(current, LvlInfo))
	assert.True(t, IsLevelActive(current, LvlWarn))
	assert.True(t, IsLevelActive(current, LvlError))

	current = LvlInfo
	assert.False(t, IsLevelActive(current, LvlDebug))
	assert.True(t, IsLevelActive(current, LvlInfo))
	assert.True(t, IsLevelActive(current, LvlWarn))
	assert.True(t, IsLevelActive(current, LvlError))

	current = LvlWarn
	assert.False(t, IsLevelActive(current, LvlDebug))
	assert.False(t, IsLevelActive(current, LvlInfo))
	assert.True(t, IsLevelActive(current, LvlWarn))
	assert.True(t, IsLevelActive(current, LvlError))

	current = LvlError
	assert.False(t, IsLevelActive(current, LvlDebug))
	assert.False(t, IsLevelActive(current, LvlInfo))
	assert.False(t, IsLevelActive(current, LvlWarn))
	assert.True(t, IsLevelActive(current, LvlError))

}
