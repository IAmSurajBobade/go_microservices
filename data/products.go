package data

import (
	"encoding/json"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"`
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

func (product *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", func(fl validator.FieldLevel) bool {
		reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
		matches := reg.FindAllString(fl.Field().String(), -1)
		return len(matches) == 1
	})
	return validate.Struct(product)
}

func (product *Product) FromJSON(r io.Reader) error {
	decode := json.NewDecoder(r)
	return decode.Decode(product)
}

type Products []*Product

func (p *Products) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(p)
}

func GetProducts() Products {
	return productList
}

func UpdateProducts(id int, prod Product) error {
	pos, err := getProduct(id)
	if err != nil {
		return err
	}
	prod.ID = id
	prod.UpdatedOn = time.Now().UTC().String()
	productList[pos] = &prod
	return nil
}

func getProduct(id int) (int, error) {
	for i, v := range productList {
		if v.ID == id {
			return i, nil
		}
	}
	return 0, ErrProdNotFound
}
func AddProducts(prod Product) {
	prod.ID = getNextID()
	productList = append(productList, &prod)
}

func getNextID() int {
	prod := productList[len(productList)-1]
	return prod.ID + 1
}

var productList = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Milky coffee",
		Price:       2.45,
		SKU:         "prod001",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Coffee without milk",
		Price:       1.99,
		SKU:         "prod002",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
		DeletedOn:   "",
	},
}
