package the_aliased_package // import "github.com/chintanparikh/counterfeiter/fixtures/aliased_package"

//go:generate counterfeiter . InAliasedPackage
type InAliasedPackage interface {
	Stuff(int) string
}
