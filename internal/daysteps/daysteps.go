package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат строки данных")
	}

	ds.Steps, err = strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("ошибка при парсинге шагов: %w", err)
	}

	durationStr := parts[1]
	ds.Duration, err = time.ParseDuration(strings.ReplaceAll(durationStr, "h", "h"))
	if err != nil {
		return fmt.Errorf("ошибка при парсинге продолжительности: %w", err)
	}

	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <= 0 {
		return "", fmt.Errorf("продолжительность должна быть больше 0")
	}

	distance := float64(ds.Steps) * 0.0006 // 1 шаг = 0.0006 км
	calories, err := ds.Personal.WalkingSpentCalories(ds.Steps)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.", ds.Steps, distance, calories)
	return result, nil
}
