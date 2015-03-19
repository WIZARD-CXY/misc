type Queue struct{
	elements []interface{}
}

func NewQueue() *Queue{
	return &Queue{make([]interface{},10)}
}

func (*Queue) Push(e interface{}){
	panic("not implemented")
}

func (self *Queue) length() int {
	return len(self.elements)
}