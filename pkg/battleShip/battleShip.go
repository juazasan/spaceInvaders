package battleShip

import "github.com/juazasan/spaceInvaders/pkg/ovni"

type BattleShip struct {
	position  ovni.Position
	maxRow    int
	maxColumn int
	view      string
}

func (a *BattleShip) Create(firstPosition ovni.Position, maxRow int, maxColumn int) {
	a.position = firstPosition
	a.maxRow = maxRow
	a.maxColumn = maxColumn
	a.view = "^"
}

func (a *BattleShip) GetPosition() ovni.Position {
	return a.position
}

func (a *BattleShip) UpdatePosition() {
	newColumn := a.position.Column + 1
	if newColumn == a.maxColumn {
		newColumn = 0
	}
	a.position.Column = newColumn
}

func (a *BattleShip) Render() string {
	return a.view
}

func (a *BattleShip) MoveRight() {
	a.position.Column++
	if a.position.Column == a.maxColumn {
		a.position.Column = 0
	}
}

func (a *BattleShip) MoveLeft() {
	a.position.Column--
	if a.position.Column < 0 {
		a.position.Column = a.maxColumn - 1
	}
}
