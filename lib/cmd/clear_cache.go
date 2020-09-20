package cmd

import (
	"github.com/niccoloCastelli/gohls/lib/hls"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(clearCacheCmd)
}

var clearCacheCmd = &cobra.Command{
	Use:   "clear-cache",
	Short: "Clears all caches and temporary files",
	Long:  `Clears all caches and temporary files`,
	Run: func(cmd *cobra.Command, args []string) {
		hls.ClearCache()
	},
}
