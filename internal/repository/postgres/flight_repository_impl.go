package postgresrepository

import (
	"context"
	"errors"
	"log"

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

func (f *flightPostgresRepository) get(ctx context.Context, query string, args ...interface{}) (result []domain.Flight, err error) {
	row, err := f.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	defer func() {
		errRow := row.Close()
		if errRow != nil {
			log.Panicln(errRow)
		}
	}()

	result = make([]domain.Flight, 0)

	for row.Next() {
		t := domain.Flight{}
		
		err = row.Scan(
			&t.Id,
			&t.CategoryId,
			&t.FlightNumber,
			&t.Departure,
			&t.DepartureTime,
			&t.Arrive,
			&t.TimeArrive,
			&t.Seats,
			&t.Price,
			&t.UpdatedAt,
			&t.CreatedAt,
			&t.IsDeleted,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, t)
	}

	return result, nil

}

func (f *flightPostgresRepository) GetAll(ctx context.Context)(result []domain.Flight, err error) {
	query := "SELECT * FROM flights"
	//err = f.Conn.SelectContext(ctx ,&result, query)
	result, err = f.get(ctx, query)
	if err != nil {
		return nil, err
	}

	defer f.Conn.Close()

	return result, nil
}


func (f *flightPostgresRepository) Add(ctx context.Context, payload *domain.Flight)(id string, err error) {
	query := "INSERT INTO flights (id) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?) RETURNING id"
	err = f.Conn.QueryRowContext(ctx, query, payload.Id, payload.CategoryId, payload.FlightNumber, payload.Departure, payload.Arrive, payload.TimeArrive, payload.Seats, payload.Price).Scan(&id)

	defer f.Conn.Close()
	
	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *flightPostgresRepository) VerifyAvailable(ctx context.Context, flightNumber string)error {
	query := "SELECT * FROM flights WHERE flight_number = ?"

	list, err := f.get(ctx, query, flightNumber)

	if err != nil {
		return err
	}

	if len(list) > 0 {
		return errors.New("flight not available")
	}

	defer f.Conn.Close()

	return nil
}