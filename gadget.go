package main

import (
	"awesomeProject/gadget"
)

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Царица", "Поломанные", "Time is running out"}
	playList(player, mixtape)
}
