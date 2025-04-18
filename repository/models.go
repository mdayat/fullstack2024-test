// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type MyClient struct {
	ID           int32            `json:"id"`
	Name         string           `json:"name"`
	Slug         string           `json:"slug"`
	IsProject    string           `json:"is_project"`
	SelfCapture  string           `json:"self_capture"`
	ClientPrefix string           `json:"client_prefix"`
	ClientLogo   string           `json:"client_logo"`
	Address      pgtype.Text      `json:"address"`
	PhoneNumber  pgtype.Text      `json:"phone_number"`
	City         pgtype.Text      `json:"city"`
	CreatedAt    pgtype.Timestamp `json:"created_at"`
	UpdatedAt    pgtype.Timestamp `json:"updated_at"`
	DeletedAt    pgtype.Timestamp `json:"deleted_at"`
}
