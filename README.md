# 🚀 GOTH Stack Complete Setup

[![Go Version](https://img.shields.io/badge/Go-1.25.5-00ADD8?style=flat-square&logo=go)](https://golang.org/doc/go1.25)
[![Echo](https://img.shields.io/badge/Echo-4.14.0-00ADD8?style=flat-square)](https://echo.labstack.com)
[![HTMX](https://img.shields.io/badge/HTMX-1.9.11-purple?style=flat-square)](https://htmx.org)
[![Tailwind CSS](https://img.shields.io/badge/Tailwind_CSS-3.4.1-38B2AC?style=flat-square&logo=tailwind-css)](https://tailwindcss.com)
[![Alpine.js](https://img.shields.io/badge/Alpine.js-3.13.7-2D3441?style=flat-square)](https://alpinejs.dev)
[![Templ](https://img.shields.io/badge/Templ-0.3.960-red?style=flat-square)](https://templ.guide)
[![SQLC](https://img.shields.io/badge/SQLC-1.27.0-blue?style=flat-square)](https://sqlc.dev)
[![Goose](https://img.shields.io/badge/Goose-3.21.1-orange?style=flat-square)](https://github.com/pressly/goose)

A modern, fully configured starter template for building fast, type-safe web applications using Go, Echo, HTMX, Tailwind CSS, Alpine.js, Templ, SQLC, and Goose. This stack provides an incredible developer experience with type safety from your database all the way to your HTML.

## ✨ Features

- **⚡ Echo Framework**: High performance, minimalist Go web framework
- **📜 Type-Safe Templates**: Using [Templ](https://templ.guide) for compile-time checked templates
- **🔥 Hot Reload**: Using [Air](https://github.com/cosmtrek/air) for instant feedback during development
- **🎨 HTMX + Tailwind**: [Modern, interactive UIs](https://htmx.org) without complex client-side JavaScript
- **🛠️ Modern JS Utilities**: [Alpine.js](https://alpinejs.dev) for lightweight interactivity
- **📱 Responsive Design**: [Mobile-first approach](https://tailwindcss.com) with Tailwind CSS
- **🗄️ Type-Safe Database**: [SQLC](https://sqlc.dev) for compile-time verified SQL queries
- **🔄 Database Migrations**: [Goose](https://github.com/pressly/goose) for schema versioning

## 🚀 Quick Start

### Prerequisites

- Go v1.25.5 or higher
- npm v10.9.0
- node v23.2.0
- Air v1.63.4
- Templ CLI v0.3.960
- SQLC v1.27.0
- Goose v3.21.1
- PostgreSQL (or SQLite)

### Installation

Install dependencies:
```bash
npm install
```

### Running Locally

Start the development server:
```bash
make run
```

This will:
- Start the Templ proxy server on the value of TEMPL_PROXY_PORT defined in the .env file.
- Start the Go server on the value of APP_PORT, also defined in the .env file.
- Enable hot reloading for all file changes.

Access the application at: 
```bash
http://localhost:<TEMPL_PROXY_PORT>
```

### Running as Container

1. Build the container:
```bash
docker build -t <image-tag> .
```

2. Run the container:
```bash
docker run --rm \
  --env-file .env \
  -p <local-port>:<app-port> \
  -t <image-tag>
```

Access the containerized application at:
```bash
http://localhost:<local-port>
```

Note: The `<app-port>` should match the `APP_PORT` in your `.env` file.

