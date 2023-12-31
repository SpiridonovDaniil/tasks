package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var d time.Duration = 3
	sleepWithTimer(d)
	sleepWithCtx(d)
}

func sleepWithTimer(d time.Duration) {
	if d <= 0 { //проверяем корректность полученных данных.
		return //если значение <=0 возвращаемся.
	}
	fmt.Println("goroutine sleep")
	e := time.NewTimer(d * time.Second) //создаем новый таймер, который передаст по своему каналу текущее время по истечению времени d.
	<-e.C                               //передаем текущее время.
	//по своей сути это реализация функции time.After().
	fmt.Println("goroutine awaked")
}

func sleepWithCtx(d time.Duration) {
	if d <= 0 { //проверяем корректность полученных данных.
		return //если значение <=0 возвращаемся.
	}
	ctx, cancel := context.WithTimeout(context.Background(), d)
	defer cancel()
	fmt.Println("goroutine sleep")
	select {
	case <-ctx.Done():
		fmt.Println("goroutine awaked")
		return
	}
}

//func timeSleep(ns int64) {
//	if ns <= 0 {
//		return
//	}
//
//	gp := getg()
//	t := gp.timer
//	if t == nil {
//		t = new(timer)
//		gp.timer = t
//	}
//	t.f = goroutineReady
//	t.arg = gp
//	t.nextwhen = nanotime() + ns
//	if t.nextwhen < 0 { // check for overflow.
//		t.nextwhen = maxWhen
//	}
//	gopark(resetForSleep, unsafe.Pointer(t), waitReasonSleep, traceEvGoSleep, 1)
//}

//Разберёмся, что здесь происходит:
//
//Первым делом мы проверяем, что ns больше нуля. Если меньше, то просто возвращаемся. ns - это количество наносекунд, которые мы хотим подождать.
//Далее мы получаем текущую горутину с помощью getg(). Эта функция возвращает указатель на структуру g, которая описывает горутину
//Получаем таймер из горутины с помощью gp.timer. Если таймера нет, то мы создаем новый.
//Это нужно для того, чтобы не создавать новый таймер на каждый вызов time.Sleep(), переиспользуя существующий - это экономит память и время
//Таймер - это структура, описывающая событие, которое должно произойти в будущем. В нашем случае такое событие - это готовность горутины
//Устанавливаем функцию, которая будет вызвана по истечении таймера и аргументы для этой функции.
//Функция goroutineReady просто устанавливает флаг ready в true для указанной горутины. gp - это указатель на текущую горутину, который мы получили выше
//Устанавливаем время, когда таймер должен сработать: получаем текущее время в наносекундах с помощью nanotime(), прибавляем к нему ns и сохраняем в t.nextwhen
//Если t.nextwhen < 0, значит произошло переполнение. В этом случае мы устанавливаем t.nextwhen равным maxWhen (максимальное значение int64)
//Вызываем gopark() для ожидания. gopark() - это функция, которая переводит горутину в состояние ожидания.
//В нашем случае до тех пор, пока не сработает таймер. У неё пять аргументов:
//Функция, которая будет вызвана, когда таймер сработает. В нашем случае это resetForSleep(), которая сбрасывает таймер
//Указатель на таймер, который мы создали или получили выше
//Причина, по которой горутина переводится в состояние ожидания. В нашем случае это waitReasonSleep
//Событие, которое будет записано в трассировку. В нашем случае это traceEvGoSleep
//Флаг, который указывает, что горутина должна быть заблокирована
//Когда таймер сработает, мы вызываем goroutineReady(), которая устанавливает флаг горутины ready в true.
//
//Теперь, когда состояние горутины ready, ей осталось лишь дождаться, когда планировщик её снова запустит.
