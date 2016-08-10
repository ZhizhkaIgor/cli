package commands

import "github.com/jessevdk/go-flags"

type ExtendedCommander interface {
	flags.Commander
	Setup() error
}
