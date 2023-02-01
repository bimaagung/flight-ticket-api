package domain

import (
	"context"
	"time"
)

type Flight struct {
	Id             string
	CategoryId    string
	FlightNumber  string
	Departure      string
	DepartureTime string
	Arrive         string
	TimeArrive    string
	Seats          int
	Price          int
	CreatedAt     time.Time
	UpdatedAt time.Time
	IsDeleted bool
}

type FlightReq struct {
	id             string
	category_id    string
	flight_number  string
	departure      string
	departure_time string
	arrive         string
	time_arrive    string
	seats          int 	
	price          int  
}

type FlightRes struct {
	id             string	 	`json:"id"`
	category_id    string 		`json:"category_id"`
	flight_number  string 		`json:"flight_number`
	departure      string 		`json:"departure"`
	departure_time string 		`json:"departure_time"`
	arrive         string  		`json:"arrive"`
	time_arrive    string  		`json:"time_arrive"`
	seats          int 	`json:"seats"`
	price          int  		`json:"price"`
	created_at     time.Time	`json:"created_at"`
	updated_at time.Time		`json:"updated_at"`
}

type FlightRepository interface {
	AddFlight(c context.Context, f Flight) (*Flight, error)
	VerifyAvailableFlight(c context.Context, id string) (string, error)
}