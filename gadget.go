package gadget

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapePlayer) Stop() {
	fmt.Println("Stopping")
}

type TapeRecorder struct {
	Microfones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}
func (t TapeRecorder) Stop() {
	fmt.Println("Stopping")
}
func (t TapeRecorder) Recorder() {
	fmt.Println("Recording")
}
