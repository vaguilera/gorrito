package models

type content struct {
	Locale  string `json:"locale"`
	Content string `json:"content"`
}

type status struct {
	ID                int       `json:"id"`
	MaintenanceStatus string    `json:"maintenance_status"`
	IncidentSeverity  string    `json:"incident_severity"`
	Titles            []content `json:"titles"`
	Updates           []struct {
		ID               int       `json:"id"`
		Author           string    `json:"author"`
		Publish          bool      `json:"publish"`
		PublishLocations []string  `json:"publish_locations"`
		Translations     []content `json:"translations"`
		CreatedAt        string    `json:"created_at"`
		UpdatedAt        string    `json:"updated_at"`
	} `json:"updates"`
	CreatedAt string   `json:"created_at"`
	ArchiveAt string   `json:"archive_at"`
	UpdatedAt string   `json:"updated_at"`
	Platforms []string `json:"platforms"`
}

type ShardStatus struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	Locales      []string `json:"locales"`
	Maintenances []status `json:"maintenances"`
	Incidents    []status `json:"incidents"`
}
