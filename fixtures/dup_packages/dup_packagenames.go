package dup_packages // import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages"

import (
	"github.com/chintanparikh/counterfeiter/fixtures/dup_packages/a/foo"
	bfoo "github.com/chintanparikh/counterfeiter/fixtures/dup_packages/b/foo"
)

//go:generate counterfeiter . AB
type AB interface {
	A() foo.S
	foo.I
	B() bfoo.S
	bfoo.I
}
