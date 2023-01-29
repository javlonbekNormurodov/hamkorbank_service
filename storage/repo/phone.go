package repo

import (
	"context"
	"hamkorbank/models"
)

type PhoneRepoI interface {
	GetAllPhones(ctx context.Context, req *models.GetAllPhonesRequest) (*models.GetAllPhonesResponse, error)
	GetPhoneById(ctx context.Context, req *models.GetByIdPhoneRequest) (*models.GetByIdPhoneResponse, bool, error)
}
