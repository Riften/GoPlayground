package Interface

import "fmt"

type TInterface interface {
	PrintValue()
}

type tInterface struct {
	value string
}

func (t *tInterface) PrintValue() {
	fmt.Println(t.value)
}

func NewTInterface(str string) TInterface{
	return &tInterface{value: str}
}