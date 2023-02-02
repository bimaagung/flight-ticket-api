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