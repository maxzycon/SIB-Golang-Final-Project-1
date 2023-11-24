package dto

type UserRow struct {
	ID       uint    `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Phone    string  `json:"phone"`
	Role     uint    `json:"role"`
	Image    *string `json:"image"`
}

type UserRowDetail struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Username     string  `json:"username"`
	Email        string  `json:"email"`
	Phone        string  `json:"phone"`
	Address      string  `json:"address"`
	Role         uint    `json:"role"`
	Image        *string `json:"image"`
	DepartmentID *uint   `json:"department_id"`
}

type PayloadUpdateUser struct {
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Password    *string `json:"password"`
	ProfilePath *string `json:"profile_path"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	Role        uint    `json:"role"`
	Department  *uint   `json:"department"`
}

type PayloadLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type PayloadUpdateProfile struct {
	Password string `json:"password"`
}
type PayloadCreateUser struct {
	Name        string  `json:"name"`
	Username    string  `json:"username"`
	Password    string  `json:"password"`
	ProfilePath *string `json:"profile_path"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Address     string  `json:"address"`
	Role        uint    `json:"role"`
	Department  *uint   `json:"department"`
}

type LoginRes struct {
	ID          uint    `json:"id"`
	Avatar      *string `json:"avatar"`
	Username    string  `json:"username"`
	Name        string  `json:"name"`
	AccessToken string  `json:"accessToken"`
	Role        uint    `json:"role"`
	Exp         int64   `json:"exp"`
}

type UserPaginatedRow struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Merchant string `json:"merchant"`
	Role     uint   `json:"role"`
}
