# Interactive Fiction Voice

[![Build Status](https://travis-ci.org/ricallinson/ifv.svg?branch=master)](https://travis-ci.org/ricallinson/ifv) [![Build status](https://ci.appveyor.com/api/projects/status/me65m9g0ji5aj431/branch/master?svg=true)](https://ci.appveyor.com/project/ricallinson/ifv/branch/master)

## Testing

	cd $GOPATH/src/github.com/ricallinson/ifv
	go test

## Coverage

	cd $GOPATH/src/github.com/ricallinson/ifv
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out
