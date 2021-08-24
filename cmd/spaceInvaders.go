package main

import "github.com/juazasan/spaceInvaders/pkg/runtime"

func main() {
	spaceInvaders := runtime.NewSpaceInvadersRuntime()
	spaceInvaders.Start()
}
