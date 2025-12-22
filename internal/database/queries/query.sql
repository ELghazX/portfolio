-- Projects queries
-- name: GetAllProjects :many
SELECT * FROM projects WHERE status = 'published' ORDER BY created_at DESC;

-- name: GetProjectBySlug :one
SELECT * FROM projects WHERE slug = $1 AND status = 'published';

-- Posts queries  
-- name: GetAllPosts :many
SELECT * FROM posts ORDER BY created_at DESC;

-- name: GetPostBySlug :one
SELECT * FROM posts WHERE slug = $1;

-- Experiences queries
-- name: GetAllExperiences :many
SELECT * FROM experiences ORDER BY start_date DESC;