package loopback

import "testing"

func TestIsHost(t *testing.T) {
	cases := map[string]bool{
		"localhost":           true,
		"LocalHost":           true,
		"127.0.0.1":           true,
		"127.1.2.3":           true,
		"::1":                 true,
		"::ffff:127.0.0.1":    true,
		"":                    false,
		"localhost.":          false,
		"localhost.evil.com":  false,
		"127.0.0.1.nip.io":    false,
		"0.0.0.0":             false,
		"127.1":               false,
		"api.checkout.com":    false,
		"ｌｏｃａｌｈｏｓｔ":           false,
		"[::1]":               false,
		"::ffff:10.0.0.1":     false,
		"evil.example\x00127": false,
	}
	for host, want := range cases {
		if got := IsHost(host); got != want {
			t.Errorf("IsHost(%q) = %v, want %v", host, got, want)
		}
	}
}
