package models

type Product struct {
	Idproducts         int    `json:"id"`
	Price              int    `json:"price"`
	Name               string `json:"name"`
	Article            string `json:"article"`
	Length             int    `json:"length"`
	Photo              string `json:"photo"`
	Width              int    `json:"width"`
	Amount             int    `json:"amount"`
	Brands_name        string `json:"brands_name"`
	Models_name        string `json:"models_name"`
	Year_model_name    string `json:"year_model_name"`
	Glass_types_name   string `json:"glass_types_name"`
	Glass_options_name string `json:"glass_options_name"`
}
