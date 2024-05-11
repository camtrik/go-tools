package cmd

import "github.com/spf13/cobra"

var wordCam = &cobra.Command{
	Use:   "word",
	Short: "Convert word format",
	Long:  "Convert word format in many ways",
	Run:   func(cmd *cobra.Command, args []string) {},
}
