🔗 LinkedIn Automation Assignment (Proof of Concept)

This repository contains a Proof-of-Concept (PoC) LinkedIn automation system built using Go (Golang) and the Rod browser automation framework.

The primary objective of this project is to demonstrate human-like browser automation techniques while minimizing bot detection through stealth behavior, session reuse, and controlled interaction patterns.

⚠️ Note: This project is intended strictly for educational and demonstration purposes.

🚀 Key Features
🔐 Authentication & Session Management

-> Secure login using credentials loaded from environment variables
-> Session persistence using cookies to avoid repeated logins
-> Automatic fallback to manual login when session cookies are invalid
-> Graceful handling of LinkedIn security checkpoints (captcha / verification)

🤖 Human-Like Interaction & Anti-Bot Techniques

The system implements multiple stealth techniques to closely mimic real human behavior:
-> Human-like typing with randomized delays
-> Safe typing mode for credentials (no typos)
-> Randomized scrolling behavior
-> Bézier-curve based mouse movement
-> Hover simulation over page elements
-> Browser fingerprint masking using go-rod/stealth
-> Rate-controlled connection requests
-> Session reuse to reduce automation fingerprints

🤝 Connection Automation

-> Navigates LinkedIn My Network recommendations
-> Identifies visible Connect buttons dynamically
-> Sends connection requests with controlled timing
-> Enforces daily connection limits to prevent aggressive behavior

🔍 Search & Targeting (PoC)

-> Implements keyword-based people search logic
-> Extracts profile URLs from search results
-> Handles scenarios where LinkedIn restricts DOM access gracefully
-> Falls back to My Network recommendations when search results are unavailable

🛑 Security Checkpoint Detection

-> Detects LinkedIn security checkpoint URLs
-> Detects CAPTCHA iframes
-> Pauses automation for manual intervention instead of bypassing security
-> Prevents blind or unsafe automation actions

🛠️ Tech Stack

-> Language: Go (Golang)
-> Browser Automation: Go-Rod
-> Stealth Layer: go-rod/stealth
-> Environment Config: godotenv
-> Browser: Chromium / Chrome (non-headless for demo)

⚙️ Setup & Usage
1️⃣ Clone the Repository
git clone <repository-url>
cd linkedin-bot

2️⃣ Install Dependencies
go mod tidy

3️⃣ Configure Environment Variables

Create a .env file in the project root:

LINKEDIN_EMAIL=your_email@example.com
LINKEDIN_PASSWORD=your_password


⚠️ Credentials are never hardcoded.

4️⃣ Run the Application
go run cmd/main.go


-> If session cookies exist, login is skipped
-> If cookies are deleted or invalid, credentials are typed manually
-> CAPTCHA or verification pages pause automation safely

🔒 Security & Privacy

-> Sensitive runtime files such as cookies.json are excluded using .gitignore
-> No credentials or personal data are committed to version control
-> The system does not bypass LinkedIn security mechanisms

🎯 Design Philosophy

-> Focus on human-like behavior, not aggressive scraping
-> Prioritize safety, reliability, and ethical automation
-> Modular design for easy future extensions

🔮 Future Enhancements

-> Advanced pagination handling
-> Profile-level connection workflow
-> Automated follow-up messaging
-> Analytics-based rate limiting
-> Improved search fallback strategies

⚠️ Disclaimer

This project is a Proof-of-Concept created for academic and learning purposes only.
The author does not encourage or endorse misuse or violation of LinkedIn’s terms of service.


## 🎥 Demo Video

The project demonstration video is available at the link below:

🔗 **Demo Video (Google Drive):** https://drive.google.com/file/d/1HujhhSoG5UnwsWT0XLRvEh456LLqjG-C/view?usp=sharing

