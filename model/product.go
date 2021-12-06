package model

type ProductModel struct {
	Id                 int
	ProductName        string
	Price              float32
	ProductDescription string `json:"description"` //ini seperti alias, jadi parameter yg dikirim adalah description dan property penyimpanannya adalah Product Description
}
