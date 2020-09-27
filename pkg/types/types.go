package types


//Money int64
type Money int64

//Category string
type Category string


//Payment struct have ID, Amount, Category
type Payment struct {
	ID       int
	Amount   Money
	Category Category
}
