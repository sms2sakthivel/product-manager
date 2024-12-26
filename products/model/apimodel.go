package model

type ProductCreateRequest struct {
	Name        string `json:"product_name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Price       uint   `json:"price"`
}

type ProductUpdateRequest struct {
	ID          uint   `json:"product_id"`
	Name        string `json:"product_name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Price       uint   `json:"price"`
}

type ProductResponse struct {
	ID          uint   `json:"product_id"`
	Name        string `json:"product_name"`
	Brand       string `json:"brand"`
	Category    string `json:"category"`
	SubCategory string `json:"sub_category"`
	Price       uint   `json:"price"`
}

func (pcr *ProductCreateRequest) GetDBObject() *Product {
	return &Product{Name: pcr.Name, Brand: pcr.Brand, Category: pcr.Category, SubCategory: pcr.SubCategory, Price: pcr.Price}
}

func (pur *ProductUpdateRequest) GetDBObject() *Product {
	return &Product{ID: pur.ID, Name: pur.Name, Brand: pur.Brand, Category: pur.Category, SubCategory: pur.SubCategory, Price: pur.Price}
}
