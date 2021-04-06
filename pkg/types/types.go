package types

type Money int64

type Category string

type Payment	struct {
	ID int
	Amount Money
	Category Category
	Status Status

}
//lawkjdlawkjdlkawjdkawd

type Status string

const (
	StatusOk Status = "OK"
	StatusFaiL Status = "Fail"
	StatusInProgress Status ="Inprogress"
	)




