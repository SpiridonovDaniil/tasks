package main

import (
	"fmt"
	"sync"
)

func main() {
	m := [5]int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup   //Создаем группу ожидания (счетчик)
	wg.Add(len(m))          //указываем количество ожидаемых горутин
	ch := make(chan int, 5) //инициализируем канал для получения данных из горутин
	//буферизированный, чтобы горутины не ждали, пока из них считают информацию

	for i := 0; i < len(m); i++ {
		//запускаем горутины
		go func(ch chan int, count int) {
			defer wg.Done()     //уменьшаем счетчик группы ожидания по окончанию работы горутин
			ch <- count * count //пишем результат в канал
		}(ch, m[i])
	}

	wg.Wait() //ждем окончания выполнения всех горутин(счетчик равен 0)

	close(ch) //закрываем канал, чтобы при чтении горутина мейн не ожидала больше значений из канала
	var res int
	for value := range ch {
		res += value //читаем результаты из канала
	}

	fmt.Println(res) //выводим результат в консоль
}

// TODO попробывать оптимизировать код
