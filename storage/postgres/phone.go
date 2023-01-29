package postgres

import (
	"context"
	"errors"
	"hamkorbank/storage/repo"

	"hamkorbank/models"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) repo.PhoneRepoI {
	return &userRepo{
		db: db,
	}
}

func (u userRepo) GetAllPhones(ctx context.Context, req *models.GetAllPhonesRequest) (*models.GetAllPhonesResponse, error) {
	res := &models.GetAllPhonesResponse{}
	offset := (req.Offset - 1) * req.Limit
	rows, err := u.db.Query(ctx, `
	select 
		count(*) over(),
		id, 
		phone_number, 
		created_at 
	from 
		phone
	limit: $1 offset: $2`, req.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		phone := &models.GetByIdPhoneResponse{}
		err = rows.Scan(
			&res.Count,
			&phone.ID,
			&phone.PhoneNumber,
			&phone.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		res.Response = append(res.Response, *phone)
	}

	return res, nil
}

func (u userRepo) GetPhoneById(ctx context.Context, req *models.GetByIdPhoneRequest) (*models.GetByIdPhoneResponse, bool, error) {
	res := &models.GetByIdPhoneResponse{}
	err := u.db.QueryRow(ctx, `select id, phone_number, created_at where id = $1`, req.ID).Scan(
		&res.ID,
		&res.PhoneNumber,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, false, err
	}
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, true, err
	}

	return res, false, nil
}
