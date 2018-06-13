package cmd

import (
	"io/ioutil"
	"os"

	"github.com/circleci/circleci-cli/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the CircleCI GraphQL API.",
	Run:   query,
}

func query(cmd *cobra.Command, args []string) {
	c := client.NewClient(viper.GetString("endpoint"), Logger)

	query, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		Logger.FatalOnError("Unable to read query", err)
	}

	resp, err := client.Run(c, viper.GetString("token"), string(query))
	if err != nil {
		Logger.FatalOnError("Error occurred when running query", err)
	}

	Logger.Prettyify(resp)
}
