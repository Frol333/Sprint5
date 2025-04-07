package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Frol333/Sprint5/tree/main/internal/spentenergy"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")
	if len(data) != 3 {
		return fmt.Errorf("invalid data format")
	}

	t.Steps, err = strconv.Atoi(data[0])
	if err != nil {
		return fmt.Errorf("invalid steps: %w", err)
	}

	t.TrainingType = data[1]

	t.Duration, err = time.ParseDuration(data[2])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	if t.Duration <= 0 {
		return "", fmt.Errorf("duration must be greater than zero")
	}

	distance := spentenergy.Distance(t.Steps, t.Personal.Height)
	speed := spentenergy.MeanSpeed(distance, t.Duration)

	var calories float64
	var err error

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.Running(t.Steps, t.Personal.Weight, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.Walking(t.Steps, t.Personal.Weight, t.Duration)
	default:
		return "неизвестный тип тренировки", fmt.Errorf("unknown training type: %s", t.TrainingType)
	}

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f",
		t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
