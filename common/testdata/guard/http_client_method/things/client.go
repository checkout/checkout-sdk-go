package things

type sender struct {
	HttpClient interface{ Get(url string) error }
}

func Send(s sender) error {
	return s.HttpClient.Get("https://example.com/things")
}
