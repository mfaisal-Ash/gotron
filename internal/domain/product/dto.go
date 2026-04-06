package product

type ProductResponse struct {
	ID       string  `json:"id"`
	Code     string  `json:"code"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Rate     float64 `json:"rate"`
	ImageURL string  `json:"image_url"`
}

type ProductListResponse struct {
	Products []ProductResponse `json:"products"`
}

func NewProductResponse(product *Product) *ProductResponse {
	return &ProductResponse{
		ID:       product.ID,
		Code:     product.Code,
		Name:     product.Name,
		Price:    product.Price,
		Rate:     product.Rate,
		ImageURL: product.ImageURL,
	}
}

func NewProductListResponse(products []*Product) *ProductListResponse {
	productResponses := make([]ProductResponse, len(products))
	for i, product := range products {
		productResponses[i] = *NewProductResponse(product)
	}
	return &ProductListResponse{
		Products: productResponses,
	}
}

func (r *ProductRequest) ToProduct() *Product {
	return &Product{
		Code:     r.Code,
		Name:     r.Name,
		Price:    r.Price,
		Rate:     r.Rate,
		ImageURL: r.ImageURL,
	}
}

func (r *ProductRequest) UpdateProduct(product *Product) {
	product.Code = r.Code
	product.Name = r.Name
	product.Price = r.Price
	product.Rate = r.Rate
	product.ImageURL = r.ImageURL
}
