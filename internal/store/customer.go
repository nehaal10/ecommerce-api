package store

type CustomerDatabase struct {
	UserID          string     `bson:"user_id"`
	Cart            []Product  `bson:"cart"`
	EmailID         string     `bson:"email_id"`
	Address         []Address  `bson:"Adress"`
	ProductWishList []WishList `bson:"wishlist"`
	Password        string     `bson:"password"`
}

type WishList struct {
	UserId   string `bson:"user_id"`
	Products []Product
}
type Product struct {
	ProductID   string
	ProductName string
	Price       int
	Quantity    int
}

func Customer() {
	//get all data from cookie

}
