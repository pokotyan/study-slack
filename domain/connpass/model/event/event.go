package event

type series struct {
	URL   string `json:"url"`
	ID    int    `json:"id"`
	Title int    `json:"title"`
}

type Event struct {
	EventURL         string `json:"event_url"`
	EventType        string `json:"event_type"`
	OwnerNickname    string `json:"owner_nickname"`
	Series           series `json:"series"`
	UpdatedAt        string `json:"updated_at"`
	StartedAt        string `json:"started_at"`
	EndedAt          string `json:"ended_at"`
	HashTag          string `json:"hash_tag"`
	Title            string `json:"title"`
	EventID          int    `json:"event_id"`
	Lat              string `json:"lat"`
	Lon              string `json:"lon"`
	Waiting          int    `json:"waiting"`
	Limit            int    `json:"limit"`
	OwnerID          int    `json:"owner_id"`
	OwnerDisplayName string `json:"owner_display_name"`
	Description      string `json:"description"`
	Address          string `json:"address"`
	Catch            string `json:"catch"`
	Accepted         int    `json:"accepted"`
	Place            string `json:"place"`
}
