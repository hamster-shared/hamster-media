package consts

type ActivityStatus int

const (
	NonExistent ActivityStatus = iota
	NotStarted
	Active
	Ended
)
