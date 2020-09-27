package types


//Money int64
type Money int64

//Category string
type Category string

//Status string
type Status string

//Statuses
const (
  StatusOk Status = "OK"
  StatusFail Status = "FAIL"
  StatusInProgress Status = "INPROGRESS"
)

//Payment struct have ID, Amount, Category
type Payment struct {
  ID       int
  Amount   Money
  Category Category
  Status   Status
}
