package handler

import (
	"embed"
	"strings"
	"time"

	"github.com/elghazx/portfolio/internal/ui/components"
	"github.com/elghazx/portfolio/internal/ui/pages"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	StaticFiles embed.FS
	// Add database/queries dependency here when needed
	// db *database.Queries
}

func NewHandler(staticFiles embed.FS) *Handler {
	return &Handler{
		StaticFiles: staticFiles,
		// Initialize dependencies here
	}
}

func (h *Handler) homeHandler(c echo.Context) error {
	projects := []components.Project{
		{
			ID:          1,
			Title:       "Portfolio V1",
			Description: "Modern web portfolio built with Go, HTMX, and Tailwind CSS. Features responsive design and interactive terminal.",
			Tags:        []string{"go", "templ", "htmx", "tailwind"},
			Image:       "",
			Status:      "completed",
			Slug:        "portfolio-v1",
		},
		{
			ID:          2,
			Title:       "Terminal Access",
			Description: "Interactive shell emulator with command history and real-time navigation.",
			Tags:        []string{"alpine", "javascript", "terminal"},
			Image:       "",
			Status:      "in-progress",
			Slug:        "terminal-access",
		},
	}

	return Render(c, pages.Home(projects))
}

func (h *Handler) projectsHandler(c echo.Context) error {
	projects := []components.Project{
		{
			ID:          1,
			Title:       "Portfolio V1",
			Description: "Modern web portfolio built with Go, HTMX, and Tailwind CSS. Features responsive design and interactive terminal.",
			Tags:        []string{"go", "templ", "htmx", "tailwind"},
			Image:       "",
			Status:      "completed",
			Slug:        "portfolio-v1",
		},
		{
			ID:          2,
			Title:       "Terminal Access",
			Description: "Interactive shell emulator with command history and real-time navigation.",
			Tags:        []string{"alpine", "javascript", "terminal"},
			Image:       "",
			Status:      "in-progress",
			Slug:        "terminal-access",
		},
		{
			ID:          3,
			Title:       "E-Commerce API",
			Description: "RESTful API for e-commerce platform with authentication, product management, and order processing.",
			Tags:        []string{"go", "postgresql", "jwt", "rest"},
			Image:       "",
			Status:      "completed",
			Slug:        "ecommerce-api",
		},
		{
			ID:          4,
			Title:       "Chat Application",
			Description: "Real-time chat application using WebSockets with room management and message history.",
			Tags:        []string{"websocket", "redis", "react"},
			Image:       "",
			Status:      "planning",
			Slug:        "chat-app",
		},
	}

	search := strings.ToLower(c.QueryParam("search"))
	status := c.QueryParam("status")

	var filteredProjects []components.Project
	for _, project := range projects {
		matchSearch := search == "" ||
			strings.Contains(strings.ToLower(project.Title), search) ||
			strings.Contains(strings.ToLower(project.Description), search)

		matchStatus := status == "" || project.Status == status

		if matchSearch && matchStatus {
			filteredProjects = append(filteredProjects, project)
		}
	}

	return Render(c, pages.Projects(filteredProjects))
}

func (h *Handler) experienceHandler(c echo.Context) error {
	experiences := []pages.ExperienceItem{
		{
			ID:          1,
			Type:        "work",
			Period:      "2024-2025",
			Title:       "Laboratory Assistant",
			Institution: "Universitas Mulawarman",
			Description: "As The Coordinator of Laboratory Assistant in the Operating Systems course, I not only facilitated student learning in using virtual machines and cloud-based systems but, also participated in team of teaching assistant in several projects, lab activities. I helped students learn the basics of operating systems, including process management, memory management, file systems, and security.",
		},
		{
			ID:          2,
			Type:        "education",
			Period:      "2023-Present",
			Title:       "Informatics Student",
			Institution: "Universitas Mulawarman",
			Description: "Focus on web development, and cloud computing",
		},
		{
			ID:          3,
			Type:        "certification",
			Period:      "2023-Present",
			Title:       "Certification",
			Institution: "",
			Description: "",
		},
	}

	return Render(c, pages.Experience(experiences))
}

func (h *Handler) notFoundHandler(c echo.Context) error {
	return Render(c, pages.NotFound())
}

func (h *Handler) postsHandler(c echo.Context) error {
	allPosts := []components.Post{
		{
			ID:        1,
			Title:     "Getting Started with Go and HTMX",
			Slug:      "getting-started-go-htmx",
			Summary:   "Learn how to build modern web applications using Go backend with HTMX for dynamic frontend interactions without complex JavaScript frameworks.",
			Content:   "Full content here...",
			Tags:      []string{"Go", "HTMX", "Web Development"},
			CreatedAt: time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        2,
			Title:     "Linux System Administration Tips",
			Slug:      "linux-system-administration-tips",
			Summary:   "Essential Linux commands and system administration techniques for managing servers and development environments effectively.",
			Content:   "Full content here...",
			Tags:      []string{"Linux & Open Source", "System Administration"},
			CreatedAt: time.Date(2023, 12, 10, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        3,
			Title:     "Building RESTful APIs with Echo Framework",
			Slug:      "building-restful-apis-echo-framework",
			Summary:   "Complete guide to creating robust REST APIs using Echo framework in Go, including middleware, validation, and error handling.",
			Content:   "Full content here...",
			Tags:      []string{"Go", "API Development", "Backend"},
			CreatedAt: time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        4,
			Title:     "Docker Container Optimization",
			Slug:      "docker-container-optimization",
			Summary:   "Best practices for optimizing Docker containers, reducing image size, and improving build performance for production deployments.",
			Content:   "Full content here...",
			Tags:      []string{"Docker", "DevOps", "Optimization"},
			CreatedAt: time.Date(2023, 11, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        5,
			Title:     "Gaming Setup for Developers",
			Slug:      "gaming-setup-developers",
			Summary:   "How to balance a productive development environment with a gaming setup, including hardware recommendations and workspace organization.",
			Content:   "Full content here...",
			Tags:      []string{"Gaming", "Hardware & Setup", "Productivity"},
			CreatedAt: time.Date(2023, 11, 20, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        6,
			Title:     "Network Security Fundamentals",
			Slug:      "network-security-fundamentals",
			Summary:   "Understanding basic network security concepts, common vulnerabilities, and essential tools for protecting your infrastructure.",
			Content:   "Full content here...",
			Tags:      []string{"Networking & IT", "Security", "Infrastructure"},
			CreatedAt: time.Date(2023, 11, 15, 0, 0, 0, 0, time.UTC),
		},
	}

	search := strings.ToLower(c.QueryParam("search"))

	var filteredPosts []components.Post
	for _, post := range allPosts {
		matchSearch := search == "" ||
			strings.Contains(strings.ToLower(post.Title), search) ||
			strings.Contains(strings.ToLower(post.Summary), search)

		if matchSearch {
			filteredPosts = append(filteredPosts, post)
		}
	}

	if c.Request().Header.Get("HX-Request") == "true" {
		return Render(c, pages.PostsList(filteredPosts))
	}

	return Render(c, pages.Posts(filteredPosts))
}
