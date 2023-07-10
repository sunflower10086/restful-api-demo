package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sunflower10086/restful-api-demo/version"
)

var (
	vers bool
)

var RootCmd = &cobra.Command{
	Use:     "demo-api",
	Long:    "demo API后端",
	Short:   "demo API后端",
	Example: "demo API后端 commands",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Printf("demo-api version %s\n", version.Version)
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "domo version")
}
