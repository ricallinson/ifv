# Interactive Fiction Voice



## Testing

	cd $GOPATH/src/git@github.com/ricallinson/ifv
	go test

## Coverage

	cd $GOPATH/src/git@github.com/ricallinson/ifv
	go test -covermode=count -coverprofile=count.out; go tool cover -html=count.out
