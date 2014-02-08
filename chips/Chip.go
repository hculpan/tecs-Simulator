package chips

import (
	"fmt"
)

type ChipInInterface struct {
	name  string
	value int
}

type ChipOutInterface struct {
	name        string
	connectedTo *Chip
}

type Chip struct {
	name string
	in   []*ChipInInterface
	out  []*ChipOutInterface
}

func (c *Chip) SetName(name string) {
	c.name = name
}

func (c *Chip) SetInputs(a ...interface{}) {
	c.in = make([]*ChipInInterface, len(a))
	index := 0
	for _, param := range a {
		c.in[index] = &ChipInInterface{param.(string), -1}
		index++
	}
}

func (c *Chip) SetInput(name string, value int) {
	for _, inParam := range c.in {
		if name == inParam.name {
			inParam.value = value
			break
		}
	}
}

func (c *Chip) SetOutputs(a ...interface{}) {
	c.out = make([]*ChipOutInterface, len(a))
	index := 0
	for index < len(a) {
		name := a[index].(string)
		chip := a[index+1].(Chip)
		c.out[index] = &ChipOutInterface{name, &chip}
		index += 2
	}
}

func (c *Chip) HasAllInput() bool {
	result := true
	for _, inParam := range c.in {
		if inParam.value == -1 {
			result = false
			break
		}
	}
	return result
}

func (c *Chip) EchoChip() {
	fmt.Println("name =", c.name)
	for _, inParam := range c.in {
		fmt.Println("   in={", inParam.name, ",", inParam.value, "}")
	}
	for _, outParam := range c.out {
		fmt.Println("   out={", outParam.name, ",", outParam.connectedTo, "}")
	}
}

func (c *Chip) Process(a ...interface{}) {
	fmt.Println(len(a))
	for _, param := range a {
		fmt.Println(param)
	}
}
