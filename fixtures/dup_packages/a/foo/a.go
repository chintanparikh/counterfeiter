package foo // import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages/a/foo"

type S struct{}

//go:generate counterfeiter . I
type I interface {
	FromA() S
}
