package postgresrepository_test

import (
	"context"
	"testing"
	"time"

	"github.com/bimmagung/flight-ticket-api/domain"
	flightPostgresRepository "github.com/bimmagung/flight-ticket-api/internal/repository/postgres"
	"github.com/stretchr/testify/assert"
	sqlxmock "github.com/zhashkevych/go-sqlxmock"
)

func TestGetAll(t *testing.T) {
	// Arrange
	db, mock, err := sqlxmock.Newx()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mockFlight := []domain.Flight{
		{
			Id: "flight-123", CategoryId: "category-123", FlightNumber: "A-30J", Departure: "Bali", DepartureTime: time.Now().Local().String(), Arrive: "Jakarta", TimeArrive: time.Now().Local().String(), Seats: 158, Price: 1200000, CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false,
		},
		{
			Id: "flight-124", CategoryId: "category-123", FlightNumber: "A-90K", Departure: "Jakarta", DepartureTime: time.Now().Local().String(), Arrive: "Bali", TimeArrive: time.Now().Local().String(), Seats: 158, Price: 1200000, CreatedAt: time.Now(), UpdatedAt: time.Now(), IsDeleted: false,
		},
	}

	rows := sqlxmock.NewRows([]string{"id", "category_id", "flight_number", "departure", "departure_time", "arrive", "time_arrive", "seats", "price", "created_at", "updated_at", "is_deleted"}).
	AddRow(mockFlight[0].Id, mockFlight[0].CategoryId, mockFlight[0].FlightNumber, mockFlight[0].Departure, mockFlight[0].DepartureTime, mockFlight[0].Arrive, mockFlight[0].TimeArrive, mockFlight[0].Seats, mockFlight[0].Price, mockFlight[0].CreatedAt, mockFlight[0].UpdatedAt, mockFlight[0].IsDeleted).
	AddRow(mockFlight[1].Id, mockFlight[1].CategoryId, mockFlight[1].FlightNumber, mockFlight[1].Departure, mockFlight[1].DepartureTime, mockFlight[1].Arrive, mockFlight[1].TimeArrive, mockFlight[1].Seats, mockFlight[1].Price, mockFlight[1].CreatedAt, mockFlight[1].UpdatedAt, mockFlight[1].IsDeleted)

	mock.ExpectQuery("SELECT (.+) FROM flights").WillReturnRows(rows)

	flightRepository :=  flightPostgresRepository.NewFlightRepositoryPostgres(db)

	// Action
	list, err := flightRepository.GetAll(context.Background())
	
	// Assert
	assert.NoError(t, err)
	assert.Len(t, list, 2)
	assert.Equal(t, list[0].Id, mockFlight[0].Id)
	assert.Equal(t, list[1].Id, mockFlight[1].Id)
}