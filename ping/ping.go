package ping

import (
	"github.com/spf13/cobra"
	"github.com/yakumo-saki/phantasma-flow-cli/client"
)

func Ping(cmd *cobra.Command, args []string) {
	cli := client.NewClient(":5000")
	cli.TestConn(":5000")
}
