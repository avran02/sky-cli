# SKY_CLI tool for generating project templates

1. [Introduction](#introduction)
2. [Warnings](#warnings)
3. [Installation](#installation)
4. [Usage](#usage)
5. [Plugin development](#plugin-development)

## Introduction

sky-cli is a tool for generating your projects templates. It consists of three main parts:
- sky-cli-lib - a common library for plugins and the cli. It contains basic structures, their methods and generation logic.
- Plugins - are represented by .so files and should contain information about the project structure, OS commands and files templates
- sky-cli - a utility that realizes the logic of plugins connection and command line interface

Since the structure is described in the plugin, sky-cli is a very flexible tool and is not bind to a specific programming language or design pattern.
## Warnings
- Plugins require an installed go compiler for installation
- Not supported for windows. Maybe it will work in wsl
- Always generate the project in an empty directory to avoid losing files that have already been in the directory
- It is better to compile the plugin and sky-cli with the same compiler. The library used is very sensitive to version changes
## Installation
Install the go compiler

Install sky-cli 
```
go install github.com/avran02/sky-cli
```
To create a plugin directory in ~/.config/sky-cli and install the base plugins, run 
```
sky-cli configure
```
## Usage
To create a project template, the 'init' flag is used to specify the plugin to use for generation, for example to create a plugin template, run 
```
sky-cly init plugin
```
To get a list of all available plugins use 
```
sky-cli list
```
To install a plugin, execute.

 __WARNING: You _must_ specify the exact version of the plugin and cannot use the @latest__ tag
 This is because your link will be used to find the directory with the plugin for later compilation
```
sky-cli install your/plugin/url@v0.0.0
```

## Plugin development
If you don't find a suitable plugin, you can always write your own! The utility is designed so that anyone can write their own plugin without even knowing golang.

### ___Let's go!___

First we need to install sky-cli. Next, we need to generate a plugin template.
```
# create a folder with the future plugin.
mkdir my-perfect-plugin

# move to it
cd my-perfect-plugin

# generate the plugin template
sky-cli init plugin
```
The utility will ask for the project name for the `go mod init` command. If you will be using github to store the plugin, I recommend typing something like _github.com/avran02/my-perfect-plugin_. In our example, the users plugin will be named _my-perfect-plugin_. Keep this in mind when choosing a name, because no one wants to write a very long name 

Then open this folder in your favorite code editor. You will see the PluginConfig structure and some of its methods. You will also find the GetPluginConfig function. ___DON'T CHANGE IT!___ I prefer to describe all project directories in this method, and to describe commands and files in a separate var block at the bottom. All we have to do is write templates for the files in the tpl package, and describe the structure in a JSON-like object.

You can always check out an example [here](github.com/avran02/go-grpc).
