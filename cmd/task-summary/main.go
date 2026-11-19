package main

import "fmt"

func statusLabel(status string) string {
	switch status {
	case "new":
		return "Новая"
	case "in_progress":
		return "В работе"
	case "done":
		return "Завершена"
	default:
		return "Неизвестный статус"
	}

}

func countStatus(statuses []string, target string) int {
	count := 0
	for _, status := range statuses {
		if status == target {
			count++
		}
	}
	return count
}

func findStatus(statuses []string, target string) int {
	for i, status := range statuses {
		if status == target {
			return i
		}
	}
	return -1
}

func completionPercent(statuses []string) float64 {
	countAll := len(statuses)
	countDone := 0
	if countAll == 0 {
		return 0
	}
	for _, status := range statuses {
		if status == "done" {
			countDone++
		}
	}
	return 100 * float64(countDone) / float64(countAll)
}

func main() {
	statuses := []string{"new", "in_progress", "done", "done"}
	statuses = append(statuses, "new")

	fmt.Printf("Задач всего: %d\n", len(statuses))
	for i, status := range statuses {
		fmt.Printf("%d. %s\n", i+1, statusLabel(status))
	}
	fmt.Printf("Завершено: %d\n", countStatus(statuses, "done"))

	counts := make(map[string]int)
	for _, status := range statuses {
		counts[status]++
	}
	fmt.Printf("Индекс первого совпавшего статуса: %d\nДоля выполненных задач: %f\n", findStatus(statuses, "done"), completionPercent(statuses))
	fmt.Printf("Статистика: new=%d, in_progress=%d, done=%d", countStatus(statuses, "new"), countStatus(statuses, "in_progress"), countStatus(statuses, "done"))
}
