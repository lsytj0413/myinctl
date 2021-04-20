package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

func NewVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:          "version",
		Short:        "Print the version of myinctl",
		SilenceUsage: true,
		Run: func(c *cobra.Command, args []string) {
			fmt.Println("0.0.1")
		},
	}
}
