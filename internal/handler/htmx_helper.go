package handler

import (
	"strings"

	"github.com/elghazx/portfolio/internal/models"
	"github.com/labstack/echo/v4"
)

type HTMXHelper struct{}

func NewHTMXHelper() *HTMXHelper {
	return &HTMXHelper{}
}

func (h *HTMXHelper) IsHTMXRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Request") == "true"
}

func (h *HTMXHelper) IsSearchRequest(c echo.Context) bool {
	return c.Request().Header.Get("HX-Trigger") == "search-form" || 
		   c.QueryParam("search") != ""
}

func (h *HTMXHelper) GetSearchTerm(c echo.Context) string {
	return strings.ToLower(strings.TrimSpace(c.QueryParam("search")))
}

func (h *HTMXHelper) FilterProjects(projects []models.Project, searchTerm string) []models.Project {
	if searchTerm == "" {
		return projects
	}

	var filtered []models.Project
	for _, project := range projects {
		if h.matchesSearch(project.Title, project.Summary, searchTerm) {
			filtered = append(filtered, project)
		}
	}
	return filtered
}

func (h *HTMXHelper) FilterPosts(posts []models.Post, searchTerm string) []models.Post {
	if searchTerm == "" {
		return posts
	}

	var filtered []models.Post
	for _, post := range posts {
		if h.matchesSearch(post.Title, post.Summary, searchTerm) {
			filtered = append(filtered, post)
		}
	}
	return filtered
}

func (h *HTMXHelper) matchesSearch(title, summary, searchTerm string) bool {
	titleLower := strings.ToLower(title)
	summaryLower := strings.ToLower(summary)
	return strings.Contains(titleLower, searchTerm) || strings.Contains(summaryLower, searchTerm)
}

func (h *HTMXHelper) SetNoCache(c echo.Context) {
	c.Response().Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Response().Header().Set("Pragma", "no-cache")
	c.Response().Header().Set("Expires", "0")
}
