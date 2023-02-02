package postgresrepository

import (
	"context"

	"github.com/bimmagung/flight-ticket-api/domain"
	"github.com/jmoiron/sqlx"
)

func NewFlightRepositoryPostgres(conn *sqlx.DB) domain.FlightRepository {
	return &flightPostgresRepository{
		Conn: conn,
	}
}

type flightPostgresRepository struct {
	Conn *sqlx.DB
}

func (f *flightPostgresRepository) GetAll(ctx context.Context)(result []domain.Flight, err error) {
	query := "SELECT * FROM flights"
	err = f.Conn.SelectContext(ctx ,&result, query)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (f *flightPostgresRepository) Add(ctx context.Context, payload *domain.Flight)(id string, err error) {
	query := "INSERT INTO flights (id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id"
	err = f.Conn.QueryRowxContext(ctx, query, payload.Id, payload.CategoryId, payload.FlightNumber, payload.Departure, payload.Arrive, payload.TimeArrive, payload.Seats, payload.Price).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, nil
}