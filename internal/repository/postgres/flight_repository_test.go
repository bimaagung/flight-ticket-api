package repository_test

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "category_id", "flight_number", "departure", "departure_time", "arrive", "time_arrive", "seats", 
	"price", "created_at", "updated_at", "is_deleted"}).
	AddRow("flight-123", "category-123", "A-30J", "Bali", time.Now(), "Jakarta", 
	time.Now(), 10, 1000000, time.Now(), time.Now(), false)
}