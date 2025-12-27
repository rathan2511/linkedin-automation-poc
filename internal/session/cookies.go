package session

import (
	"encoding/json"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const CookieFile = "cookies.json"

// SaveCookies saves browser cookies to disk
func SaveCookies(page *rod.Page) error {
	cookies, err := page.Cookies(nil)
	if err != nil {
		return err
	}

	data, err := json.Marshal(cookies)
	if err != nil {
		return err
	}

	return os.WriteFile(CookieFile, data, 0644)
}

// LoadCookies loads cookies from disk into browser
func LoadCookies(page *rod.Page) bool {
	data, err := os.ReadFile(CookieFile)
	if err != nil {
		return false
	}

	var storedCookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &storedCookies); err != nil {
		return false
	}

	// Convert NetworkCookie -> NetworkCookieParam
	var params []*proto.NetworkCookieParam
	for _, c := range storedCookies {
		params = append(params, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	_ = page.SetCookies(params)
	return true
}
