package entities

import "time"

type CdioEditHistory struct {
	ID         uint    `json:"id" db:"id"`
	ModelType  string  `json:"model_type" db:"model_type"`   // Loại hồ sơ chinh sửa (camelCase)
	ModelId    int     `json:"model_id" db:"model_id"`       // Id của hồ sơ bị chinh sửa
	Data       string  `json:"data" db:"data"`               // Dữ liệu của hồ sơ bị chinh sửa
	Active     *bool   `json:"active" db:"active"`           // Trạng thái kích hoạt của lần chỉnh sửa này
	ModifierId *int    `json:"modifier_id" db:"modifier_id"` // Id của employee thực hiện chỉnh sửa
	Note       *string `json:"note" db:"note"`               // Ghi chú của lần chỉnh sửa này

	Modifier *User `json:"modifier" db:"modifier"`

	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	DeletedAt *time.Time `json:"deleted_at" db:"deleted_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`
}
