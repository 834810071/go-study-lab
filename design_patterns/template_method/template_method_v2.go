package main

import "fmt"

// WorkInterface 定义一个基类接口，该接口包含了一系列相关的操作
type WorkInterface interface {
	GetUp()
	Working()
	Sleep()
}

// Worker 定义一个工作者的接口，每个工作者都有这些基本操作，然后将具体的实现交给子类
type Worker struct {
	WorkInterface
}

// Daily 定义一个工作者的日常行为
func (w *Worker) Daily() {
	w.getUp()
	w.GetUp()
	w.Working()
	w.Sleep()
}
func (e *Worker) getUp() {
	fmt.Println("Worker...")
}

// NewWorker 任何实现了WorkInterface的子类，都可以传入到其中，然后调用其日常方法，就能实现子类的三个具体实现，即多态的实现
func NewWorker(w WorkInterface) *Worker {
	return &Worker{w}
}

// EnvironmentalSanitationWorker 定义一个清洁员类
type EnvironmentalSanitationWorker struct {
}

func (e *EnvironmentalSanitationWorker) GetUp() {
	fmt.Println("环卫阿姨起床了...")
}

func (e *EnvironmentalSanitationWorker) Working() {
	fmt.Println("环卫阿姨正在打扫卫生...")
}

func (e *EnvironmentalSanitationWorker) Sleep() {
	fmt.Println("环卫阿姨要休息了...")
}

// Programmer 定义一个程序员类
type Programmer struct {
}

func (e *Programmer) GetUp() {
	fmt.Println("程序员起床了...")
}

func (e *Programmer) Working() {
	fmt.Println("程序员正在疯狂敲代码...")
}

func (e *Programmer) Sleep() {
	fmt.Println("程序员睡觉了...")
}
