package cmd

import (
	"log"
	"strings"

	"github.com/camtrik/go-tools/internal/word"
	"github.com/spf13/cobra"
)

const (
	ModeUpper = iota + 1
	ModeLower
	ModeUnderscoreToUpperCamelCase
	ModeUnderScoreToLowerCamelCase
	ModeCammelCaseToUnderscore
)

var mode int8
var str string

var desc = strings.Join([]string{
	"This subcommand supports various word format conversions, with the following modes:",
	"1: Convert all to uppercase",
	"2: Convert all to lowercase",
	"3: Convert snake_case to UpperCamelCase",
	"4: Convert snake_case to lowerCamelCase",
	"5: Convert CamelCase to snake_case",
}, "\n")

var wordCam = &cobra.Command{
	Use:   "word",
	Short: "Convert word format",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModeLower:
			content = word.ToLower(str)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderScoreToUpperCamelCase(str)
		case ModeUnderScoreToLowerCamelCase:
			content = word.UnderScoreToLowerCamelCase(str)
		case ModeCammelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(str)
		default:
			log.Fatalf("Unsupported mode, please refer to the help documentation")
		}
		log.Printf("Output: %s", content)
	},
}

func init() {
	wordCam.Flags().StringVarP(&str, "str", "s", "", "Please input the word you want to convert")
	wordCam.Flags().Int8VarP(&mode, "mode", "m", 0, "Please select the conversion mode")
}
