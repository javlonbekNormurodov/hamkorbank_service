package storage

import (
	"hamkorbank/storage/repo"
)

type StorageI interface {
	CloseDB()
	Phone() repo.PhoneRepoI
}
