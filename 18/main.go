package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	n = 5
	t = 9
)

type counter struct { //структура счетчик.
	count int
	m     sync.Mutex
}

func newCounter() *counter { //функция инициализации структуры счетчика.
	return &counter{count: 0}
}

func (c *counter) increment() { //метод инкрементации счетчика в конкурнтной среде.
	c.m.Lock()
	c.count++
	c.m.Unlock()
}

func (c *counter) show() { //метод вывода значения счетчика.
	c.m.Lock()
	fmt.Println(c.count)
	c.m.Unlock()
}

func main() {
	counter := newCounter()                                                 //инициализируем структуру-счетчик.
	ctx, cancel := context.WithTimeout(context.Background(), t*time.Second) //инициализируем контекст с таймаутом по истечению жизни которого покажем значение счетчика.
	defer cancel()
	var wg sync.WaitGroup //инициализируем группу ожидания.
	wg.Add(n)             //увеличиваем счетчик группы ожидания на значение, равное количеству горутин.

	for i := 0; i < n; i++ { //инициализируем цикл с количеством итераций, равным количеству горутин.
		go worker(ctx, &wg, counter, i) //запускаем заданное количество горутин.
	}

	wg.Wait()      //ожидаем корректное завершение всех раболчих горутин.
	counter.show() //пишем в стандартный поток вывода значение структуры-счетчика.
}

func worker(ctx context.Context, wg *sync.WaitGroup, counter *counter, i int) {
	for { //в бесконечном цикле инкрементим структуру-счетчик каждые (i+1) секунды до закрытия контекста.
		select {
		case <-time.After(time.Duration(i+1) * time.Second):
			counter.increment()
		case <-ctx.Done():
			wg.Done() //уменьшаем счетчик группы ожидания на 1
			return
		}
	}
}
