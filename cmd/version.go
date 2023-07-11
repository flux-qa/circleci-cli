package cmd

import (
	"fmt"

	"github.com/CircleCI-Public/circleci-cli/cmd/create_telemetry"
	"github.com/CircleCI-Public/circleci-cli/settings"
	"github.com/CircleCI-Public/circleci-cli/telemetry"
	"github.com/CircleCI-Public/circleci-cli/version"
	"github.com/spf13/cobra"
)

type versionOptions struct {
	cfg  *settings.Config
	args []string
}

func newVersionCommand(config *settings.Config) *cobra.Command {
	opts := versionOptions{
		cfg: config,
	}
	var telemetryClient telemetry.Client

	return &cobra.Command{
		Use:   "version",
		Short: "Display version information",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			telemetryClient = create_telemetry.CreateTelemetry(config)
			defer telemetryClient.Close()
			_ = telemetryClient.Track(telemetry.CreateVersionEvent(version.Version))

			opts.cfg.SkipUpdateCheck = true
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			opts.args = args
		},
		Run: func(_ *cobra.Command, _ []string) {
			fmt.Printf("%s+%s (%s)\n", version.Version, version.Commit, version.PackageManager())
		},
	}
}
