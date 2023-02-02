package domain

import (
	"context"
	"time"
)

type Flight struct {
	Id             	string 		`db:"id"`	
	CategoryId    	string 		`db:"category_id"`
	FlightNumber  	string 		`db:"flight_number"`
	Departure      	string 		`db:"departure"`
	DepartureTime 	string 		`db:"departure_time"`
	Arrive         	string 		`db:"arrive"`
	TimeArrive    	string 		`db:"time_arrive"`
	Seats          	int			`db:"seats"`
	Price          	int			`db:"price"`
	CreatedAt     	time.Time	`db:"created_at"`
	UpdatedAt 	  	time.Time	`db:"updated_at"`
	IsDeleted 		bool		`db:"is_deleted"`
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
	GetAll(c context.Context) ([]Flight, error)
	Add(c context.Context, payload *Flight) (string, error)
}