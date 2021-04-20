package app

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/lsytj0413/myinctl/cmd/app/command/version"
)

func NewMyInCtlCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:  "myinctl",
		Long: "",
		PersistentPreRunE: func(*cobra.Command, []string) error {
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
		// Args: func(cmd *cobra.Command, args []string) error {
		// 	return nil
		// },
	}

	cmd.AddCommand(version.NewVersionCommand())
	return cmd
}
