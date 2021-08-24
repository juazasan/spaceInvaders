build:
	go build ./pkg/ovni
	go build ./pkg/spaceInvader
	go build ./pkg/battleShip
	go build ./pkg/keyboardManager
	go build ./pkg/runtime
	go build ./cmd/spaceInvaders.go

test:
	go test ./pkg/runtime

run: build
	go run ./cmd/spaceInvaders.go