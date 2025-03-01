package leakybucket

type Bucket interface {
	IsFull() bool
	Inc()
	Dec()
}
