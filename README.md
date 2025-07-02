# Groupie Tracker

A web-based Go application that visually presents data from a RESTful music artist API. Groupie Tracker displays artist profiles, their members, debut dates, and upcoming tour locations through a clean interface. The project is fully modular, combining Go’s powerful backend capabilities with HTML, CSS, and JavaScript for a dynamic frontend.

## Table of Contents

1. [📝 About](#-about)  
2. [📁 File Structure](#-file-structure)  
3. [✨ Features](#-features)  
4. [🚀 Usage Instructions](#-usage-instructions)  
   - [📦 Clone the Repository](#-clone-the-repository)  
   - [▶️ Run the web app](#️-run-the-web-app  
5. [🌐 Live Demo](#-live-demo)  
6. [🔭 Future Plans](#-future-plans)  
7. [🤝 Contributions](#-contributions)  
8. [🙏 Acknowledgements](#-acknowledgements)  
9. [📄 License](#-license)  

## 📝 About

Groupie Tracker is a music-focused web application that fetches and displays data from a third-party API to showcase artists, their members, debut albums, and upcoming tour dates. The project helped me understand how to consume and integrate public APIs, develop a Go backend that communicates with a REST API and display said data, and design a responsive user interface using templating and static assets. It mimics a booking-style showcase for bands/artists and serves as a foundation for future feature-rich web apps.

## 📁 File Structure

```
.
├── README.md
├── assets # Static frontend assets
│   ├── css # page styling
│   │   ├── about.css
│   │   └── index-css # Component-specific CSS files
│   │       ├── index_banner.css
│   │       ├── index_button.css
│   │       ├── index_card_content.css
│   │       ├── index_card_layout.css
│   │       ├── index_footer.css
│   │       └── index_header.css
│   ├── images # Image assets
│   │   ├── 01_founders_logo.jpeg
│   │   ├── github.png
│   │   ├── guitar.jpg
│   │   ├── guitarist.jpg
│   │   ├── headphones_icon.jpg
│   │   ├── html_black_tag.jpg
│   │   ├── html_tag.jpg
│   │   ├── json_icon.jpg
│   │   ├── music_note_clipart.png
│   │   ├── singer.jpg
│   │   └── wallpaper.jpg
│   └── js # Javascript for interactivity
│       └── script.js
├── cmd # Main server entry point and program command
│   └── main.go
├── go.mod
├── handlers # Route logic
│   ├── about.go
│   ├── index.go
│   ├── templates.go
│   └── tourdates.go
├── internal # Internal API fetching logic
│   └── fetch.go
├── models # Structs and shared data
│   └── artist.go
└── templates # HTML files rendered to the user
    ├── about.html
    ├── index.html
    └── tour-dates.html

11 directories, 31 files
```

## ✨ Features

- 🎸 Displays band/artist profiles, members, and debut info
- 📍 View a list of upcoming concert/tour dates
- 🌐 Fully functional web interface powered by Go’s `net/http`
- 🧩 Modular structure with separation of concerns (handlers, models, internal logic)
- 🎨 Clean frontend layout using CSS and responsive image assets

## 🚀 Usage Instructions

### 📦 Clone the Repository

Clone the repository to your local machine:
```bash
git clone https://learn.01founders.co/git/iyoussef/Groupie-Tracker.git
cd groupie-tracker
```

- ### ▶️ Run the web app 
```
go run ./cmd
```
Visit http://localhost:8080 to view the application in your browser.

## 🌐 Live Demo

🔗 Coming soon!

## 🔭 Future Plans
Here are a few enhancements I plan to add in future updates:

- **🗃️ Filtering search feature**: Allow users to filter artist data

- **🌍 Geolocalisation**: Display concert locations on a visual map

- **🔎 Search text input functionality**: Allows users to search for specific text input or artist data

##  🤝 Contributions
Contributions are welcome! If you'd like to help improve **Groupie Tracker**, please follow these steps:

1. **Fork the Repository:**  
   Click the "Fork" button at the top-right of the repository page to create your own copy of the project.

2. **Create a New Branch:**  
   Create a new branch for your feature or bug fix:
   ```bash
    git checkout -b feature-or-bugfix-description
   ```
3. **Make your Changes:**
Implement your changes and ensure that your code adheres to the project's style guidelines.
Tip: Write or update tests as needed.

4. **Commit and Push your Changes**:
Commit your changes with a clear, descriptive message and push your branch to your forked repository:
    ```bash
    git commit -m "Add: description of your changes"
    git push origin feature-or-bugfix-description
    ```
5. **Open a Pull Request**:
Open a pull request (PR) from your branch to the main repository. Please include a clear description of your changes and the motivation behind them.
If you're not sure about a major change, open an issue first to discuss your ideas.

Thank you for helping make ascii-art-web even better!


## 🙏 Acknowledgements
- Created as part of my Go learning journey at 01 Founders


## 📄 License
This project is licensed under the [MIT License](LICENSE).

Acknowledgements
Special Thanks:
Thanks to all contributors, mentors, and peers who provided feedback and support during the development of go-reloaded.

Inspiration:
This project was inspired by best practices in Go development and the need for automated text formatting solutions.

Resources:

The MIT License
Various open-source projects and communities that encourage collaboration and learning.
Thank you for checking out go-reloaded! We hope this tool helps streamline your text processing tasks and that you find it both useful and easy to contribute to.