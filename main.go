package main

import (
	"fmt"
	"os"
	"strconv"
)

func displayTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Задачи не найдены")
		return
	}

	for _, task := range tasks {
		statusEmoji := "📝"
		switch task.Status {
		case StatusInProgress:
			statusEmoji = "🔄"
		case StatusDone:
			statusEmoji = "✅"
		}

		fmt.Printf("%s ID: %d | %s\n", statusEmoji, task.ID, task.Title)
		fmt.Printf("   Описание: %s\n", task.Description)
		fmt.Printf("   Статус: %s | Создана: %s\n\n",
			task.Status, task.CreatedAt.Format("2006-01-02 15:04"))
	}
}

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

func createService(useMemoryStorage bool) (TaskService, error) {
	var repository TaskRepository
	repository = NewFileRepository("tasks.json")

	return NewTaskService(repository)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	useMemoryStorage := false
	if len(os.Args) > 1 && os.Args[1] == "--memory" {
		useMemoryStorage = true
		os.Args = append(os.Args[:1], os.Args[2:]...)
	}

	service, err := createService(useMemoryStorage)
	if err != nil {
		fmt.Printf("Ошибка инициализации сервиса: %v\n", err)
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Ошибка: недостаточно аргументов")
			fmt.Println("Использование: task-tracker add <заголовок> <описание>")
			os.Exit(1)
		}
		title := os.Args[2]
		description := os.Args[3]
		if err := service.Add(title, description); err != nil {
			fmt.Printf("Ошибка добавления задачи: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Задача успешно добавлена!")

	case "update":
		if len(os.Args) < 5 {
			fmt.Println("Ошибка: недостаточно аргументов")
			fmt.Println("Использование: task-tracker update <id> <заголовок> <описание>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Ошибка: неверный ID задачи: %v\n", err)
			os.Exit(1)
		}
		title := os.Args[3]
		description := os.Args[4]
		if err := service.Update(id, title, description); err != nil {
			fmt.Printf("Ошибка обновления задачи: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Задача успешно обновлена!")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка: недостаточно аргументов")
			fmt.Println("Использование: task-tracker delete <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Ошибка: неверный ID задачи: %v\n", err)
			os.Exit(1)
		}
		if err := service.Delete(id); err != nil {
			fmt.Printf("Ошибка удаления задачи: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Задача успешно удалена!")

	case "start":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка: недостаточно аргументов")
			fmt.Println("Использование: task-tracker start <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Ошибка: неверный ID задачи: %v\n", err)
			os.Exit(1)
		}
		if err := service.UpdateStatus(id, StatusInProgress); err != nil {
			fmt.Printf("Ошибка обновления статуса: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Задача помечена как 'в процессе'!")

	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Ошибка: недостаточно аргументов")
			fmt.Println("Использование: task-tracker done <id>")
			os.Exit(1)
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Printf("Ошибка: неверный ID задачи: %v\n", err)
			os.Exit(1)
		}
		if err := service.UpdateStatus(id, StatusDone); err != nil {
			fmt.Printf("Ошибка обновления статуса: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Задача помечена как 'выполнена'!")

	case "list":
		tasks := service.GetAll()
		displayTasks(tasks)

	case "list-todo":
		tasks := service.GetAllByStatus(StatusToDo)
		fmt.Println("=== Невыполненные задачи ===")
		displayTasks(tasks)

	case "list-progress":
		tasks := service.GetAllByStatus(StatusInProgress)
		fmt.Println("=== Задачи в процессе ===")
		displayTasks(tasks)

	case "list-done":
		tasks := service.GetAllByStatus(StatusDone)
		fmt.Println("=== Выполненные задачи ===")
		displayTasks(tasks)

	case "help", "-h", "--help":
		printUsage()

	default:
		fmt.Printf("Неизвестная команда: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}
