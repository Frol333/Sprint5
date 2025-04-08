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
	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(distance, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
