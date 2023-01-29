package models

type GetAllPhonesRequest struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type GetAllPhonesResponse struct {
	Count    int                    `json:"count"`
	Response []GetByIdPhoneResponse `json:"response"`
}

type GetByIdPhoneRequest struct {
	ID string `json:"id"`
}

type GetByIdPhoneResponse struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
}
