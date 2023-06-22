package components

import "github.com/yohamta/donburi"

type InputData struct {
	Left    bool
	Right   bool
	Up      bool
	Down    bool
	ButtonA bool
}

var Input = donburi.NewComponentType[InputData]()

func (i *InputData) Reset() {
	i.Left = false
	i.Right = false
	i.Up = false
	i.Down = false
	i.ButtonA = false
}
