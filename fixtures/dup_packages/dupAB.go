package dup_packages // import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages"

//go:generate counterfeiter . DupAB
type DupAB interface {
	DupA
	DupB
}
