package timeout

type Timeout struct{}

func NewTimeout() *Timeout {
	return &Timeout{}
}

func (t *Timeout) Run() {
}
