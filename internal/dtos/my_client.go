package dtos

type CreateMyClientRequest struct {
	Name         string `json:"name" validate:"required,max=250"`
	Slug         string `json:"slug" validate:"required,max=100"`
	IsProject    string `json:"is_project" validate:"required,max=30,oneof=0 1"`
	SelfCapture  string `json:"self_capture" validate:"required,max=1"`
	ClientPrefix string `json:"client_prefix" validate:"required,max=4"`
	ClientLogo   string `json:"client_logo" validate:"required,max=255"`
	Address      string `json:"address" validate:"omitempty"`
	PhoneNumber  string `json:"phone_number" validate:"omitempty,max=50"`
	City         string `json:"city" validate:"omitempty,max=50"`
}

type MyClientResponse struct {
	Id           int32  `json:"id"`
	Name         string `json:"name"`
	Slug         string `json:"slug"`
	IsProject    string `json:"is_project"`
	SelfCapture  string `json:"self_capture"`
	ClientPrefix string `json:"client_prefix"`
	ClientLogo   string `json:"client_logo"`
	Address      string `json:"address"`
	PhoneNumber  string `json:"phone_number"`
	City         string `json:"city"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	DeletedAt    string `json:"deleted_at"`
}
