package models

type Service struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     int    `json:"version"`
}

type ServiceVersion struct {
	ServiceID string `json:"service_id"`
	Version   int    `json:"version"`
}
