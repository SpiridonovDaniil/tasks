package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	var n int
	_, err := fmt.Scan(&n) //получаем данные с консоли.
	if err != nil {
		log.Fatalln(err) //завершаем работу при ошибке получения данных с консоли.
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(n)*time.Second) //инициализируем контекст, считающийся исполненным по истечению заданного промежутка времени.
	defer cancel()                                                                         //высвобождаем ресурсы, в случае, если медленная операция завершается до истечения времени ожидания.
	ch := make(chan int)                                                                   //инициализируем канал для передачи данных.
	defer close(ch)                                                                        //закрытие канала по завершению работы.

	var wg sync.WaitGroup //создаем группу ожидания.
	wg.Add(1)             //добавляем к счетчику 1, так как читать канал будет только одна горутина.

	go func() { //запускаем горутину, читающую из канала.
		defer wg.Done() //по завершению работы горутины уменьшаем счетчик группы ожидания.
		for {
			select {
			case <-ctx.Done(): //при получении данных о исполненности контекста, завершаем работу горутины.
				return
			case <-time.After(1 * time.Second): //в цикле раз в секунду(для наглядности) читаем данные из канала.
				fmt.Println(<-ch)
			}
		}
	}()

	for {
		select {
		case <-ctx.Done(): //получаем информацию о исполненности контекста.
			fmt.Println(ctx.Err()) //выводим в консоль причину исполненности контекста.
			wg.Wait()              //ожидаем завершения работы читающей горутины.
			return                 //завершаем функцию main.
		case ch <- n - 1: //в цикле пишем в канал данные.
			n -= 1 //для реализации счетчика, имитирующего время до исполнения контекста(для наглядности).
		}
	}
}