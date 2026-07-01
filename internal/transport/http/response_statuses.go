package http

type ResponceStatus struct {
	code    int
	message string
}

type Status interface {
	Code() int
	Message() string
}

func (r ResponceStatus) Code() int       { return r.code }
func (r ResponceStatus) Message() string { return r.message }
