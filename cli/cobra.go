package cli

import (
	"log"

	"github.com/spf13/cobra"
)

//NewCliCommand returns the root cmd. All flags related to cli configuration are declared here
func NewCliCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "avel [ACTION]",
		Short: "avel is a dummy cli",
		Long:  "avel is a dummy cli",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			err := setLogLevel(logLevel)
			if err != nil {
				log.Fatalf("Unable to set loglevel: %v", err)
			}
		},
	}
	cmd.PersistentFlags().StringVarP(&logLevel, "loglevel", "l", "info", "log level. For debugging purposes")
	return cmd, nil
}
