package dto

import "github.com/Edw-Castro/Preis/internal/core/domain"

type ResponseProductGet struct {
	Products []domain.Product
}

func BuildResponseProductGet(products []domain.Product) ResponseProductGet {
	response := ResponseProductGet{
		Products: products,
	}
	return response
}

// import "github.com/Edw-Castro/Preis/internal/core/domain"

// type ResponseProductGet domain.Product

// func BuildResponseProductGet(product domain.Product) ResponseProductGet {
// 	return ResponseProductGet(product)
// }
