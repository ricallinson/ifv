# Interactive Fiction Voice

[![Build Status](https://travis-ci.org/ricallinson/ifv.svg?branch=master)](https://travis-ci.org/ricallinson/ifv) [![Build status](https://ci.appveyor.com/api/projects/status/me65m9g0ji5aj431/branch/master?svg=true)](https://ci.appveyor.com/project/ricallinson/ifv/branch/master)

## Install

You must first install the `Go` environment for your operating system from [golang.org](https://golang.org/dl/). Once installed open your terminal program.

* MAC: Terminal. To open the terminal, `apple+space` and type `terminal`.
* Windows: Command Prompt. To open the command prompt, `windows+r` and type `cmd`.
* Linux: If you don't know you shouldn't be using Linux.

You should now see your command line interface. From here you will install the program and execute it. Enter the following at the command line;

	go get github.com/ricallinson/ifv
	go install github.com/ricallinson/ifv
	ifv -story $GOPATH/src/github.com/ricallinson/ifv/fixtures/story.yaml

The game should now be running and you can play it by entering an option number and pressing enter.

## Usage

The game interface is currently text based. It will print out some of the story and when wait for a choice to be made. To make a choice enter the number of the option you want and then press the enter key. This will repeat until the game completes or is exited.

### story (required)

The path to the location of the story file to use for the game.

	ifv -story /path/to/story.yaml

## Story Development

A story is represented in a [YAML](http://yaml.org/) file. An example of which can be found [here](https://github.com/ricallinson/ifv/blob/master/fixtures/story.yaml).

Unfortunately there are no instructions for how the YAML file is constructed at the moment. It is reasonably straight forward and I encourage reverse engineering it as an interim until the manual is written.

You can make a copy of the YAML file and change it in a text editor. Once changed you'll need to reload the game program with `-story` being the path to your edited file.

## Coding Development Environment

	go get github.com/ricallinson/ifv
	go install github.com/ricallinson/ifv
	cd $GOPATH/src/github.com/ricallinson/ifv

### Testing

	cd $GOPATH/src/github.com/ricallinson/ifv
	go test

### Coverage

	cd $GOPATH/src/github.com/ricallinson/ifv
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out
