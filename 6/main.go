package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	stop := make(chan struct{})                             //инициализируем канал для остановки горутины, пустая структура, потому что она ничего не весит.
	defer close(stop)                                       //закрываем канал при завершении работы.
	ctx, cancel := context.WithCancel(context.Background()) //инициализируем контекст с вохможностью закрыть его.
	var wg sync.WaitGroup                                   //инициализируем группу ожидания.
	wg.Add(4)                                               //увеличиваем счетчик группы ожидания на количество горутин.

	go sigInChan(stop, &wg)   //запускаем горутину с прекращением работы через сигнал в канале.
	go closeContext(ctx, &wg) //запускаем горутину с прекращением работы через исполнение контекста.
	go returnExample(&wg)     //запускаем горутину, которая завершит работу по окончанию операций.
	go func() {
		defer wg.Done() //при завершении работы горутины вычитаем счетчик группы ожидания.
		for {
			runtimeGoExit()
			//TODO work
		}
	}() //запускаем горутину с прекращением работы функцией runtime.GoExit(не затрагивает другие горутины, выполняет отложенные функции)

	time.Sleep(6 * time.Second) //имитируем работу.
	stop <- struct{}{}          //сигналом в канал останавливаем горутину sigInChan.
	cancel()                    //считаем контекст исполненным, останавливая горутину closeContext.
	//fmt.Println("main exit")    //раскомментировать для примера использования Goexit в главной горутине, остальные горутины должны продолжать работу,
	//runtime.Goexit()			  //но приложение завершится Exit(2) fatal error: no goroutines (main called runtime.Goexit) - deadlock!
	wg.Wait()           //ждем завершения всех горутин.
	fmt.Println("Exit") //уведомляем о завершении работы функции main.
}

func sigInChan(stop <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done() //при завершении работы горутины вычитаем счетчик группы ожидания.
	for {
		select {
		default: //в цикле каждые две секунды имитируем работу.
			time.Sleep(2 * time.Second)
			fmt.Println("sigInChan work")
		case <-stop: //при получении сигнала в канале завершаем работу горутины.
			fmt.Println("sigInChan stop")
			return
		}

	}
}

func closeContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done() //при завершении работы горутины вычитаем счетчик группы ожидания.
	for {
		select {
		default: //в цикле каждые три секунды имитируем работу.
			time.Sleep(3 * time.Second)
			fmt.Println("closeContext work")
		case <-ctx.Done(): //при получении сигнала о закрытии контекста завершаем работу горутины.
			fmt.Println("closeContext stop")
			return
		}
	}
}

func returnExample(wg *sync.WaitGroup) {
	defer wg.Done() //при завершении работы горутины вычитаем счетчик группы ожидания.
	var i int       //инициализируем счетчик.
	for {
		time.Sleep(4 * time.Second) //в цикле каждые четыре секунды имитируем работу.
		fmt.Println("returnExample work")
		i++         //инкрементим счетчик.
		if i == 2 { //заданного значение завершаем работу горутины.
			fmt.Println("returnExample stop")
			return
		}
	}
}

func runtimeGoExit() {
	time.Sleep(4 * time.Second) //четыре секунды имитируем работу.
	fmt.Println("runtimeGoExit work")
	time.Sleep(2 * time.Second)
	fmt.Println("runtimeGoExit stop")
	runtime.Goexit() //не паника, любые отложенные функции восстановления будут возвращать nil.
	//редко когда-либо понадобится использовать runtime.Goexit.
	//Он используется для завершения текущей горутины, когда вы не можете вернуться с места вызова,
	//например, когда вы находитесь в вызове функции внутри горутины.
	//В стандартной библиотеке пакет тестирования использует его в таких функциях, как FailNow и SkipNow, для немедленного завершения текущего теста.

}

//также горутина может завершить работу при выполнении всех своих операций.
//также горутина завершит работу при завершении работы функции main, если она не была закрыта функцией runtime.goexit().
