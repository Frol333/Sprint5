package personaldata

// Personal - структура, содержащая персональные данные пользователя.
type Personal struct {
	Name   string  // Имя пользователя.
	Weight float64 // Вес пользователя в килограммах.
	Height float64 // Рост пользователя в сантиметрах.
}

// Print - метод структуры Personal, выводящий данные пользователя на экран.
func (p Personal) Print() {
	println("Имя:", p.Name)    // Вывод имени пользователя.
	println("Вес:", p.Weight)  // Вывод веса пользователя.
	println("Рост:", p.Height) // Вывод роста пользователя.
}
