package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// Task описывает структуру задачи
type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить новую задачу",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Ошибка: введите текст задачи")
			return
		}

		// Объединяем аргументы в одну строку, если ввели без кавычек
		taskText := strings.Join(args, " ")

		// 1. Создаем экземпляр структуры
		newTask := Task{
			Text: taskText,
			Done: false,
		}

		// 2. Создаем переменную под список задач
		var tasks []Task
		// 3. Читаем файл. Если файла нет — это не ошибка, просто идем дальше
		fileData, err := os.ReadFile("tasks.json")
		if err == nil {
			// Если файл есть, "распаковываем" JSON в наш слайс
			json.Unmarshal(fileData, &tasks)
		}

		tasks = append(tasks, newTask)

		// 4. Кодируем ВЕСЬ список обратно в JSON
		data, err := json.MarshalIndent(tasks, "", "  ")
		if err != nil {
			fmt.Println("Ошибка кодирования:", err)
			return
		}

		// 3. Сохраняем в файл (пока что просто перезаписываем файл tasks.json)
		err = os.WriteFile("tasks.json", data, 0644)
		if err != nil {
			fmt.Println("Ошибка записи в файл:", err)
			return
		}

		fmt.Printf("Добавлено! Теперь у тебя задач в списке: %d\n", len(tasks))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
