package modals

type Order struct {
	ID         string `bson:"_id" json:"id"`
	CustomerId string `bson:"customerid" json:"customerid"`
}
