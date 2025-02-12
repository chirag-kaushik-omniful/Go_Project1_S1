package modals

type Product struct {
	ID          string `bson:"_id"`
	Name        string `bson:"name"`
	Description string `bson:"description"`
	Price       string `bson:"price"`
}
