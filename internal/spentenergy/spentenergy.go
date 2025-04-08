package spentenergy

import (
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
	if duration <= 0 {
		return 0
	}
	distance := Distance(steps, height)
	return distance / duration.Hours()
}

// WalkingSpentCalories вычисляет калории, потраченные при ходьбе.
func WalkingSpentCalories(steps int, weight float64, height float64, duration time.Duration) float64 {
	if weight <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	meanSpeed := MeanSpeed(steps, height, duration)
	return ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Minutes()
}

// RunningSpentCalories вычисляет калории, потраченные при беге.
func RunningSpentCalories(steps int, weight float64, height float64, duration time.Duration) float64 {
	if weight <= 0 || duration <= 0 {
		return 0
	}
	meanSpeed := MeanSpeed(steps, height, duration)

	return (runningCaloriesMeanSpeedMultiplier*meanSpeed - runningCaloriesMeanSpeedShift) * weight
}
