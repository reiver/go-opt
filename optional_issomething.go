package opt

func (receiver Optional[T]) IsSomething() bool {
	return receiver.something
}
