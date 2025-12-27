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

	"linkedin-bot/internal/humanizer"
	"linkedin-bot/internal/search"
	"linkedin-bot/internal/security"
	"linkedin-bot/internal/session"
)

func main() {
	// -------------------------------
	// 1. LOAD ENV VARIABLES
	// -------------------------------
	err := godotenv.Load()
	if err != nil {
		_ = godotenv.Load("../.env")
	}

	username := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if username == "" || password == "" {
		log.Fatal("❌ LINKEDIN_EMAIL or LINKEDIN_PASSWORD missing in .env")
	}

	// -------------------------------
	// 2. LAUNCH STEALTH BROWSER
	// -------------------------------
	fmt.Println("Launching Stealth Browser...")

	u := launcher.New().
		Headless(false).
		Leakless(false).
		MustLaunch()

	browser := rod.New().
		ControlURL(u).
		MustConnect()
	defer browser.MustClose()

	page := stealth.MustPage(browser)

	loggedIn := false

	// -------------------------------
	// 3. TRY RESTORING SESSION
	// -------------------------------
	fmt.Println("Attempting to restore previous session...")

	if session.LoadCookies(page) {
		page.MustNavigate("https://www.linkedin.com/feed/")
		time.Sleep(6 * time.Second)

		if security.DetectCheckpoint(page) {
			fmt.Println("🛑 Checkpoint detected.")
			fmt.Println("👉 Please solve it manually in the browser.")
			fmt.Println("👉 Press ENTER after resolving the checkpoint.")

			fmt.Scanln() // wait instead of exiting
		}

		if strings.Contains(page.MustInfo().URL, "/feed") {
			fmt.Println("✅ Session restored. Skipping login.")
			loggedIn = true
		}
	}

	// -------------------------------
	// 4. LOGIN IF NEEDED
	// -------------------------------
	if !loggedIn {
		fmt.Println("Navigating to LinkedIn Login...")
		page.MustNavigate("https://www.linkedin.com/login")
		time.Sleep(4 * time.Second)

		fmt.Println("Typing credentials...")
		humanizer.TypeLikeHuman(page.MustElement("#username"), username, true)
		humanizer.TypeLikeHuman(page.MustElement("#password"), password, true)

		loginBtn, _ := page.Element(".btn__primary--large")
		if loginBtn != nil {
			loginBtn.MustClick()
		}

		fmt.Println("------------------------------------------------")
		fmt.Println("⚠️ Solve captcha / 2FA if prompted")
		fmt.Println("Press ENTER once LinkedIn Feed is visible")
		fmt.Println("------------------------------------------------")
		fmt.Scanln()
		if security.DetectCheckpoint(page) {
			fmt.Println("🛑 Checkpoint detected.")
			fmt.Println("👉 Please solve it manually in the browser.")
			fmt.Println("👉 Press ENTER after resolving the checkpoint.")

			fmt.Scanln() // wait instead of exiting
		}

		fmt.Println("Saving session cookies...")
		_ = session.SaveCookies(page)
	}

	// =================================================
	// 5. SEARCH PHASE (PHASE-1: FIRST PAGE ONLY)
	// =================================================
	searchKeyword := "Data Analyst Hyderabad"

	profiles := search.SearchPeopleAndCollectProfiles(page, searchKeyword)

	fmt.Println("------------------------------------------------")
	fmt.Println("PROFILE URLS FROM SEARCH (FIRST PAGE)")
	for i, p := range profiles {
		fmt.Printf("%d. %s\n", i+1, p)
	}
	fmt.Println("------------------------------------------------")

	// -------------------------------
	// 6. NAVIGATE TO MY NETWORK
	// -------------------------------
	fmt.Println("Navigating to My Network...")
	page.MustNavigate("https://www.linkedin.com/mynetwork/")
	time.Sleep(5 * time.Second)

	// -------------------------------
	// 7. SEND CONNECTION REQUESTS
	// -------------------------------
	fmt.Println("------------------------------------------------")
	fmt.Println("STARTING CONNECTION RUN (Max 3)")
	fmt.Println("------------------------------------------------")

	maxInvites := 3
	invitesSent := 0

	for invitesSent < maxInvites {
		fmt.Printf("[%d/%d] Searching for Connect button...\n", invitesSent+1, maxInvites)

		buttons, _ := page.Elements("button")
		clicked := false

		for _, btn := range buttons {
			txt, err := btn.Text()
			if err != nil {
				continue
			}

			if strings.TrimSpace(txt) == "Connect" {
				visible, _ := btn.Visible()
				if !visible {
					continue
				}

				btn.ScrollIntoView()
				humanizer.RandomSleep(500, 1000)

				fmt.Println("   → Clicking Connect")
				if err := btn.Click(proto.InputMouseButtonLeft, 1); err == nil {
					invitesSent++
					clicked = true
					break
				}
			}
		}

		if !clicked {
			fmt.Println("   → No Connect button found. Scrolling...")
			page.Mouse.MustScroll(0, 500)
			time.Sleep(2 * time.Second)
			continue
		}

		time.Sleep(1 * time.Second)

		popupBtns, _ := page.Elements("button")
		for _, pBtn := range popupBtns {
			txt, _ := pBtn.Text()
			if strings.Contains(txt, "Send") {
				fmt.Println("   → Sending invite")
				pBtn.Click(proto.InputMouseButtonLeft, 1)
				break
			}
		}

		fmt.Println("   ✅ Invite sent. Cooling down...")
		humanizer.RandomSleep(5000, 6000)
	}

	// -------------------------------
	// 8. EXIT
	// -------------------------------
	fmt.Println("------------------------------------------------")
	fmt.Println("Automation complete. Closing browser in 5s.")
	time.Sleep(5 * time.Second)
}
