package main

import (
	"fmt"

	"golang.design/x/hotkey"
)

func strToKey(in string) (hotkey.Key, error) {
	switch in {
	case "0":
		return hotkey.Key0, nil
	case "1":
		return hotkey.Key1, nil
	case "2":
		return hotkey.Key2, nil
	case "3":
		return hotkey.Key3, nil
	case "4":
		return hotkey.Key4, nil
	case "5":
		return hotkey.Key5, nil
	case "6":
		return hotkey.Key6, nil
	case "7":
		return hotkey.Key7, nil
	case "8":
		return hotkey.Key8, nil
	case "9":
		return hotkey.Key9, nil
	case "a":
		return hotkey.KeyA, nil
	case "b":
		return hotkey.KeyB, nil
	case "c":
		return hotkey.KeyC, nil
	case "d":
		return hotkey.KeyD, nil
	case "e":
		return hotkey.KeyE, nil
	case "f":
		return hotkey.KeyF, nil
	case "g":
		return hotkey.KeyG, nil
	case "h":
		return hotkey.KeyH, nil
	case "i":
		return hotkey.KeyI, nil
	case "j":
		return hotkey.KeyJ, nil
	case "k":
		return hotkey.KeyK, nil
	case "l":
		return hotkey.KeyL, nil
	case "m":
		return hotkey.KeyM, nil
	case "n":
		return hotkey.KeyN, nil
	case "o":
		return hotkey.KeyO, nil
	case "p":
		return hotkey.KeyP, nil
	case "q":
		return hotkey.KeyQ, nil
	case "r":
		return hotkey.KeyR, nil
	case "s":
		return hotkey.KeyS, nil
	case "t":
		return hotkey.KeyT, nil
	case "u":
		return hotkey.KeyU, nil
	case "v":
		return hotkey.KeyV, nil
	case "w":
		return hotkey.KeyW, nil
	case "x":
		return hotkey.KeyX, nil
	case "y":
		return hotkey.KeyY, nil
	case "z":
		return hotkey.KeyZ, nil
	default:
		return 0, fmt.Errorf("%v is not a valid hotkey", in)
	}
}

func strToMod(in string) (hotkey.Modifier, error) {
	switch in {
	case "win":
		return hotkey.ModWin, nil
	case "ctrl":
		return hotkey.ModCtrl, nil
	case "alt":
		return hotkey.ModAlt, nil
	case "shift":
		return hotkey.ModShift, nil
	default:
		return 0, fmt.Errorf("%v is not a valid modifier key", in)
	}
}
