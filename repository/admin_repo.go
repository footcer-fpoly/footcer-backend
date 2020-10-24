package repository

import (
	"context"
)

type AdminRepository interface {
	AcceptStadium(context context.Context, id string) error
}
