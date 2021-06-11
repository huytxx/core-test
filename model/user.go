package model

import "time"

type User struct {
	UserId      string    `json:"-" db:"user_id,omitempty"`
	FullName    string    `json:"fullName,omitempty" db:"full_name,omitempty"`
	Email       string    `json:"email,omitempty" db:"email,omitempty"`
	Phone       string    `json:"phone,omitempty" db:"phone,omitempty"`
	Password    string    `json:"-" db:"password,omitempty"`
	DateOfBirth string    `json:"dateOfBirth,omitempty" db:"birth_of_date,omitempty"`
	Title       string    `json:"title,omitempty" db:"title,omitempty"`
	Avatar      string    `json:"avatar,omitempty" db:"avatar,omitempty"`
	DepartmentId string `json:"departmentId,omitempty" db:"department_id,omitempty"`
	ManagerId   string    `json:"managerId,omitempty" db:"manager_id"`
	Role        string    `json:"-" db:"role,omitempty"`
	Token       string    `json:"token,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" db:"updated_at,omitempty"`
}