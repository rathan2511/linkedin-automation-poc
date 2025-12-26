package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
	"github.com/joho/godotenv"

	// Import our local package
	"linkedin-bot/internal/humanizer"
)

func main() {
	// --- 1. LOAD CONFIGURATION ---
	// We look for .env in the root directory (../.env) because main is inside cmd/
	// However, if you run "go run cmd/main.go" from root, it looks in root.
	err := godotenv.Load()
	if err != nil {
		// Fallback: try loading from one directory up if running from inside cmd
		_ = godotenv.Load("../.env")
	}

	// If still empty, warn user
	if os.Getenv("LINKEDIN_EMAIL") == "" {
		log.Println("Warning: Could not load .env file. Make sure it exists in the project root.")
	}

	linkedInUsername := os.Getenv("LINKEDIN_EMAIL")
	linkedInPassword := os.Getenv("LINKEDIN_PASSWORD")

	if linkedInUsername == "" || linkedInPassword == "" {
		log.Fatal("Error: Credentials empty. Check .env file.")
	}

	fmt.Println("Launching Stealth Browser...")
	u := launcher.New().
		Headless(false).
		Leakless(false).
		MustLaunch()

	browser := rod.New().ControlURL(u).MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser)

	// --- 2. LOGIN ---
	fmt.Println("Navigating to LinkedIn Login...")
	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitStable()

	fmt.Println("Typing Credentials...")
	// USAGE OF NEW PACKAGE HERE:
	humanizer.TypeLikeHuman(page.MustElement("#username"), linkedInUsername)
	humanizer.TypeLikeHuman(page.MustElement("#password"), linkedInPassword)

	// Safe click for login button
	loginBtn, _ := page.Element(".btn__primary--large")
	if loginBtn != nil {
		loginBtn.MustClick()
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("⚠️  CHECK BROWSER: Solve Captcha if needed.")
	fmt.Println("Press ENTER when you see the LinkedIn Feed.")
	fmt.Println("------------------------------------------------")
	fmt.Scanln()

	// --- 3. NAVIGATE TO 'MY NETWORK' ---
	fmt.Println("Navigating to 'My Network'...")
	page.MustNavigate("https://www.linkedin.com/mynetwork/")

	fmt.Println("Waiting 5 seconds for cards to load...")
	time.Sleep(5 * time.Second)

	// --- 4. CONNECT FROM LIST ---
	fmt.Println("------------------------------------------------")
	fmt.Println("STARTING CONNECTION RUN (Max 3 people)")
	fmt.Println("------------------------------------------------")

	maxInvites := 3
	invitesSent := 0

	for i := 0; i < maxInvites; i++ {
		fmt.Printf("[%d/%d] Scanning for 'Connect' button...\n", i+1, maxInvites)

		buttons, _ := page.Elements("button")
		clicked := false

		for _, btn := range buttons {
			txt, err := btn.Text()
			if err != nil {
				continue
			}

			if strings.TrimSpace(txt) == "Connect" {
				if visible, _ := btn.Visible(); visible {

					btn.ScrollIntoView()
					humanizer.RandomSleep(500, 1000) // USAGE HERE

					fmt.Println("   -> Found a Connect button! Clicking...")

					err := btn.Click(proto.InputMouseButtonLeft, 1)
					if err == nil {
						clicked = true
						invitesSent++
						break
					}
				}
			}
		}

		if !clicked {
			fmt.Println("   -> No visible buttons. Scrolling down...")
			page.Mouse.MustScroll(0, 500)
			time.Sleep(2 * time.Second)
			i--
			if invitesSent == 0 && i < -5 {
				fmt.Println("   -> Stopping to prevent infinite loop.")
				break
			}
			continue
		}

		time.Sleep(1 * time.Second)
		popupButtons, _ := page.Elements("button")
		for _, pBtn := range popupButtons {
			pTxt, _ := pBtn.Text()
			if strings.Contains(pTxt, "Send") {
				fmt.Println("   -> Clicking 'Send' in popup...")
				pBtn.Click(proto.InputMouseButtonLeft, 1)
				break
			}
		}

		fmt.Println("   -> ✅ Invite Sent! Waiting 5s...")
		humanizer.RandomSleep(5000, 6000) // USAGE HERE
	}

	fmt.Println("------------------------------------------------")
	fmt.Println("Automation Complete. Closing browser in 5s.")
	time.Sleep(5 * time.Second)
}
