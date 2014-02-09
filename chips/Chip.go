package chips

import (
	list "container/list"
	"errors"
	"fmt"
	"strconv"
)

var chipList *list.List
var chipCycles int
var chipSetName string
var verboseOutput bool = false
var toProcess *list.List
var readyToProcess *list.List

func Init(ChipSetName string) {
	chipSetName = ChipSetName
	chipList = list.New()
	readyToProcess = list.New()
}

func Reset() {
	readyToProcess = list.New()
	for e := chipList.Front(); e != nil; e = e.Next() {
		var c *chip = e.Value.(*chip)
		c.Reset()
	}
}

func VerboseOutput() {
	verboseOutput = true
}

func BriefOutput() {
	verboseOutput = false
}

func Process() {
	EchoChipHeader()
	chipCycles = 0

	toProcess = readyToProcess
	readyToProcess = list.New()
	for toProcess.Len() > 0 {
		chipCycles++
		if verboseOutput {
			fmt.Println(" Chip cycle started: " + strconv.Itoa(chipCycles))
		}

		for e := toProcess.Front(); e != nil; e = e.Next() {
			var c *chip = e.Value.(*chip)
			if !c.HasProcessed() {
				c.process()
			}
		}

		toProcess = readyToProcess
		readyToProcess = list.New()
	}
	EchoChipFooter()
}

type ChipOut struct {
	Name        string
	ConnectedTo *chip
	InputName   string
}

func EchoChipHeader() {
	fmt.Println("==================================================")
	fmt.Println(" Chip set: " + chipSetName)
	fmt.Println("==================================================")
}

func EchoChipFooter() {
	nandCount := 0
	fmt.Println("==================================================")
	for e := chipList.Front(); e != nil; e = e.Next() {
		var c *chip = e.Value.(*chip)
		if c.name == "Nand" {
			nandCount++
		}
	}
	fmt.Println(" Nand gates:", nandCount)
	fmt.Println(" total cycles:", chipCycles)
	fmt.Println("==================================================")
}

func NewChip(name string, outs ...interface{}) (*chip, error) {
	var c *chip = new(chip)

	if name == "Nand" {
		c.name = "Nand"
		c.SetInputs("a", "b")
	} else if name == "Out" {
		c.name = "Out"
		c.nickname = "out"
	} else if name == "In" {
		c.name = "In"
		c.nickname = "in"
	} else {
		return nil, errors.New("unimplemented chip '" + name + "'")
	}

	c.processed = false
	c.out = make([]*ChipOut, len(outs))
	index := 0
	for _, param := range outs {
		c.out[index] = &ChipOut{param.(ChipOut).Name, param.(ChipOut).ConnectedTo, param.(ChipOut).InputName}
		index++
	}

	chipList.PushFront(c)
	return c, nil
}

type chipIn struct {
	name  string
	value int
}

type chip struct {
	name      string
	nickname  string
	processed bool
	in        []*chipIn
	out       []*ChipOut
}

func (c *chip) Reset() {
	for _, inParam := range c.in {
		inParam.value = -1
	}

	c.processed = false
}

func (c *chip) SetNickname(nickname string) {
	c.nickname = nickname
}

func (c *chip) setName(name string) {
	c.name = name
}

func (c *chip) SetInputs(a ...interface{}) {
	c.in = make([]*chipIn, len(a))
	index := 0
	for _, param := range a {
		c.in[index] = &chipIn{param.(string), -1}
		index++
	}
}

func (c *chip) GetInput(name string) int {
	var result int = 0
	for _, inParam := range c.in {
		if name == inParam.name {
			result = inParam.value
			break
		}
	}
	return result
}

func (c *chip) SetInput(name string, value int) {
	for _, inParam := range c.in {
		if name == inParam.name {
			inParam.value = value
			if c.HasAllInput() {
				readyToProcess.PushFront(c)
			}
			break
		}
	}
}

func (c *chip) HasAllInput() bool {
	result := true
	for _, inParam := range c.in {
		if inParam.value == -1 {
			result = false
			break
		}
	}
	return result
}

func (c *chip) HasProcessed() bool {
	return c.processed
}

func (c *chip) EchoChip() {
	fmt.Println("name =", c.name)
	for _, inParam := range c.in {
		fmt.Println("   in={", inParam.name, ",", inParam.value, "}")
	}
	for _, outParam := range c.out {
		fmt.Println("   out={", outParam.Name, ",", outParam.ConnectedTo, "}")
	}
}

func (c *chip) setOutput(value int) {
	for _, outChip := range c.out {
		if outChip.ConnectedTo != nil {
			outChip.ConnectedTo.SetInput(outChip.Name, value)
		}
	}
}

func (c *chip) process() {
	if c.HasAllInput() {
		if verboseOutput {
			fmt.Println("   processing " + c.nickname)
		}
		switch c.name {
		case "Nand":
			if c.in[0].value == 0 && c.in[1].value == 0 {
				c.setOutput(1)
			} else {
				c.setOutput(0)
			}
			c.processed = true
		case "Out":
			fmt.Println(" Out: " + c.in[0].name + "=" + strconv.Itoa(c.in[0].value))
			c.processed = true
		case "In":
			c.processIn()
			c.processed = true
		}
	}
}

func (c *chip) processIn() {
	for _, outChip := range c.out {
		if outChip.ConnectedTo != nil {
			value := c.GetInput(outChip.InputName)
			outChip.ConnectedTo.SetInput(outChip.Name, value)
		}
	}
}
