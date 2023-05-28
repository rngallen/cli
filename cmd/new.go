/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/rngallen/mycli/data"
	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create a new note",
	Long:  "create a new note",
	Run: func(cmd *cobra.Command, args []string) {
		createNewNote()
	},
}

type promptContent struct {
	errorMsg string
	lable    string
}

func init() {
	noteCmd.AddCommand(newCmd)

}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) < 1 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	prompt := promptui.Prompt{
		Label:     pc.lable,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed : %v \n", err)
	}
	log.Printf("Input: %s \n", result)
	return result
}

func promptGetSelect(pc promptContent) string {

	items := []string{"animal", "food", "person", "object"}
	index := -1

	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.lable,
			Items:    items,
			AddLabel: "Other",
		}
		index, result, err = prompt.Run()
		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		log.Fatalf("Prompt faile %v \n", err)
	}
	log.Printf("Input: %s \n", result)
	return result
}

func createNewNote() {
	wordPromptContent := promptContent{
		"Please provide a word ",
		"What would you like to make a note of? ",
	}
	word := promptGetInput(wordPromptContent)

	definitionPromptContent := promptContent{
		"Please provide a definition ",
		fmt.Sprintf("What is the definition of %s? ", word),
	}
	definition := promptGetInput(definitionPromptContent)

	categoryPromptContent := promptContent{
		"Please provide a category ",
		fmt.Sprintf("What category does %s belong to ? ", word),
	}

	category := promptGetSelect(categoryPromptContent)

	note := data.StudyBuddy{Word: word, Definition: definition, Category: category}
	data.CreateNote(&note)
}
