package keyboardManager

import (
	"github.com/eiannone/keyboard"
)

func ListenToKeyboardInputs(k chan keyboard.Key) {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		//fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		k <- key

	}
}
