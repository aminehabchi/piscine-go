package main

import "github.com/01-edu/z01"

type Door struct {
	state int
}

func PrintStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func CloseDoor(Door *Door) bool {
	PrintStr("Door Closing...")
	Door.state = CLOSE
	return true
}

func OpenDoor(Door *Door) {
	PrintStr("Door Opening...")
	Door.state = OPEN
}

func IsDoorOpen(Door *Door) bool {
	PrintStr("is the Door opened ?")
	Door.state = OPEN
	return true
}

func IsDoorClose(ptrDoor *Door) bool {
	PrintStr("is the Door closed ?")
	return true
}

func main() {
	door := &Door{}

	OpenDoor(door)
	if IsDoorClose(door) != true {
		OpenDoor(door)
	}
	if IsDoorOpen(door) != true {
		CloseDoor(door)
	}
	if door.state == OPEN {
		CloseDoor(door)
	}
}

const (
	OPEN  = 1
	CLOSE = 0
)
