package dto

type CreateMasterSystemOption struct {
	OptionName string `json:"optionName,omitempty" binding:"required"`
	Value      string `json:"value,omitempty" binding:"required"`
}
type UpdateMasterSystemOption struct {
	OptionName string `json:"optionName,omitempty" binding:"required"`
	Value      string `json:"value,omitempty" binding:"required"`
}
