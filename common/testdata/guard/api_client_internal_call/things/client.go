package things

func Send(c interface{ invoke(string) error }) error {
	return c.invoke("/things")
}
