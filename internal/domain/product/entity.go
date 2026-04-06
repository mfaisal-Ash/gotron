package product

type Product struct {
	ID          string  `json:"id"`
	Code		string  `json:"code"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Rate		float64 `json:"rate"`
	ImageURL    string  `json:"image_url"`

}

type ProductRequest struct {
	Code		string  `json:"code" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Price       float64 `json:"price" validate:"required,gt=0"`
	Rate		float64 `json:"rate" validate:"required,gt=0"`
	ImageURL    string  `json:"image_url" validate:"required,url"`
}

