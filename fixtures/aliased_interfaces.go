package fixtures

import alias "github.com/chintanparikh/counterfeiter/fixtures/another_package"

//go:generate counterfeiter . AliasedInterface

// AliasedInterface is an interface that embeds an interface in an aliased package.
type AliasedInterface interface {
	alias.AnotherInterface
}
