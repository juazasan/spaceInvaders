package bullet

import (
	"math/rand"
	"time"

	"github.com/juazasan/spaceInvaders/pkg/ovni"
)

type Bullet struct {
	position     ovni.Position
	maxRow       int
	maxColumn    int
	view         string
	speed        time.Duration // milliseconds to wait before next move
	lastMovement time.Time
}

func CreateBullet(firstPosition ovni.Position, maxRow int, maxColumn int, speed int) *Bullet {
	return &Bullet{
		position:     firstPosition,
		maxRow:       maxRow,
		maxColumn:    maxColumn,
		view:         "|",
		speed:        time.Duration(int64(speed) * int64(time.Millisecond)),
		lastMovement: time.Now(),
	}
}

func (a *Bullet) GetPosition() ovni.Position {
	return a.position
}

func (a *Bullet) UpdatePosition() {
	if time.Now().After(a.lastMovement.Add(a.speed)) {
		rand.Seed(time.Now().UnixNano())
		a.position.Row++
		if a.position.Row > a.maxRow {
			a.position.Row = a.maxRow
		}
		a.lastMovement = time.Now()
	}
}

func (a *Bullet) Render() string {
	return a.view
}
