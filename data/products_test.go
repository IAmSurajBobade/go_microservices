package data

import (
	"testing"
)

func TestProduct_Validate(t *testing.T) {
	tests := []struct {
		name       string
		product    *Product
		wantErrMsg string
		wantErr    bool
	}{
		{"empty struct", &Product{}, `Key: 'Product.Name' Error:Field validation for 'Name' failed on the 'required' tag
Key: 'Product.Price' Error:Field validation for 'Price' failed on the 'gt' tag
Key: 'Product.SKU' Error:Field validation for 'SKU' failed on the 'required' tag`, true},
		{"only name", &Product{Name: "Some product"}, `Key: 'Product.Price' Error:Field validation for 'Price' failed on the 'gt' tag
Key: 'Product.SKU' Error:Field validation for 'SKU' failed on the 'required' tag`, true},
		{
			name:       "name and price",
			product:    &Product{Name: "Some product", Price: 1.7},
			wantErrMsg: `Key: 'Product.SKU' Error:Field validation for 'SKU' failed on the 'required' tag`,
			wantErr:    true,
		},
		{
			name:       "invalid sku",
			product:    &Product{Name: "Some product", Price: 1.7, SKU: "asd"},
			wantErrMsg: "Key: 'Product.SKU' Error:Field validation for 'SKU' failed on the 'sku' tag",
			wantErr:    true,
		},
		{
			name:       "valid product",
			product:    &Product{Name: "Some product", Price: 1.7, SKU: "asd-qwe-rew"},
			wantErrMsg: "",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.product.Validate()
			if err != nil {
				// Check if we expected error
				if !tt.wantErr {
					t.Errorf("Product.Validate() did not expect error [wantErr %v]", tt.wantErr)
				}
				// Check error is as expected
				if err.Error() != tt.wantErrMsg {
					t.Errorf("Product.Validate() error = \n%v, wantErr \n%v", err.Error(), tt.wantErrMsg)
				}
			}
		})
	}
}
