package domain

import "errors"

var ErrInvalidProductID    = errors.New("productId is invalid")
var ErrInvalidProductPrice = errors.New("product price must be greater than 0")
var ErrProductNotFound     = errors.New("there is no product with such id")
