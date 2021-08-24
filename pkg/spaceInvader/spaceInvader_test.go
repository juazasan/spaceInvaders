package spaceInvader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpaceInvader(t *testing.T) {
	t.Run("test create spaceInvader with valid first position", func(t *testing.T) {})

	t.Run("test create spaceInvader with invalid first position", func(t *testing.T) {})

	t.Run("test update position", func(t *testing.T) {
		// setup
		spaceInvader := createSpaceInvader(Position{Row: 5, Column: 1}, 10, 10)
		currentPosition := spaceInvader.GetPosition()
		spaceInvader.UpdatePosition()
		newPosition := spaceInvader.GetPosition()
		assert.NotEqual(t, currentPosition, newPosition)
		assert.Equal(t, currentPosition.Column+1, newPosition.Column)
	})

	t.Run("test render", func(t *testing.T) {
		// setup
		spaceInvader := createSpaceInvader(Position{Row: 5, Column: 1}, 10, 10)

		// act

		// assert
		assert.Equal(t, "ø", spaceInvader.Render(), "spaceInvader format invalid")
	})

}
