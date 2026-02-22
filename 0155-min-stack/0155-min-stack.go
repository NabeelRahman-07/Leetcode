type MinStack struct {
    items []int
    min []int
}


func Constructor() MinStack {
    stack := MinStack{
        items : []int{},
        min : []int{},
    }
    return stack
}


func (this *MinStack) Push(val int)  {
    if len(this.min) == 0 || val <= this.min[len(this.min)-1] {
        this.min = append(this.min, val)
    }
    this.items = append(this.items, val)
}


func (this *MinStack) Pop()  {
    temp := this.items[len(this.items)-1]
    this.items = this.items[:len(this.items)-1]
    if this.min[len(this.min)-1] == temp {
        this.min = this.min[:len(this.min)-1]
    }
}


func (this *MinStack) Top() int {
    return this.items[len(this.items)-1]
}


func (this *MinStack) GetMin() int {   
    return this.min[len(this.min)-1]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */