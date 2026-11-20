package main

import "fmt"

type Task struct {
	ID       int
	Title    string
	Status   string
	Priority int
}

func (t Task) Summary() string {
	// TODO: верните "#ID Название [статус], приоритет N".
	return fmt.Sprintf("#%d %s [%s], приоритет %d", t.ID, t.Title, t.Status, t.Priority)
}

func (t *Task) Complete() {
	t.Status = "done"
	// TODO: присвойте статус done исходной задаче.
}

func (t *Task) SetPriority(priority int) bool {
	if priority < 1 || priority > 3 {
		return false
	}
	t.Priority = priority
	return true
	// TODO: проверьте диапазон 1–3, при ошибке объект не менять.
}

func tasksByStatus(tasks []Task, status string) []Task {
	needTask := []Task{}
	for _, task := range tasks {
		if task.Status == status {
			needTask = append(needTask, task)
		}
	}
	return needTask
}

func findTaskByID(tasks []Task, id int) *Task {
	for i := range tasks {
		if tasks[i].ID == id {
			return &tasks[i]
		}
	}
	// TODO: верните &tasks[i] при совпадении, иначе nil.
	return nil
}

func main() {
	tasks := []Task{
		{ID: 1, Title: "Установить Go", Status: "done", Priority: 1},
		{ID: 2, Title: "Освоить функции", Status: "in_progress", Priority: 2},
		{ID: 3, Title: "Изучить структуры", Status: "new", Priority: 3},
	}
	// TODO: найдите ID=2, завершите и установите приоритет 3.
	// TODO: посчитайте завершённые задачи.
	found := findTaskByID(tasks, 2)
	if found != nil {
		found.Complete()
		found.SetPriority(3)
	}

	count := 0
	for _, task := range tasks {
		if task.Status == "done" {
			count++
		}
	}

	for _, task := range tasks {
		fmt.Println(task.Summary())
	}

	fmt.Printf("Задач завершено: %d\n", count)

	if findTaskByID(tasks, 99) == nil {
		fmt.Println("Задача не найдена")
	}

	fmt.Print(tasksByStatus(tasks, "done"))
}
