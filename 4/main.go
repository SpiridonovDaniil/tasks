package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {

	ch := make(chan int)                                                  //инициализируем канал.
	defer close(ch)                                                       //отложенная функция закрытия канала.
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt) //инициализируем контекст, канал получения нужного сигнала. Контекст считается выполненным при получении сигнала, вызове возвращаемой функции, либо закрытии родительского канала.
	defer stop()                                                          //не забываем закрыть канал получения сигналов, если контекст уже закрыт.
	var num int
	_, err := fmt.Scan(&num) //получаем данные от пользователя.
	if err != nil {
		log.Fatalln(err) //завершаем работу при ошибке получения данных от пользователя.
	}

	var wg sync.WaitGroup //инициализируем группу ожидания.
	wg.Add(num)           //устанавливаем счетчик в количестве, указанном пользователем.

	for i := 0; i < num; i++ { //запускаем столько горутин, сколько захочет пользователь.
		go func() {
			defer wg.Done() //по завершению работы горутины вычитаем счетчик группы ожидания.
			worker(ctx, ch) //функция worker бесконечно читает канал до получения данных о закрытии контекста.
		}()
	}
	for {
		select {
		case <-ctx.Done(): //кейс при закрытии контекста.
			fmt.Println(ctx.Err())
			stop()    //как можно скорее прекращаем получать уведомления о сигналах.
			wg.Wait() //ждем завершения работы всех горутин.
			return    //завершаем работу функции main.
		case ch <- rand.Intn(100): //бесконечно пишем в канал.
		}
	}

}

func worker(ctx context.Context, ch chan int) {
	for {
		select {
		case <-ctx.Done(): //при получении данных(закрытый канал) о закрытии контекста, завершаем работу функции worker.
			return
		case <-time.After(1 * time.Second): //задержка для наглядности количества горутин и избежания ужаса в терминале. :)
			fmt.Println(<-ch) //читаем с канала.
		}
	}
}

//Выбор способа завершения всех горутин обоснован простотой и лаконичностью кода. Не используются сторонние библиотеки.
//Группа ожидания дает возможность и время правильно завершить работу(освободить ресурсы, в том числе соединения с базами
//данных и удаленными сервисами(если есть) обработать уже полученные запросы, не потеряв данные, не принимать новые данные).
//Контекст дает возможность оповестить о необходимости завершения работы всех горутин сразу.