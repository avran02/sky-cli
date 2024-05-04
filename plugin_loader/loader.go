package pluginloader

import (
	"fmt"
	"os"
	"plugin"

	skyclilib "github.com/avran02/sky-cli-lib"
)

func getPluginsRoot() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return home + "/.config/sky-cli/plugins"
}

// Search for PluginConfig getter in plugin object, named "GetPluginConfig" and call it, returning PluginConfiger object
func LoadConf(pluginName string) skyclilib.PluginConfiger {
	p := mustLoadPlugn(pluginName)
	symbol, err := p.Lookup("GetPluginConfig") // search for PluginConfig getter
	if err != nil {
		fmt.Println("can't load symbol:", err)
		fmt.Println("make sure your sky-cli plugin is valid. \nLooks like you have not have a PluginConfig object")
		os.Exit(1)
	}

	getConf, ok := symbol.(func() skyclilib.PluginConfiger)
	if !ok {
		fmt.Println("make sure your sky-cli plugin is valid. \nLooks like your PluginConfig object doesn't implement PluginConfiger interface")
		os.Exit(1)
	}
	return getConf()
}

// Search for plugin in plugins folder + pluginName, try to load it. If error occurs, os.Exit(1)
func mustLoadPlugn(pluginName string) *plugin.Plugin {
	pluginsAvailable := MustGetPluginNames()
	for _, p := range pluginsAvailable {
		if pluginName == p {
			plug, err := plugin.Open(getPluginsRoot() + p) // TODO: parcer plugin name
			if err != nil {
				fmt.Println("can't open plugin:", err)
				os.Exit(1)
			}
			return plug
		}
	}
	fmt.Println("Unknown plugin. Available plugins:")
	for _, pluginName := range pluginsAvailable {
		fmt.Println(pluginName[:len(pluginName)-3])
	}
	os.Exit(1)
	return nil
}

// Search for all plugins in plugins folder
func MustGetPluginNames() []string {
	dir, err := os.Open(getPluginsRoot())
	if err != nil {
		fmt.Println("can't open plugins folder:", err)
		os.Exit(1)
	}
	defer dir.Close()

	files, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println("can't read plugins folder:", err)
		os.Exit(1)
	}
	fileNames := []string{}
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}
	return fileNames
}
