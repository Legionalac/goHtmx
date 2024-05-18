package authservice

import (
	"context"
	"goHtmx/internal/core"

	"github.com/jackc/pgx/v5"
)

type AuthService struct {
	Db libs.CockroachDb
}
func (service *AuthService) CreateAccount(email string, pass string) error{
	query := "INSERT INTO profile (email, pass) VALUES (@email, @pass)"
	args := pgx.NamedArgs{
		"email": email,
		"pass": pass,
	}
	_, err := service.Db.ConnectionPool.Exec(context.Background(), query, args)
	if err != nil {
        return err
    }
    return nil
}