package provisioner

import (
	"context"

	"github.com/urfave/cli"

	"github.com/smallstep/cli/command/ca/policy/policycontext"
	"github.com/smallstep/cli/command/ca/policy/ssh"
	"github.com/smallstep/cli/command/ca/policy/x509"
)

// Command returns the policy subcommand.
func Command(ctx context.Context) cli.Command {
	ctx = policycontext.NewContextWithProvisionerPolicyLevel(ctx)
	return cli.Command{
		Name:        "provisioner",
		Usage:       "manage certificate issuance policies for provisioners",
		UsageText:   "**step beta ca policy provisioner** <subcommand> [arguments] [global-flags] [subcommand-flags]",
		Description: `**step beta ca policy provisioner** command group provides facilities for managing certificate issuance policies for provisioners.`,
		Subcommands: cli.Commands{
			x509.Command(ctx),
			ssh.Command(ctx),
			removeCommand(ctx),
			viewCommand(ctx),
		},
	}
}
