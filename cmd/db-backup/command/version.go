package command

import (
	"fmt"
	"github.com/omegion/go-db-backup/pkg/info"

	"github.com/spf13/cobra"
)

// Version prints version/build.
func Version() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version/build number",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println(fmt.Sprintf("%s %s\n", info.AppName, info.Version))

			return nil
		},
	}

	return cmd
}
