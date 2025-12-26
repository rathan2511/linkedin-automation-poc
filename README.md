# LinkedIn Automation Assignment (PoC)

This is a Proof-of-Concept (PoC) automation tool built with **Go** and **Go-Rod**. It automates the process of logging into LinkedIn and sending connection requests while implementing stealth techniques to avoid bot detection.

## 🚀 Features implemented
1. **Stealth Mode:** Uses `go-rod/stealth` to mask `navigator.webdriver` and browser fingerprints.
2. **Human Simulation:** Implements realistic typing delays and scrolling behavior.
3. **Environment Security:** Uses `.env` files to protect credentials.
4. **Resilient Selectors:** Uses a grid-scanning strategy on the "My Network" page to ensure valid connection buttons are found.

## 🛠️ Tech Stack
- **Language:** Go (Golang)
- **Library:** Go-Rod (Browser Automation)
- **Config:** Godotenv

## ⚙️ Setup & Usage

1. **Clone the repository**
2. **Install dependencies:**
   ```bash
   go mod tidy