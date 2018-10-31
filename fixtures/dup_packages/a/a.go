package a // import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages/a"

import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages/a/foo"

//go:generate counterfeiter . A
type A interface {
	V1() foo.I
}
