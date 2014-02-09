/**
 * Created with IntelliJ IDEA.
 * User: harry
 * Date: 2/5/14
 * Time: 1:26 PM
 */
package main

import (
	//"fmt"
	chips "github.com/hculpan/tecs-Simulator/chips"
	"github.com/jimlawless/cfg"
	"log"
	//	"os"
	"os/user"
	"path/filepath"
)

func main() {
	configOptions := &ConfigOptions{}

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	mymap := make(map[string]string)
	propFilename := filepath.Join(usr.HomeDir, ".tecsim")
	err = cfg.Load(propFilename, mymap)
	if err != nil {
		log.Fatal(err)
	}

	if val, ok := mymap["installed-chips"]; ok {
		configOptions.installed_chips = val
	} else {
		configOptions.installed_chips = filepath.Join(usr.HomeDir, "installed-chips")
	}

	chips.Init("And gate")
	chips.VerboseOutput()

	output, _ := chips.NewChip("Out")
	output.SetInputs("a")

	nand, _ := chips.NewChip("Nand", chips.ChipOut{Name: "a", ConnectedTo: output})
	nand.SetNickname("nand")

	not1, _ := chips.NewChip("Nand", chips.ChipOut{Name: "a", ConnectedTo: nand})
	not1.SetNickname("not1")

	not2, _ := chips.NewChip("Nand", chips.ChipOut{Name: "b", ConnectedTo: nand})
	not2.SetNickname("not2")

	input, _ := chips.NewChip("In",
		chips.ChipOut{"a", not1, "a"}, chips.ChipOut{"b", not1, "a"},
		chips.ChipOut{"a", not2, "b"}, chips.ChipOut{"b", not2, "b"})
	input.SetInputs("a", "b")
	input.SetInput("a", 1)
	input.SetInput("b", 1)

	chips.Process()

	chips.Reset()

	input.SetInput("a", 0)
	input.SetInput("b", 0)

	chips.Process()

	chips.Reset()

	input.SetInput("a", 1)
	input.SetInput("b", 0)

	chips.Process()

	chips.Reset()

	input.SetInput("a", 0)
	input.SetInput("b", 1)

	chips.Process()

	chips.Reset()

	input.SetInput("a", 1)
	input.SetInput("b", 1)

	chips.Process()
}
