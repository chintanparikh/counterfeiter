package dup_packages // import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages"

import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages/b/foo"

//go:generate counterfeiter . DupB
type DupB interface {
	B() foo.S
}
