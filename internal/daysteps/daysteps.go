package daysteps

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Yandex-Practicum/go1fl-4-sprint-final/internal/spentcalories"
)

var (
	StepLength = 0.65 // длина шага в метрах
)

func parsePackage(data string) (int, time.Duration, error) {
	var parts []string
	current := ""

	for _, ch := range data {
		if ch == ',' {
			parts = append(parts, current)
			current = ""
			continue
		}
		current += string(ch)
	}
	parts = append(parts, current) // добавим последний элемент

	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неправильное количество аргументов")
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil || steps <= 0 {
		return 0, 0, fmt.Errorf("невалидное количество шагов: %v", err)
	}

	dur, err := time.ParseDuration(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("ошибка при парсинге продолжительности: %v", err)
	}

	return steps, dur, nil
}

// DayActionInfo обрабатывает входящий пакет, который передаётся в
// виде строки в параметре data. Параметр storage содержит пакеты за текущий день.
// Если время пакета относится к новым суткам, storage предварительно
// очищается.
// Если пакет валидный, он добавляется в слайс storage, который возвращает
// функция. Если пакет невалидный, storage возвращается без изменений.
func DayActionInfo(data string, weight, height float64) string {
	steps, duration, err := parsePackage(data)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return ""
	}

	if steps <= 0 {
		return ""
	}

	// Вычисляем дистанцию в километрах
	distance := float64(steps) * StepLength / 1000

	// Вычисляем потраченные калории (используем функцию из пакета spentcalories)
	calories := spentcalories.WalkingSpentCalories(steps, weight, height, duration)

	// Формируем строку результата
	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", steps, distance, calories)
	return result
}
