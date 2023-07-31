package cmd

import (
	"github.com/electricbubble/gidevice-cli/internal"
	"github.com/spf13/cobra"
)

// rebootCmd represents the reboot command
var rebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot device",
	Run: func(cmd *cobra.Command, args []string) {
		udid, _ := cmd.Flags().GetString("udid")

		d, err := internal.GetDeviceFromCommand(udid)
		internal.ErrorExit(err)

		err = d.Reboot()
		internal.ErrorExit(err)
	},
}

func init() {
	rootCmd.AddCommand(rebootCmd)

	rebootCmd.Flags().StringP("udid", "u", "", "Device uuid")
}
