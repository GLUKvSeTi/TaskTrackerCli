package main

import "fmt"

func printUsage() {
	fmt.Println(`Использование:
  task-tracker add <заголовок> <описание>    - Добавить задачу
  task-tracker update <id> <заголовок> <описание> - Обновить задачу
  task-tracker delete <id>                   - Удалить задачу
  task-tracker start <id>                    - Начать задачу (in-progress)
  task-tracker done <id>                     - Завершить задачу (done)
  task-tracker list                          - Показать все задачи
  task-tracker list-todo                     - Показать невыполненные задачи
  task-tracker list-progress                 - Показать задачи в процессе
  task-tracker list-done                     - Показать выполненные задачи
  task-tracker info                          - Показать информацию о хранилище
  task-tracker help                          - Показать эту справку`)
}

func main() {
	printUsage()
}
