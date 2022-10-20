package mail

// Config
// contains config parameters for mailgun.
type Config struct {
	Enabled bool   `koanf:"enabled"`
	Domain  string `koanf:"domain"`
	APIKEY  string `koanf:"api_key"`
}
