package config

type Config struct {
	SpotifyClientId     string `mapstructure:"SPOTIFY_CLIENT_ID"`
	SpotifyClientSecret string `mapstructure:"SPOTIFY_CLIENT_SECRET"`
	SpotifyBaseUrlApi   string `mapstructure:"SPOTIFY_BASE_URL_API"`
	SpotifyAuthApi      string `mapstructure:"SPOTIFY_AUTH_API"`
}
