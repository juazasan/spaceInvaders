package runtime

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/juazasan/spaceInvaders/pkg/battleShip"
	"github.com/juazasan/spaceInvaders/pkg/bullet"
	"github.com/juazasan/spaceInvaders/pkg/keyboardManager"
	"github.com/juazasan/spaceInvaders/pkg/ovni"
	"github.com/juazasan/spaceInvaders/pkg/spaceInvader"
)

type SpaceInvadersRuntime struct {
	spaceInvaders []*spaceInvader.SpaceInvader
	battleShip    battleShip.BattleShip
	bullets       []*bullet.Bullet
}

// Game configuration
const maxRows = 30
const maxColumns = 40
const maxSpaceInvadersRows = 3

//const frames = 5

func NewSpaceInvadersRuntime() SpaceInvadersRuntime {
	var spaceInvadersRuntime SpaceInvadersRuntime
	var battleShip battleShip.BattleShip
	battleShip.Create(ovni.Position{Row: 0, Column: 2}, 0, maxColumns)
	spaceInvadersRuntime.battleShip = battleShip
	columnOffset := 1
	for j := 0; j < maxSpaceInvadersRows; j = j + 2 {
		for i := columnOffset; i < maxColumns-1; i = i + 2 {
			spaceInvadersRuntime.spaceInvaders = append(spaceInvadersRuntime.spaceInvaders, spaceInvader.CreateSpaceInvader(ovni.Position{Row: maxRows - 2 - j, Column: i}, maxRows, maxColumns, 2000))
		}
		if columnOffset == 1 {
			columnOffset = 2
		} else {
			columnOffset = 1
		}
	}
	return spaceInvadersRuntime
}

func (a *SpaceInvadersRuntime) update() {

	j := len(a.spaceInvaders)
	i := 0
	for i < j {
		if a.spaceInvaders[i].IsAlive() {
			a.spaceInvaders[i].UpdatePosition()
			i++
		} else {
			a.spaceInvaders[i] = a.spaceInvaders[j-1]
			a.spaceInvaders = a.spaceInvaders[:j-1]
			j--
		}
	}
	j = len(a.bullets)
	i = 0
	for i < j {
		a.bullets[i].UpdatePosition()
		if a.bullets[i].GetPosition().Row == maxRows {
			a.bullets[i] = a.bullets[j-1]
			a.bullets = a.bullets[:j-1]
			j--
			break
		}
		for _, spaceInvader := range a.spaceInvaders {
			if a.bullets[i].GetPosition().Row == spaceInvader.GetPosition().Row && a.bullets[i].GetPosition().Column == spaceInvader.GetPosition().Column {
				spaceInvader.Destroy()
				a.bullets[i] = a.bullets[j-1]
				a.bullets = a.bullets[:j-1]
				j--
				break
			}
		}
		i++
	}
}

func (a *SpaceInvadersRuntime) Start() {
	var k chan keyboard.Key = make(chan keyboard.Key)
	gameEscaped := false
	a.renderSpaceInvadersScreen()
	go keyboardManager.ListenToKeyboardInputs(k)
	go func() {
		for {
			time.Sleep(300 * time.Millisecond)
			a.update()
			a.renderSpaceInvadersScreen()
		}
	}()

	for !gameEscaped {
		key := <-k
		switch key {
		case keyboard.KeyEsc:
			gameEscaped = true
		case keyboard.KeyArrowLeft:
			a.battleShip.MoveLeft()
		case keyboard.KeyArrowRight:
			a.battleShip.MoveRight()
		case keyboard.KeySpace:
			firstPosition := a.battleShip.GetPosition()
			firstPosition.Row++
			a.bullets = append(a.bullets, bullet.CreateBullet(firstPosition, maxRows, firstPosition.Column, 300))
		}
	}
}

func (a *SpaceInvadersRuntime) renderSpaceInvadersScreen() {
	screen := new([maxRows][maxColumns]string)
	for i := 0; i < maxRows; i++ {
		for j := 0; j < maxColumns; j++ {
			screen[i][j] = " "
		}
	}
	battleShipPosition := a.battleShip.GetPosition()
	screen[battleShipPosition.Row][battleShipPosition.Column] = a.battleShip.Render()
	for _, spaceInvader := range a.spaceInvaders {
		spaceInvaderPosition := spaceInvader.GetPosition()
		screen[spaceInvaderPosition.Row][spaceInvaderPosition.Column] = spaceInvader.Render()
	}
	for _, bullet := range a.bullets {
		bulletPosition := bullet.GetPosition()
		screen[bulletPosition.Row][bulletPosition.Column] = bullet.Render()
	}
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
	newline := ""
	for j := maxRows - 1; j >= 0; j-- {
		for i := 0; i < maxColumns; i++ {
			newline = newline + screen[j][i]
		}
		fmt.Println(newline)
		newline = ""
	}
	fmt.Println("")
	fmt.Println("")
	fmt.Println("Press ESC to quit")
}
