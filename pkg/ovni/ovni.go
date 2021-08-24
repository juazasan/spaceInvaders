package ovni

type Position struct {
	Row    int
	Column int
}

// Ovni is the interface all objects represented in the Space must implement
type Ovni interface {
	GetPosition() Position
	Render() string
}
