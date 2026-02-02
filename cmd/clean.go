/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Очистить список",
	Run: func(cmd *cobra.Command, args []string) {
		tasks := []Task{}

		fileData, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			fmt.Println("Ошибка при подготовке данных:", err)
			return
		}

		err = os.WriteFile("tasks.json", fileData, 0644)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}

		fmt.Println("Список задач успешно очищен!")
	},
}

func init() {
	rootCmd.AddCommand(cleanCmd)
}
