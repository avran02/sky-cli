// go build -buildmode plugin -o ../../dist/
package main

import (
	skyclilib "github.com/avran02/sky-cli-lib"
	"github.com/avran02/sky-cli/plugins/plugin/tpl"
)

// must be named GetPluginConfig
func GetPluginConfig() skyclilib.PluginConfiger {
	return &PluginConfig{
		commands: []skyclilib.OsCommand{
			{
				Name: "go",
				Args: []skyclilib.CommandArg{
					{
						Name:            "mod",
						NeedGetFromUser: false,
					},
					{
						Name:            "init",
						NeedGetFromUser: false,
					},
					{
						Name:            "Project name",
						NeedGetFromUser: true,
					},
				},
			},
		},
		fs: &skyclilib.Folder{
			IsOptional: false,
			FolderStructure: map[string]interface{}{
				"main.go": skyclilib.File{
					IsOptional:     false,
					RequiredValues: map[string]string{},
					UserValues:     map[string]string{},
					Tpl:            tpl.Main(),
				},
				"tpl": skyclilib.Folder{
					IsOptional: false,
				},
			},
		},
	}
}

// must be named PluginConfig
type PluginConfig struct {
	commands []skyclilib.OsCommand
	fs       *skyclilib.Folder
}

// Return slice of os commands that will be executed before generation
func (p *PluginConfig) GetOsCommands() []skyclilib.OsCommand {
	return p.commands
}

// Return JSON-like structure describing virtual file system
func (p *PluginConfig) GetVirtualFs() *skyclilib.Folder {
	return p.fs
}
