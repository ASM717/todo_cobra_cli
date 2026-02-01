package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Показать список всех задач",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. Читаем файл
		fileData, err := os.ReadFile("tasks.json")
		if err != nil {
			fmt.Println("Список задач пуст или файл не найден.")
			return
		}

		// 2. Распаковываем JSON в слайс
		var tasks []Task
		err = json.Unmarshal(fileData, &tasks)
		if err != nil {
			fmt.Println("Ошибка при чтении данных:", err)
			return
		}

		// 3. Выводим заголовок
		fmt.Println("Ваши задачи:")

		// 4. Цикл по задачам
		// i - индекс (0, 1, 2...), task - копия элемента
		for i, task := range tasks {
			status := " "
			if task.Done {
				status = "x"
			}
			// Выводим в формате: 1. [ ] Текст задачи
			fmt.Printf("%d. [%s] %s\n", i+1, status, task.Text)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

/*
range:
Это самый частый способ перебора элементов в Go.
Он возвращает два значения: индекс и копию элемента.
_ (нижнее подчеркивание): Если бы нам не нужен был индекс i, мы бы написали for _, task := range tasks.
В Go нельзя оставлять неиспользуемые переменные — будет ошибка компиляции.
fmt.Printf: Позволяет форматировать строку. %d — для чисел, %s — для строк.
*/
