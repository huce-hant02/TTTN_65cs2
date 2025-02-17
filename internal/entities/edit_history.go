package entities

import "time"

type CdioEditHistory struct {
	ID         uint    `json:"id"`
	ModelType  string  `json:"modelType"`  // Loại hồ sơ chinh sửa (camelCase)
	ModelId    uint    `json:"modelId"`    // Id của hồ sơ bị chinh sửa
	Data       string  `json:"data"`       // Dữ liệu của hồ sơ bị chinh sửa
	Active     *bool   `json:"active"`     // Trạng thái kích hoạt của lần chỉnh sửa này
	ModifierId *uint   `json:"modifierId"` // Id của employee thực hiện chỉnh sửa
	Note       *string `json:"note"`       // Ghi chú của lần chỉnh sửa này

	Modifier *User `json:"modifier"`

	CreatedAt time.Time  `json:"createdAt" swaggerignore:"true"`
	DeletedAt *time.Time `json:"deletedAt" swaggerignore:"true"`
	UpdatedAt *time.Time `json:"updatedAt" swaggerignore:"true"`
}
