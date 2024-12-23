package main

import (
	"awesomeProject/gadget"
)

type Player interface {
	Play(string)
	Stop()
}

func TryOut(player Player) {
	player.Play("Тестируемый трек")
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
	player.Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	TryOut(gadget.TapeRecorder{})
	TryOut(gadget.TapePlayer{})
}
