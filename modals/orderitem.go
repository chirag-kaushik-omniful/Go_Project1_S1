package modals

type Orderitem struct {
	ID        string `bson:"_id"`
	OrderId   string `bson:"orderid"`
	ProductId string `bson:"productid"`
}
