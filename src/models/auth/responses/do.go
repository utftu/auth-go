package responses

type Do struct {
		AccessToken  string `json:"access_token"`
		Bearer       string `json:"bearer"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
		Info         struct {
			Name     string `json:"name"`
			Email    string `json:"email"`
			UUID     string `json:"uuid"`
			TeamUUID string `json:"team_uuid"`
			TeamName string `json:"team_name"`
		} `json:"info"`
}

type DoAccount struct {
	Account struct {
		DropletLimit    int    `json:"droplet_limit"`
		FloatingIPLimit int    `json:"floating_ip_limit"`
		ReservedIPLimit int    `json:"reserved_ip_limit"`
		VolumeLimit     int    `json:"volume_limit"`
		Email           string `json:"email"`
		UUID            string `json:"uuid"`
		EmailVerified   bool   `json:"email_verified"`
		Status          string `json:"status"`
		StatusMessage   string `json:"status_message"`
		Team            struct {
			UUID string `json:"uuid"`
			Name string `json:"name"`
		} `json:"team"`
	} `json:"account"`
}