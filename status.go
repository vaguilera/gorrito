package gorrito

type ShardStatus struct {
	Name      string `json:"name"`
	RegionTag string `json:"region_tag"`
	Hostname  string `json:"hostname"`
	Services  []struct {
		Status    string `json:"status"`
		Incidents []struct {
			Active    bool   `json:"active"`
			CreatedAt string `json:"created_at"`
			ID        int64  `json:"id"`
			Updates   []struct {
				Severity     string `json:"severity"`
				Author       string `json:"author"`
				CreatedAt    string `json:"created_at"`
				UpdatedAt    string `json:"updated_at"`
				Content      string `json:"content"`
				ID           string `json:"id"`
				Translations []struct {
					Locale    string `json:"locale"`
					Content   string `json:"content"`
					UpdatedAt string `json:"updated_at"`
				} `json:"translations"`
			} `json:"updates"`
		} `json:"incidents"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"services"`
	Slug    string   `json:"slug"`
	Locales []string `json:"locales"`
}

func (gorrito *Gorrito) lolStatus(shardStatus *ShardStatus) int {

	return gorrito.requestAPI(lolStatus, shardStatus, nil)
}
