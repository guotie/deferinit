deferinit
=========

defer init/fini of golang

缘由：
========

1. golang提供了内置的init函数，在程序启动之前执行，在一些情况下，我们需要在init函数执行之前做一些工作，例如读取配置文件，设置log，等等。


2. 在一些时候，我们需要对系统进行清理，即对应于init函数的fini函数


用法：
========

## 1. deferinit.AddInit(f, fi func())

参数为两个func()类型的函数，f为init函数，fi为fini函数，f和fi可以为nil


## 2. deferinit.InitAll()

执行初始化函数

## 3. deferinit.FiniAll()

执行对应的清理函数，按照与init函数相反的顺序执行


示例
========

	func TestDeferInit(t *testing.T) {
		AddInit(func() {
			fmt.Println("1")
		}, nil)
		AddInit(func() {
			fmt.Println("2")
		}, nil)
		AddInit(nil, func() {
			fmt.Println("-3")
		})
		AddInit(func() {
			fmt.Println("4")
		}, func() {
			fmt.Println("-4")
		})

		InitAll()
		FiniAll()
	}

执行的结果为：

	1
	2
	4
	-4
	-3
