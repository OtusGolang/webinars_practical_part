package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "clncnd",
	Short: "CleanCalendar is a calendar micorservice demo",
}

func init() {
	RootCmd.AddCommand(GrpcServerCmd)
	RootCmd.AddCommand(GrpcClientCmd)
}
