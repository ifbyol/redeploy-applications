package model

import "time"

// Application represents an application deployed within an Okteto namespace
type Application struct {
	Name        string     `json:"name"`
	Status      string     `json:"status"`
	LastUpdated *time.Time `json:"lastUpdated"`
}
