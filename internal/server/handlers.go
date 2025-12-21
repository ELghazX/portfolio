package server

import (
	"github.com/elghazx/portfolio/internal/components"
	"github.com/elghazx/portfolio/internal/components/pages"
	"github.com/labstack/echo/v4"
)

func homeHandler(c echo.Context) error {
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

func projectsHandler(c echo.Context) error {
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

	return Render(c, pages.Projects(projects))
}

func experienceHandler(c echo.Context) error {
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

func notFoundHandler(c echo.Context) error {
	return Render(c, pages.NotFound())
}
