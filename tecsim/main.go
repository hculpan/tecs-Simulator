/**
 * Created with IntelliJ IDEA.
 * User: harry
 * Date: 2/5/14
 * Time: 1:26 PM
 */
package main

import (
	"fmt"
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

	//	fmt.Printf("%v\n", mymap)
	//pwd, _ := os.Getwd()
	//}" "	fmt.Println(pwd)

	chips.Init("Not gate")

	output, _ := chips.NewChip("out")

	// not
	chip, err := chips.NewChip("Nand", chips.ChipOut{"a", output})
	if err != nil {
		fmt.Println("Unable to create chip:", err.Error())
		return
	}

	input, _ := chips.NewChip("in", chips.ChipOut{"a", chip}, chips.ChipOut{"b", chip})

	input.SetInput("a", 0)

	chips.Process()
}
