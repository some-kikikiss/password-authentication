package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"password-authentication/internal/generator"
)

var (
	length           uint
	withSymbols      bool
	withNumbers      bool
	withUpperLetters bool
	withEngLang      bool
	withRusLang      bool
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create password with your specifications  ",
	Long:  "Create password with your length and some specifications",
	Run:   generate,
}

func generate(_ *cobra.Command, args []string) {
	config := generator.Config{
		Length:           length,
		WithSymbols:      withSymbols,
		WithNumbers:      withNumbers,
		WithUpperLetters: withUpperLetters,
		WithEngLang:      withEngLang,
		WithRusLang:      withRusLang,
		Alphabet:         "",
	}
	g, err := generator.New(&config)
	if err != nil {
		log.Fatalf("Create generator troubles. Error : %s", err)
	}
	pass, errGen := g.Generate()
	if errGen != nil {
		log.Fatalf("Create password troubles. Error : %s", errGen)
	}
	fmt.Println(config.String())
	fmt.Println(*pass)
}

func init() {
	createCmd.PersistentFlags().UintVarP(&length, "length", "l", generator.DefaultConfig.Length,
		"length of the password")
	createCmd.PersistentFlags().BoolVarP(&withSymbols, "symbols", "s", generator.DefaultConfig.WithSymbols,
		"include symbols")
	createCmd.PersistentFlags().BoolVarP(&withNumbers, "numbers", "n", generator.DefaultConfig.WithNumbers,
		"include numbers")
	createCmd.PersistentFlags().BoolVarP(&withUpperLetters, "upper", "u", generator.DefaultConfig.WithUpperLetters,
		"include uppercase letters")
	createCmd.PersistentFlags().BoolVarP(&withEngLang, "eng", "e",
		generator.DefaultConfig.WithEngLang, "include English letters")
	createCmd.PersistentFlags().BoolVarP(&withRusLang,
		"ru", "r", generator.DefaultConfig.WithRusLang, "include Russian letters")

	rootCmd.AddCommand(createCmd)
}
