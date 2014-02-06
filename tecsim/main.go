/**
 * Created with IntelliJ IDEA.
 * User: harry
 * Date: 2/5/14
 * Time: 1:26 PM
 * To change this template use File | Settings | File Templates.
 */
package main

 import (
 	"fmt"
 	"os"
 	"os/user"
 	"path/filepath"
 	"log"
 	"github.com/jimlawless/cfg"
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

	fmt.Printf("%v\n", mymap)
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
}

