package gorrito

type Client struct {
	region   string
	token    string
	regionV5 string
}

// Move it to another file if it grows
type Config struct {
	Region string
}

func NewClient(token string, cfg *Config) *Client {
	region5 := ""

	switch cfg.Region {
	case "NA", "BR", "LAN", "LAS", "OCE":
		region5 = "AMERICAS"
	case "KR", "JP":
		region5 = "ASIA"
	case "EUNE", "EUW", "TR", "RU":
		region5 = "EUROPE"
	default:
		region5 = ""
	}
	if region5 == "" {
		return nil
	}

	c := Client{
		token:    token,
		region:   cfg.Region,
		regionV5: region5,
	}

	return &c
}
