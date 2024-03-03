package implement_stack_using_queues

type MyStack struct {
	queue1, queue2 []int
}

func Constructor() (s MyStack) {
	return
}

func (this *MyStack) Push(x int) {
	this.queue2 = append(this.queue2, x)
	for len(this.queue1) != 0 {
		ele := this.queue1[0]
		this.queue1 = this.queue1[1:]
		this.queue2 = append(this.queue2, ele)
	}
	this.queue1, this.queue2 = this.queue2, this.queue1
}

func (this *MyStack) Pop() int {
	v := this.queue1[0]
	this.queue1 = this.queue1[1:]
	return v
}

func (this *MyStack) Top() int {
	return this.queue1[0]
}

func (this *MyStack) Empty() bool {
	return len(this.queue1) == 0
}
