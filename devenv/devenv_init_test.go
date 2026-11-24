package devenv

import "testing"

func TestInit(t *testing.T) {
	Init() // pass args through the command line
}

// go test -v -run TestInit -args -name=myapp -dir=./test/myapp
