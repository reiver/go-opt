package opt

func (receiver Optional[T]) IsNothing() bool {
	return !(receiver.something)
}
