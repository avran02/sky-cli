// go build -buildmode plugin -o ../../dist/
package main

import (
	"fmt"
	"os"
	"path"

	"github.com/avran02/go-grpc/tpl"
	skyclilib "github.com/avran02/sky-cli-lib"
)

// must be named GetPluginConfig
func GetPluginConfig() skyclilib.PluginConfiger {
	return &PluginConfig{
		commands: []skyclilib.OsCommand{
			initGo,
			initCobraCli,
		},
		fs: &skyclilib.Folder{
			IsOptional: false,
			FolderStructure: skyclilib.FolderStructure{
				"db": skyclilib.Folder{
					IsOptional: true,
					FolderStructure: skyclilib.FolderStructure{
						"repo": skyclilib.Folder{
							IsOptional: false,
							FolderStructure: skyclilib.FolderStructure{
								"repo.go": gormRepo,
							},
						},
						"models": skyclilib.Folder{
							IsOptional:      false,
							FolderStructure: skyclilib.FolderStructure{},
						},
					},
				},
				"server": skyclilib.Folder{
					IsOptional: false,
					FolderStructure: skyclilib.FolderStructure{
						"server.go": server,
					},
				},
				"dockerfile":    dockerfile,
				".dockerignore": dockerignore,
				".golangci.yml": linter,
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

var (
	// files
	dockerfile = skyclilib.File{
		IsOptional: true,
		Tpl:        tpl.Dockerfile,
	}

	dockerignore = skyclilib.File{
		IsOptional: true,
		Tpl:        tpl.Dockerignore,
	}

	linter = skyclilib.File{
		IsOptional: true,
		Tpl:        tpl.Linter,
	}

	gormRepo = skyclilib.File{
		IsOptional: false,
		RequiredValues: map[string]string{
			"ProjectName": getCurrentDirName(),
		},
		Tpl: tpl.GormRepo,
	}

	server = skyclilib.File{
		IsOptional: false,
		UserValues: map[string]string{"ProtoPath": ""},
		Tpl:        tpl.Server,
	}

	// commands
	initGo = skyclilib.OsCommand{
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
				Name:            "module name",
				NeedGetFromUser: true,
			},
		},
	}

	initCobraCli = skyclilib.OsCommand{
		Name: "cobra-cli",
		Args: []skyclilib.CommandArg{
			{
				Name:            "init",
				NeedGetFromUser: false,
			},
		},
	}
)

func getCurrentDirName() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("can't get current dir")
		os.Exit(1)
	}
	return path.Base(dir)
}
