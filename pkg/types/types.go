package types


//Money int64
type Money int64

//Category string
type PaymentCategory string

//Status string
type PaymentStatus string

//Statuses
const (
  PaymentStatusOk          PaymentStatus = "OK"
  PaymentStatusFail        PaymentStatus = "FAIL"
  PaymentStatusInProgress  PaymentStatus = "INPROGRESS"
)

//Payment struct have ID, Amount, Category
type Payment struct {
  ID       int
  Amount   Money
  Category PaymentCategory
  Status   PaymentStatus
}

type Phone string

//Account informatio user
type Account struct {
  ID       int64
  Phone    Phone
  Balance  Money
}

