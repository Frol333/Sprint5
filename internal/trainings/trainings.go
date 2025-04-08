package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse разбирает строку данных и заполняет поля структуры Training.
func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("неверный формат данных")
	}

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("некорректное количество шагов: %w", err)
	}

	// Проверяем, что количество шагов не меньше или равно нулю.
	if t.Steps <= 0 {
		return fmt.Errorf("количество шагов должно быть больше нуля")
	}

	t.TrainingType = data[1]

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("некорректная длительность: %w", err)
	}

	// Проверяем, что длительность не меньше или равна нулю.
	if t.Duration <= 0 {
		return fmt.Errorf("длительность должна быть больше нуля")
	}

	return nil
}

// ActionInfo формирует строку с информацией о тренировке.
func (t Training) ActionInfo() (string, error) {
	// Вычисляем пройденную дистанцию
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)

	// Вычисляем среднюю скорость
	steps := 0 // тут steps не нужен, так как его нет в spentenergy.MeanSpeed
	speed := spentenergy.MeanSpeed(steps, distance, t.Duration)

	var calories float64

	// В зависимости от типа тренировки вычисляем количество сожженных калорий
	switch t.TrainingType {
	case "Бег":
		calories, _ = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	case "Ходьба":
		calories, _ = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
	default:
		// Возвращаем ошибку, если тип тренировки неизвестен
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	// Формируем строку с информацией о тренировке
	result := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories)

	return result, nil
}
