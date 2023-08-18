package consts

type ActivityStatus int

const (
	NonExistent ActivityStatus = iota + 1
	NotStarted
	Active
	Ended
)
