package models

import "time"

// Project represents a project in the portfolio
type Project struct {
	ID        int32
	Title     string
	Slug      string
	Summary   string
	Content   string
	RepoUrl   string
	LiveUrl   string
	Tags      []string
	ImageUrl  string
	Status    string
	CreatedAt time.Time
}

// Post represents a blog post
type Post struct {
	ID        int32
	Title     string
	Slug      string
	Summary   string
	Content   string
	Tags      []string
	CreatedAt time.Time
}

// Experience represents work/education experience
type Experience struct {
	ID          int
	Type        string
	Title       string
	Institution string
	Location    string
	StartDate   time.Time
	EndDate     *time.Time
	Description string
	CreatedAt   time.Time
}
