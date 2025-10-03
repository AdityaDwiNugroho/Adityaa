package handler

import (
	"net/http"
)

type Project struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tech        []string `json:"tech"`
	Link        string   `json:"link"`
	Status      string   `json:"status"`
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "text/html")
	
	projects := []Project{
		{
			Title:       "Journal App - Rust Fullstack",
			Description: "Full-stack journal application built with Rust and Axum framework. Features JWT authentication, PostgreSQL database, and Askama templates for server-side rendering. Clean architecture with secure authentication flow.",
			Tech:        []string{"Rust", "Axum", "PostgreSQL", "JWT", "Askama"},
			Link:        "https://github.com/AdityaDwiNugroho/journal-app",
			Status:      "Active",
		},
		{
			Title:       "Perpustakaan - Library Management",
			Description: "Vue.js-based library management system. Track books, manage borrowing records, and handle inventory. Responsive design with modern UI/UX. Currently under maintenance with planned updates.",
			Tech:        []string{"Vue.js", "JavaScript", "CSS"},
			Link:        "https://github.com/AdityaDwiNugroho/Perpustakaan",
			Status:      "Maintenance",
		},
		{
			Title:       "Random Group Generator",
			Description: "Smart group randomizer tool for team formation. Useful for classrooms, workshops, atau event organizing. Clean JavaScript implementation with intuitive interface. 2 forks from community usage.",
			Tech:        []string{"JavaScript", "HTML", "CSS"},
			Link:        "https://github.com/AdityaDwiNugroho/Random-Group",
			Status:      "Complete",
		},
		{
			Title:       "Study Timer - Pomodoro Technique",
			Description: "Productivity timer implementing 50/10 study technique - 50 minutes focus, 10 minutes break. Clean HTML/CSS interface untuk study sessions. Simple but effective productivity tool.",
			Tech:        []string{"HTML", "CSS", "JavaScript"},
			Link:        "https://github.com/AdityaDwiNugroho/study-time-50-10",
			Status:      "Complete",
		},
		{
			Title:       "Frontend Mentor Challenges",
			Description: "Collection of frontend challenges showcasing responsive design, SCSS proficiency, and modern CSS techniques. Demonstrates UI/UX implementation skills and attention to detail.",
			Tech:        []string{"SCSS", "HTML", "CSS", "JavaScript"},
			Link:        "https://github.com/AdityaDwiNugroho/FrontendMentor",
			Status:      "Ongoing",
		},
	}
	
	html := ""
	for _, p := range projects {
		techTags := ""
		for _, tech := range p.Tech {
			techTags += `<span class="tag">` + tech + `</span>`
		}
		
		html += `
		<div class="project-card glass">
			<div class="project-header">
				<h3>` + p.Title + `</h3>
				<span class="project-status">` + p.Status + `</span>
			</div>
			<p>` + p.Description + `</p>
			<div class="project-tags">` + techTags + `</div>
			<a href="` + p.Link + `" class="project-link" target="_blank">
				View Project [>>]
			</a>
		</div>
		`
	}
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}
