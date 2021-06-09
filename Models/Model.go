package Models

import "github.com/jinzhu/gorm"

type Manager struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

type Pharmacist struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

type Salesman struct {
	gorm.Model
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

type Medices struct {
	gorm.Model
	Name string `json:"name"`
	Mg   int    `json:"mg"`
	Salt int    `json:"salt"`

}
