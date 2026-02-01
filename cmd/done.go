package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv" // Пакет для конвертации строк в числа

	"github.com/spf13/cobra"
)

var doneCmd = &cobra.Command{
	Use:   "done [номер задачи]",
	Short: "Отметить задачу как выполненную",
	Args:  cobra.ExactArgs(1), // Cobra проверит, что передан ровно 1 аргумент
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Конвертируем строку в число (индекс)
		idx, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("Ошибка: введите корректный номер задачи")
			return
		}

		// 2. Читаем файл
		fileData, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("Файл не найден")
			return
		}

		var tasks []Task
		json.Unmarshal(fileData, &tasks)

		// 3. Проверяем, существует ли такой индекс
		// Мы выводили список начиная с 1, поэтому отнимаем 1 для работы со слайсом
		realIdx := idx - 1
		if realIdx < 0 || realIdx >= len(tasks) {
			fmt.Println("Ошибка: задачи с таким номером нет")
			return
		}

		// 4. Меняем статус
		tasks[realIdx].Done = true

		// 5. Сохраняем всё назад
		data, _ := json.MarshalIndent(tasks, "", "  ")
		os.WriteFile("tasks.json", data, 0644)

		fmt.Printf("Задача '%s' выполнена! ✅\n", tasks[realIdx].Text)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
