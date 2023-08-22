package entity

type Inventory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CodeItem    string `json:"codeitem"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

// CREATE TABLE Inventories(
// 	ID       int auto_increment primary key,
// 	Name     string
// 	CodeItem string
// 	Stock int,
// 	Description VARCHAR(255),
// 	Status VARCHAR(255)
// )