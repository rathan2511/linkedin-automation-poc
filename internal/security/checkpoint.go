package security

import (
	"fmt"
	"strings"

	"github.com/go-rod/rod"
)

// DetectCheckpoint checks if LinkedIn triggered a security checkpoint
func DetectCheckpoint(page *rod.Page) bool {
	url := page.MustInfo().URL

	// URL-based detection (most reliable)
	if strings.Contains(url, "/checkpoint/") ||
		strings.Contains(url, "/security/") {
		fmt.Println("🚨 LinkedIn security checkpoint detected (URL)")
		return true
	}

	// CAPTCHA iframe detection
	if el, _ := page.Element(`iframe[src*="captcha"]`); el != nil {
		fmt.Println("🚨 CAPTCHA iframe detected")
		return true
	}

	return false
}
