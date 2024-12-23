package gadget

import (
	"fmt"
)

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Играть", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Остановить")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Играть", song)
}

func (t TapeRecorder) Record() {
	fmt.Println(("Записать"))
}

func (t TapeRecorder) Stop() {
	fmt.Println("Остановить")
}
