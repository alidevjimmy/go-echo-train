package domains

type Product struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Price uint   `json:"price"`
}
