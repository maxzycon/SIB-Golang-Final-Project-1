package dto

type TodosRow struct {
	ID          uint   `json:"id"`
	Status      uint   `json:"status"`
	Description string `json:"description"`
}

type PayloadTodos struct {
	Status      uint   `json:"status"`
	Description string `json:"description"`
}
