package event

type series struct {
	Url   string `json:url`
	Id    int    `json:id`
	Title int    `json:title`
}

type Event struct {
	EventUrl         string  `json:event_url`
	EventType        string  `json:event_type`
	OwnerNickname    string  `json:owner_nickname`
	Series           series  `json:series`
	UpdatedAt        string  `json:updated_at`
	Lat              float64 `json:lat`
	StartedAt        string  `json:started_at`
	HashTag          string  `json:hash_tag`
	Title            string  `json:title`
	EventId          int     `json:event_id`
	Lon              float64 `json:lon`
	Waiting          int     `json:waiting`
	Limit            int     `json:limit`
	OwnerId          int     `json:owner_id`
	OwnerDisplayName string  `json:owner_display_name`
	Description      string  `json:description`
	Address          string  `json:address`
	Catch            string  `json:catch`
	Accepted         int     `json:accepted`
	Ended_at         string  `json:ended_at`
	Place            string  `json:place`
}
