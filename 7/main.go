package main

import (
	"fmt"
	"sync"
)

type clients struct {
	data map[int]string
	m    sync.Mutex
}

func (c *clients) write(id int, name string) {
	c.m.Lock()        //блокируем структуру с картой для избежания гонки горутин, для проверки наличия гонки используйте флаг -race: go run -race 7/main.go.
	c.data[id] = name //пишем данные в карту.
	fmt.Println(id)   //демонстрируем в терминал, порядок горутин.
	c.m.Unlock()      //разблокируем структуру с картой
}

func newClients() *clients {
	data := make(map[int]string)
	return &clients{
		data: data,
	}
}

func main() {
	data := newClients() //создаем экземпляр структуры с картой
	test := "test"
	var wg sync.WaitGroup //инициализируем группу ожидания
	wg.Add(10)            //увеличим счетчик группы ожидания на 10

	for i := 1; i < 11; i++ {
		id := i
		go func() {
			defer wg.Done()
			data.write(id, test)
		}() //запускаем 10 горутин, пишущих в карту в структуре
	}

	wg.Wait()              //Ждем выполнения всех горутин
	fmt.Println(data.data) //выводим карту в консоль
}

//можно реализовать без структуры и методов, однако я посчитал, что так удобнее и красивее :)