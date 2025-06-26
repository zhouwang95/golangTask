package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
题目 ：编写一个Go程序，定义一个函数，该函数接收一个整数指针作为参数，在函数内部将该指针指向的值增加10，然后在主函数中调用该函数并输出修改后的值。
考察点 ：指针的使用、值传递与引用传递的区别。
*/
func addTen(p *int32) {
	*p = *p + 10
}

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/

func multiplyTwo(p *[]int32) {
	num := *p
	for i, v := range num {
		num[i] = v * 2
	}
}

/*
题目 ：编写一个程序，使用 go 关键字启动两个协程，一个协程打印从1到10的奇数，另一个协程打印从2到10的偶数。
考察点 ： go 关键字的使用、协程的并发执行。
*/
func printNum() {
	go func() {
		for i := 1; i < 11; i++ {
			if i%2 == 1 {
				fmt.Println(i)
			}
			time.Sleep(10 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i < 11; i++ {
			if i%2 == 0 {
				fmt.Println(i)
			}
		}
		time.Sleep(10 * time.Millisecond)
	}()

}

/*
题目 ：设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
考察点 ：协程原理、并发任务调度。
*/
func taskScheduler(tasks []func()) {
	var wg sync.WaitGroup
	for index, task := range tasks {
		wg.Add(1)
		go func(t func()) {
			defer wg.Done()
			startTime := time.Now()
			fmt.Printf("开始执行任务%d\n", index+1)
			t()
			useTime := time.Since(startTime)
			fmt.Printf("执行任务%d结束,花费时间%v\n", index+1, useTime)
		}(task)
	}
	wg.Wait()
}

/*
题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，
实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
*/
type Shape interface {
	Area()
	Perimeter()
}
type Rectangle struct {
}
type Circle struct {
}

func (r *Rectangle) Area() {
	fmt.Println("调用Rectangle实现的Area函数")
}
func (r *Rectangle) Perimeter() {
	fmt.Println("调用Rectangle实现的Perimeter函数")
}
func (c *Circle) Area() {
	fmt.Println("调用Circle实现的Area函数")
}
func (c *Circle) Perimeter() {
	fmt.Println("调用Circle实现的Perimeter函数")
}

/*
题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，
组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/
type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	EmployeeID string
}

func (e *Employee) PrintInfo() {
	fmt.Printf("Name:%s\n", e.Person.Name)
	fmt.Printf("Age:%d\n", e.Person.Age)
	fmt.Printf("EmployeeID:%s\n", e.EmployeeID)
}

/*
题目 ：编写一个程序，使用通道实现两个协程之间的通信。一个协程生成从1到10的整数，
并将这些整数发送到通道中，另一个协程从通道中接收这些整数并打印出来。
考察点 ：通道的基本使用、协程间通信。
*/
func ChannelCommunication() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go func(chan<- int) {
		defer wg.Done()
		for i := 1; i < 11; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	wg.Add(1)
	go func() {
		defer wg.Done()
		timeout := time.After(10 * time.Second)
		for {
			select {
			case re, ok := <-ch:
				if !ok {
					fmt.Println("channel 已关闭")
					return
				}
				fmt.Printf("接收数据:%d\n", re)
			case <-timeout:
				fmt.Println("操作超时")

			default:
				fmt.Println("没有数据，等待中......")
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	wg.Wait()
}

/*
题目 ：实现一个带有缓冲的通道，生产者协程向通道中发送100个整数，消费者协程从通道中接收这些整数并打印。
考察点 ：通道的缓冲机制。
*/
func Producer(ch chan<- int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(ch chan<- int) {
		defer wg.Done()
		for i := 1; i < 101; i++ {
			ch <- i
			fmt.Printf("发送数据:%d\n", i)
		}
		close(ch)
	}(ch)
}
func Consumer(ch <-chan int, wg *sync.WaitGroup) {
	wg.Add(1)
	go func(ch <-chan int) {
		defer wg.Done()
		for {
			select {
			case v, ok := <-ch:
				if !ok {
					fmt.Println("channel 已关闭")
					return
				}
				fmt.Printf("接收数据:%d\n", v)
			default:
				fmt.Println("没有数据，等待中......")
				time.Sleep(1 * time.Millisecond)
			}
		}
	}(ch)
	wg.Wait()
}

/*
题目 ：编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ： sync.Mutex 的使用、并发数据安全。
*/

func LockCounter() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	count := 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				lock.Lock()
				count++
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("count:%d\n", count)
}

/*
题目 ：使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
考察点 ：原子操作、并发数据安全。
*/
func NoLockCounter() {
	var wg sync.WaitGroup
	var count uint32 = 0
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddUint32(&count, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Printf("count:%d\n", count)
}
func main() {
	NoLockCounter()

	//LockCounter()

	/*var wg sync.WaitGroup
	ch := make(chan int, 100)
	Producer(ch, &wg)
	Consumer(ch, &wg)
	wg.Wait()*/

	//ChannelCommunication()

	/*p := Person{"anli", 20}
	emp := Employee{p, "123456"}
	emp.PrintInfo()*/

	/*var shape = &Rectangle{}
	shape.Area()
	shape.Perimeter()

	shape = &Circle{}
	shape.Area()
	shape.Perimeter()*/

	/*tasks := []func(){
		func() { fmt.Println("执行任务1"); time.Sleep(5 * time.Second) },
		func() { fmt.Println("执行任务2"); time.Sleep(4 * time.Second) },
		func() { fmt.Println("执行任务3"); time.Sleep(3 * time.Second) },
		func() { fmt.Println("执行任务4"); time.Sleep(2 * time.Second) },
		func() { fmt.Println("执行任务5"); time.Sleep(1 * time.Second) },
	}
	taskScheduler(tasks)*/

	/*printNum()
	//防止main方法结束，而两个协程未结束，导致不能打印出数据
	time.Sleep(1000 * time.Millisecond)*/

	/*s := []int32{2, 4, 6, 5}
	multiplyTwo(&s)
	fmt.Printf("out value, %d\n", s)
	*/
	/*var num int32 = 5
	addTen(&num)
	fmt.Printf("out value, %d\n", num)*/
}
