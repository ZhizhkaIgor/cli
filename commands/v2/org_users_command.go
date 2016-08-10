package v2

import (
	"os"

	"code.cloudfoundry.org/cli/cf/cmd"
	"code.cloudfoundry.org/cli/commands/flags"
)

type OrgUsersCommand struct {
	RequiredArgs flags.Organization `positional-args:"yes"`
	AllUsers     bool               `short:"a" description:"List all users in the org"`
	usage        interface{}        `usage:"CF_NAME org-users ORG"`
}

func (_ OrgUsersCommand) Setup() error {
	return nil
}

func (_ OrgUsersCommand) Execute(args []string) error {
	cmd.Main(os.Getenv("CF_TRACE"), os.Args)
	return nil
}
