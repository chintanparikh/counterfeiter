package foo // import "github.com/chintanparikh/counterfeiter/fixtures/dup_packages/b/foo"

type S struct{}

//go:generate counterfeiter . I
type I interface {
	FromB() S
}
