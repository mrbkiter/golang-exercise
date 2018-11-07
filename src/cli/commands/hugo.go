package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo Short description",
	Long: `Hogo Longer description.. 
            feel free to use a few lines here.
            `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		fmt.Println(strings.Join(args, " "))
	},
}
