package checkout

// Config ...
type Config struct {
	PublicKey         string
	SecretKey         string
	URI               string
	IdempotencyKey    string
	CancellationToken string
}
