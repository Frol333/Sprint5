package actioninfo

import (
	"fmt"
)

type DataParser interface {
	Parse(data string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// Перебираем все значения слайса dataset.
	for _, data := range dataset {
		// Пытаемся распарсить текущую строку данных.
		err := dp.Parse(data)
		if err != nil {
			// Если произошла ошибка парсинга, выводим ее и переходим к следующей итерации.
			fmt.Printf("Ошибка парсинга: %v\n", err)
			continue
		}
		// Формируем и выводим информацию об активности.
		fmt.Println(dp.ActionInfo())
	}
}
