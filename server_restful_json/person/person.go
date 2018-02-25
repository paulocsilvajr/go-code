package person

import "time"

type Person struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	Create_at time.Time `json:"create_at"`
}

type Persons []Person
