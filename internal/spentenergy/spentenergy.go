package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                              = 1000  // количество метров в километре.
	minInH                             = 60    // количество минут в часе.
	stepLengthCoefficient              = 0.45  // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesWeightMultiplier    = 0.035 // коэффициент для расчета калорий при ходьбе, зависящий от веса.
	walkingSpeedHeightMultiplier       = 0.029 // коэффициент для расчета калорий при ходьбе, зависящий от скорости и роста.
	runningCaloriesMeanSpeedMultiplier = 0.035 // коэффициент для расчета калорий при беге, зависящий от скорости.
	runningCaloriesMeanSpeedShift      = 0.029 // коэффициент для расчета калорий при беге, сдвиг.
)

// Distance вычисляет пройденное расстояние в километрах.
func Distance(steps int, height float64) float64 {
	lenStep := height * stepLengthCoefficient
	return float64(steps) * lenStep / mInKm
}

// MeanSpeed вычисляет среднюю скорость в км/ч.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 0 {
		return 0
	}
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

func main() {
	var height float64 = 0.5
	var duration time.Duration = time.Hour * 2
	var steps int = 1000

	speed := MeanSpeed(steps, height, duration)
	println(speed)
}

// WalkingSpentCalories вычисляет калории, потраченные при ходьбе.
func WalkingSpentCalories(steps int, weight float64, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 || height <= 0 || duration <= 0 || steps < 0 {
		return 0, errors.New("некорректные входные параметры")
	}

	// Здесь должна быть логика расчета калорий.
	// Временная реализация для примера.
	meanSpeed := MeanSpeed(steps, height, duration)
	walkingCaloriesWeightMultiplier := 0.035
	walkingSpeedHeightMultiplier := 0.029
	calories := ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Minutes()
	return calories, nil
}

// RunningSpentCalories вычисляет калории, потраченные при беге.
func RunningSpentCalories(steps int, weight float64, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 || duration <= 0 || steps < 0 || height <= 0 {
		return 0, errors.New("некорректные входные параметры")
	}
	meanSpeed := MeanSpeed(steps, height, duration)

	calories := (runningCaloriesMeanSpeedMultiplier*meanSpeed + runningCaloriesMeanSpeedShift) * weight * duration.Minutes()

	return calories, nil
}
