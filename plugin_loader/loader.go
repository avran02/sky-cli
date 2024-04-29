package pluginloader

import (
	"fmt"
	"os"
	"plugin"

	pluginInterface "github.com/avran02/sky-cli-lib"
)

const (
	pluginsRoot string = "/home/andrey/go/skybrary/sky-cli/dist/"
)

func LoadConf(pluginName string) pluginInterface.PluginConfiger {
	p := mustLoadPlugn(pluginName)
	fmt.Println(p)
	symbol, err := p.Lookup("GetPluginConfig") // search for PluginConfig object
	if err != nil {
		fmt.Println("can't load symbol:", err)
		fmt.Println("make sure your sky-cli plugin is valid. \nLooks like you have not have a PluginConfig object")
		os.Exit(1)
	}

	getConf, ok := symbol.(func() pluginInterface.PluginConfiger)
	if !ok {
		fmt.Println("make sure your sky-cli plugin is valid. \nLooks like your PluginConfig object doesn't implement PluginConfiger interface")
		os.Exit(1)
	}
	return getConf()
}

func mustLoadPlugn(pluginName string) *plugin.Plugin {
	pluginsAvailable := MustGetPluginNames()
	for _, p := range pluginsAvailable {
		if pluginName == p {
			plug, err := plugin.Open(pluginsRoot + p) // TODO: parcer plugin name
			if err != nil {
				fmt.Println("can't open plugin:", err)
				os.Exit(1)
			}
			return plug
		}
	}
	fmt.Println("Unknown plugin. Available plugins:")
	for _, pluginName := range pluginsAvailable {
		fmt.Println(pluginName)
	}
	os.Exit(1)
	return nil
}

func MustGetPluginNames() []string {
	dir, err := os.Open(pluginsRoot)
	if err != nil {
		fmt.Println("can't open plugins folder:", err)
		os.Exit(1)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}
