package person

import "time"

type Person struct {
	Name      string    `json:"name"`
	Active    bool      `json:"active"`
	Create_at time.Time `json:"create_at"`
}

type Persons []Person
