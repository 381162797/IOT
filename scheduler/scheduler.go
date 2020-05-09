package scheduler

//Run 接收地址,开始任务
func Run(addr string) {
	go receiveTCP(addr)
	go producer()
	go consumer()
	go loadTimer()
}
