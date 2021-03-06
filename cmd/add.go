package cmd

import (
	"fmt"
	"log"

	"github.com/cuilan/todo/entity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var priority int

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo-cli",
	Long:  "Add will create a new todo-cli item to the list",
	Run:   addRun,
}

func addRun(cmd *cobra.Command, args []string) {
	items, err := entity.ReadItems(viper.GetString("datafile"))
	// items, err := entity.ReadItems(".todo.json")
	if err != nil {
		log.Printf("%v", err)
	}
	for _, x := range args {
		item := entity.Item{Text: x}
		item.SetPriority(priority)

		items = append(items, item)
	}

	if err := entity.SaveItems(viper.GetString("datafile"), items); err != nil {
		_ = fmt.Errorf("%v", err)
	}
}

func init() {
	RootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority,
		"priority", "p", 2, "Priority:1,2,3")
}
