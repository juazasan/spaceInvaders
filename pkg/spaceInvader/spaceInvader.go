package spaceInvader

import (
	"math/rand"
	"time"

	"github.com/juazasan/spaceInvaders/pkg/ovni"
)

type SpaceInvader struct {
	position     ovni.Position
	maxRow       int
	maxColumn    int
	view         string
	validRows    [3]int
	validColumns [3]int
	speed        time.Duration // milliseconds to wait before next move
	lastMovement time.Time
	isAlive      bool
}

func CreateSpaceInvader(firstPosition ovni.Position, maxRow int, maxColumn int, speed int) *SpaceInvader {
	return &SpaceInvader{
		position:     firstPosition,
		maxRow:       maxRow,
		maxColumn:    maxColumn,
		view:         "ø",
		validRows:    [3]int{firstPosition.Row - 1, firstPosition.Row, firstPosition.Row + 1},
		validColumns: [3]int{firstPosition.Column - 1, firstPosition.Column, firstPosition.Column + 1},
		speed:        time.Duration(int64(speed) * int64(time.Millisecond)),
		lastMovement: time.Now(),
		isAlive:      true,
	}
}

func (a *SpaceInvader) GetPosition() ovni.Position {
	return a.position
}

func (a *SpaceInvader) UpdatePosition() {
	if a.isAlive && time.Now().After(a.lastMovement.Add(a.speed)) {
		rand.Seed(time.Now().UnixNano())
		newRow := a.validRows[rand.Intn(2)]
		newColumn := a.validColumns[rand.Intn(2)]
		a.position = ovni.Position{Row: newRow, Column: newColumn}
		a.lastMovement = time.Now()
	}
}

func (a *SpaceInvader) Render() string {
	return a.view
}

func (a *SpaceInvader) Destroy() {
	a.view = "X"
	a.isAlive = false
}

func (a *SpaceInvader) IsAlive() bool {
	return a.isAlive
}
