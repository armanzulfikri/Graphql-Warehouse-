package models

//Suppliers func
type Suppliers struct {
	ID           uint   `json:"id"`
	SupplierName string `json:"supplier_name"`
	Address      string `json:"address"`
	Telepon      string `json:"telepon"`
}
