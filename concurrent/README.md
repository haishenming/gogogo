# golang并发编程与Scheduler调度器

## golang



## scheduler 调度器

go程序执行由两层组成，go program, Runtime；

go program即我们编写的代码；

Runtime即golang的运行时，负责与操作系统交流，执行程序，如内存分配，goroutine调度等等；

操作系统是不认识goroutine的，它只认识process和thread，要让操作系统去执行我们的goroutine，就需要一个东西，把goroutine放到线程上去执行，这就是go语言中的scheduler调度器。



### M:N 模型
Runtime 会在程序启动的时候，创建 M 个线程（CPU 执行调度的单位），之后创建的 N 个 goroutine 都会依附在这 M 个线程上执行。这就是 M:N 模型：


