package handler

import (
	"database/sql"
	"embed"
	"net/http"
	"strings"
	"time"

	"github.com/elghazx/portfolio/internal/ui/components"
	"github.com/elghazx/portfolio/internal/ui/pages"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	StaticFiles embed.FS
}

func NewHandler(staticFiles embed.FS) *Handler {
	return &Handler{
		StaticFiles: staticFiles,
	}
}

// Dummy data
var dummyProjects = []components.Project{
	{
		ID:       1,
		Title:    "Portfolio Website",
		Slug:     "portfolio-website",
		Summary:  "Modern web portfolio built with Go, HTMX, and Tailwind CSS featuring responsive design and interactive terminal.",
		Content:  "This portfolio website showcases modern web development practices using Go for the backend, HTMX for dynamic interactions, and Tailwind CSS for styling. The project demonstrates clean architecture principles, server-side rendering, and progressive enhancement techniques. Features include a responsive design, interactive terminal component, project showcase, and blog functionality.",
		Tags:     []string{"Go", "HTMX", "Tailwind", "Web Development"},
		RepoUrl:  stringPtr("https://github.com/elghazx/portfolio"),
		LiveUrl:  stringPtr("https://portfolio.elghaz.dev"),
		ImageUrl: sql.NullString{String: "/static/images/apple-dark.jpg", Valid: true},
	},
	{
		ID:       2,
		Title:    "E-Commerce API",
		Slug:     "ecommerce-api",
		Summary:  "RESTful API for e-commerce platform with authentication, product management, and order processing.",
		Content:  "A comprehensive e-commerce API built with Go and PostgreSQL. Features include user authentication with JWT, product catalog management, shopping cart functionality, order processing, payment integration, and admin dashboard. The API follows REST principles and includes comprehensive documentation with Swagger.",
		Tags:     []string{"Go", "PostgreSQL", "JWT", "REST API"},
		RepoUrl:  stringPtr("https://github.com/elghazx/ecommerce-api"),
		ImageUrl: sql.NullString{String: "/static/images/apex_octane.jpg", Valid: true},
	},
	{
		ID:       3,
		Title:    "Chat Application",
		Slug:     "chat-application",
		Summary:  "Real-time chat application using WebSockets with room management and message history.",
		Content:  "A real-time chat application built with Go and WebSockets. Features include multiple chat rooms, private messaging, user presence indicators, message history, file sharing, and emoji support. The frontend uses vanilla JavaScript for WebSocket connections and provides a smooth, responsive user experience.",
		Tags:     []string{"Go", "WebSocket", "JavaScript", "Real-time"},
		ImageUrl: sql.NullString{String: "/static/images/AnimeRailway.png", Valid: true},
	},
}

var dummyPosts = []components.Post{
	{
		ID:        1,
		Title:     "Getting Started with Go and HTMX",
		Slug:      "getting-started-go-htmx",
		Summary:   "Learn how to build modern web applications using Go backend with HTMX for dynamic frontend interactions.",
		Content:   "HTMX allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML, using attributes. This tutorial covers setting up a Go server with Echo framework, creating HTMX-powered components, and building interactive web applications without complex JavaScript frameworks. We'll explore practical examples including form submissions, dynamic content loading, and real-time updates.",
		Tags:      []string{"Go", "HTMX", "Web Development", "Tutorial"},
		CreatedAt: time.Date(2023, 12, 15, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        2,
		Title:     "Linux System Administration Tips",
		Slug:      "linux-system-administration-tips",
		Summary:   "Essential Linux commands and system administration techniques for managing servers and development environments.",
		Content:   "A comprehensive guide to Linux system administration covering essential commands, process management, file permissions, network configuration, and security best practices. Learn how to monitor system performance, manage users and groups, configure services, and automate tasks with shell scripts. Perfect for developers and system administrators looking to improve their Linux skills.",
		Tags:      []string{"Linux", "System Administration", "DevOps", "Commands"},
		CreatedAt: time.Date(2023, 12, 10, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        3,
		Title:     "Building RESTful APIs with Echo Framework",
		Slug:      "building-restful-apis-echo-framework",
		Summary:   "Complete guide to creating robust REST APIs using Echo framework in Go with middleware and validation.",
		Content:   "Echo is a high-performance, extensible, minimalist Go web framework. This guide covers building RESTful APIs from scratch, implementing middleware for authentication and logging, request validation, error handling, and API documentation. We'll also explore testing strategies and deployment best practices for production-ready APIs.",
		Tags:      []string{"Go", "Echo", "REST API", "Backend"},
		CreatedAt: time.Date(2023, 12, 5, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        4,
		Title:     "Docker Container Optimization",
		Slug:      "docker-container-optimization",
		Summary:   "Best practices for optimizing Docker containers, reducing image size, and improving build performance.",
		Content:   "Learn advanced Docker techniques to optimize your containers for production. Topics include multi-stage builds, layer caching, security scanning, resource limits, and monitoring. We'll explore how to reduce image sizes, improve build times, and ensure your containers are secure and efficient in production environments.",
		Tags:      []string{"Docker", "DevOps", "Optimization", "Containers"},
		CreatedAt: time.Date(2023, 11, 28, 0, 0, 0, 0, time.UTC),
	},
}

func stringPtr(s string) *string {
	return &s
}

func (h *Handler) homeHandler(c echo.Context) error {
	if c.Request().Header.Get("HX-Request") == "true" {
		return Render(c, pages.HomeContent(dummyProjects[:2]))
	}
	return Render(c, pages.Home(dummyProjects[:2]))
}

func (h *Handler) projectsHandler(c echo.Context) error {
	search := strings.ToLower(c.QueryParam("search"))
	var filteredProjects []components.Project
	
	for _, project := range dummyProjects {
		if search == "" || strings.Contains(strings.ToLower(project.Title), search) || strings.Contains(strings.ToLower(project.Summary), search) {
			filteredProjects = append(filteredProjects, project)
		}
	}

	// For HTMX requests
	if c.Request().Header.Get("HX-Request") == "true" {
		if c.Request().Header.Get("HX-Trigger") == "search-form" {
			return Render(c, pages.ProjectsList(filteredProjects))
		}
		return Render(c, pages.ProjectsContent(filteredProjects))
	}
	
	return Render(c, pages.Projects(filteredProjects))
}

func (h *Handler) projectDetailHandler(c echo.Context) error {
	slug := c.Param("slug")
	
	for _, project := range dummyProjects {
		if project.Slug == slug {
			if c.Request().Header.Get("HX-Request") == "true" {
				return Render(c, pages.ProjectDetailContent(project))
			}
			return Render(c, pages.ProjectDetail(project))
		}
	}
	
	return h.notFoundHandler(c)
}

func (h *Handler) postsHandler(c echo.Context) error {
	search := strings.ToLower(c.QueryParam("search"))
	var filteredPosts []components.Post
	
	for _, post := range dummyPosts {
		if search == "" || strings.Contains(strings.ToLower(post.Title), search) || strings.Contains(strings.ToLower(post.Summary), search) {
			filteredPosts = append(filteredPosts, post)
		}
	}

	// For HTMX requests
	if c.Request().Header.Get("HX-Request") == "true" {
		if c.Request().Header.Get("HX-Trigger") == "search-form" {
			return Render(c, pages.PostsList(filteredPosts))
		}
		return Render(c, pages.PostsContent(filteredPosts))
	}
	
	return Render(c, pages.Posts(filteredPosts))
}

func (h *Handler) postDetailHandler(c echo.Context) error {
	slug := c.Param("slug")
	
	for _, post := range dummyPosts {
		if post.Slug == slug {
			if c.Request().Header.Get("HX-Request") == "true" {
				return Render(c, pages.PostDetailContent(post))
			}
			return Render(c, pages.PostDetail(post))
		}
	}
	
	return h.notFoundHandler(c)
}

func (h *Handler) experienceHandler(c echo.Context) error {
	experiences := []pages.ExperienceItem{
		{
			ID:          1,
			Type:        "work",
			Period:      "2024-2025",
			Title:       "Laboratory Assistant",
			Institution: "Universitas Mulawarman",
			Description: "Coordinator of Laboratory Assistant in Operating Systems course, facilitating student learning in virtual machines and cloud-based systems.",
		},
		{
			ID:          2,
			Type:        "education",
			Period:      "2023-Present",
			Title:       "Informatics Student",
			Institution: "Universitas Mulawarman",
			Description: "Focus on web development, cloud computing, and software engineering principles.",
		},
	}
	
	if c.Request().Header.Get("HX-Request") == "true" {
		return Render(c, pages.ExperienceContent(experiences))
	}
	return Render(c, pages.Experience(experiences))
}

func (h *Handler) notFoundHandler(c echo.Context) error {
	c.Response().Status = http.StatusNotFound
	return Render(c, pages.NotFound())
}
