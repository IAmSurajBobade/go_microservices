package data

import "fmt"

// ErrProdNotFound is an error raised when a product can not be found in the database
var ErrProdNotFound = fmt.Errorf("Product not found")
