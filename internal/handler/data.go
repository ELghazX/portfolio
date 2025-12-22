package handler

import (
	"time"

	"github.com/elghazx/portfolio/internal/models"
)

var Projects = []models.Project{
	{
		ID:      1,
		Title:   "Portfolio Website",
		Slug:    "portfolio-web",
		Summary: "A web portfolio built with Go, HTMX, and Tailwind CSS. Server-side rendering with dynamic interactions.",
		Content: `
### stack
- go with echo framework
- htmx for interactivity
- tailwind css for styling
- templ for type-safe templates

### architecture
clean separation using go patterns. static assets embedded at build time. hot reload in development.

### features
- server-side rendering
- progressive enhancement
- responsive design
- fast page loads

### deployment
containerized with docker. ci/cd via github actions. deployed on cloud infrastructure.

the project demonstrates modern web development using traditional server-side techniques enhanced with htmx for dynamic behavior without heavy javascript frameworks.
`,
		Tags:      []string{"Go", "HTMX", "Tailwind", "Web", "alpinejs"},
		RepoUrl:   "https://github.com/elghazx/portfolio",
		LiveUrl:   "https://elghaz.my.id",
		ImageUrl:  "/static/images/portfolio.webp",
		Status:    "published",
		CreatedAt: time.Date(2025, 12, 24, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:      3,
		Title:   "Go Vault",
		Slug:    "go-vault",
		Summary: "A secure file storage web app built with Go. Features JWT authentication, file upload/preview, and clean UI with Templ + HTMX.",
		Content: `

### What it does
Go-Vault lets users securely upload and manage files through a web interface. Built with Go and Echo framework.

### Features
- JWT authentication
- File upload and preview
- SQLite database
- Clean web interface
- Docker support

### Tech Stack
- Go + Echo framework
- SQLite3
- Templ templates + HTMX
- Tailwind CSS
- Docker`,
		Tags:      []string{"Go", "SQLite3", "Media"},
		RepoUrl:   "https://github.com/elghazx/go-vault",
		LiveUrl:   "https://go-vault.elghaz.my.id",
		ImageUrl:  "/static/images/go-vault.webp",
		Status:    "draft",
		CreatedAt: time.Date(2023, 12, 19, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:      2,
		Title:   "Go ShortURl",
		Slug:    "go-shorturl",
		Summary: "A URL shortener (not really cause my domain takin some space) built with Go, PostgreSQL, and Redis. Features real-time click tracking and clean Swiss design interface.",
		Content: `
### Tech Stack

Go backend with hexagonal architecture. PostgreSQL for storage, Redis for caching. HTMX + Tailwind CSS frontend.

### Features

- Short URL generation with 8-character codes
- Click tracking with real-time statistics
- Cache-first URL resolution
- Swiss design interface

### Architecture

Hexagonal pattern separating business logic from external dependencies. Core domain handles URL shortening, adapters manage database and cache.

## Technical Implementation

- **Database**: PostgreSQL with proper indexing
- **Caching**: Redis for session management
- **Deployment**: Docker containerization`,
		Tags:      []string{"Go", "PostgreSQL", "Tailwind", "HTMX"},
		RepoUrl:   "https://github.com/elghazx/go-shorturl",
		LiveUrl:   "https://go-shorturl.elghaz.my.id",
		ImageUrl:  "/static/images/go-shorturl.webp",
		Status:    "published",
		CreatedAt: time.Date(2025, 12, 18, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:      2,
		Title:   "API Booking Cosplay",
		Slug:    "booking-cosplay-backend",
		Summary: "REST API for cosplay rental booking system built with Spring Boot. Manages costume and accessory rentals with user authentication and booking status tracking.",
		Content: `
### Overview
A backend service for ChocoMintCos, a cosplay rental platform that allows users to browse, book, and manage costume and accessory rentals.

### Tech Stack
- **Framework**: Spring Boot 3.4.5
- **Language**: Java 21
- **Database**: MySQL 8.0
- **ORM**: Spring Data JPA
- **Security**: Spring Security (password encryption)
- **Build Tool**: Maven
- **Containerization**: Docker Compose

### Key Features
- User registration and authentication
- Costume and accessory catalog management
- Booking system with status tracking
- RESTful API endpoints
- Database integration with MySQL
- Docker containerization for easy deployment

### API Endpoints
- User management (registration, login)
- Item catalog (costumes, accessories)
- Booking operations (create, update, status management)
- Admin functions for inventory management

### Architecture
Clean architecture with separated layers:
- Controllers for API endpoints
- Services for business logic
- Repositories for data access
- DTOs for request/response handling
- Entities for database mapping

### Deployment
Supports both Docker and local XAMPP deployment with configurable database connections and automated setup through Docker Compose.
`,
		Tags:      []string{"Java", "Spring Boot", "API"},
		RepoUrl:   "https://github.com/elghazx/booking-cosplay-backend",
		LiveUrl:   "",
		ImageUrl:  "/static/images/cosplay.webp",
		Status:    "published",
		CreatedAt: time.Date(2025, 12, 18, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:      2,
		Title:   "EchoArtworks",
		Slug:    "echoartworks",
		Summary: "A social media platform for artists to share artwork, interact through likes and comments, with user authentication and admin management features.",
		Content: `
### Overview
EchoArtworks is a web-based social platform built with PHP and MySQL, designed for artists to showcase their work and build community connections.

### Key Features
- **User Authentication**: Registration, login, and profile management
- **Content Sharing**: Upload and display artwork with descriptions
- **Social Interaction**: Like and comment system for posts
- **Search Functionality**: Find posts and users easily
- **Admin Panel**: User and content management tools
- **Responsive Design**: Mobile-friendly interface

### Technical Stack
- **Backend**: PHP with MySQL database
- **Frontend**: HTML, CSS, JavaScript
- **Database**: MySQL with tables for users, posts, comments, and likes
- **Assets**: Image upload and management system

### Core Functionality
- User registration and authentication system
- Post creation with image uploads
- Interactive like and comment features
- Admin dashboard for content moderation
- Search and discovery tools
- Profile management and customization

### Database Structure
- Users table for account management
- Posts table for artwork storage
- Comments and likes for social interactions
- Secure session management

This project demonstrates full-stack web development skills, database design, user experience design, and social media platform architecture.
`,
		Tags:      []string{"Java", "Spring Boot", "API"},
		RepoUrl:   "https://github.com/elghazx/echoartworks",
		LiveUrl:   "",
		ImageUrl:  "/static/images/echoartworks.webp",
		Status:    "published",
		CreatedAt: time.Date(2025, 12, 18, 0, 0, 0, 0, time.UTC),
	},
}

var Posts = []models.Post{
	{
		ID:      1,
		Title:   "FIX LAG! Hyprland NVIDIA Dual Monitor FPS DROP",
		Slug:    "fix-lag-hyprland-nvidia-dual-monitor-fps-drop",
		Summary: "Masalah FPS drop saat memakai Hyprland + NVIDIA pada konfigurasi dual monitor seringkali berkaitan dengan modeset dan urutan modul kernel. Mengaktifkan modeset untuk nvidia-drm, memastikan modul NVIDIA (dan i915 untuk setup hybrid) dimasukkan ke mkinitcpio, dan menambahkan beberapa environment variables di konfigurasi Hyprland biasanya menyelesaikan masalah dan mengembalikan FPS yang stabil. Setup ini sudah dites pada sistem berbasis Arch (termasuk CachyOS) dan pada laptop Lenovo Ideapad Gaming dengan RTX 3050. Jika butuh bantuan lebih lanjut, periksa log sistem dan konfigurasi yang dibuat, atau coba sesuaikan variabel lingkungan sesuai kebutuhan aplikasi.",
		Content: `
## Video Guide
{{< youtube RMY-dZod5ik >}}

---

Pernah mengalami FPS drop atau tampilan jadi patah-patah setelah menyambungkan laptop ke monitor eksternal saat menggunakan Hyprland dengan kartu grafis NVIDIA? Meski layar laptop dan monitor eksternal sama-sama punya refresh rate tinggi, frame rate bisa anjlok dan tidak stabil. Di bawah ini langkah-langkah praktis yang saya pakai untuk memperbaiki masalah ini — sederhana, langsung ke inti, dan terbukti meningkatkan stabilitas FPS pada setup dual monitor.
## Ringkasan masalah dan setup saya
Gejala: saat laptop dicolok ke monitor eksternal, FPS turun drastis dan tidak stabil pada kedua layar. Ini terjadi padahal refresh rate hardware mendukung (contoh: laptop 165 Hz dan monitor eksternal 60 Hz).

Perangkat yang diuji:
- Lenovo Ideapad Gaming 3i 15IAH7 (15" 1080p 165 Hz)
- NVIDIA RTX 3050
- Monitor eksternal (all-in-one) 24" 1080p 60 Hz
- Distribusi berbasis Arch (mis. CachyOS, Manjaro, atau Arch sendiri)

## Mengapa ini bisa terjadi

Secara garis besar masalah ini biasanya terkait integrasi driver NVIDIA dengan lingkungan Wayland (Hyprland). Mode setting kernel untuk driver NVIDIA dan urutan modul inisialisasi (mkinitcpio) berperan penting. Jika modul tidak dimuat dengan benar atau modeset tidak diaktifkan, compositor dan driver bisa gagal memanfaatkan kemampuan penuh GPU sehingga FPS menjadi tidak stabil.

## Langkah perbaikan (step-by-step)
Langkah-langkah berikut memperbaiki kasus yang saya alami. Pastikan kamu sudah memasang driver NVIDIA yang sesuai untuk distribusi-mu sebelum mengikuti langkah ini.
### 1. Aktifkan modeset untuk driver NVIDIA (modprobe)

Buat atau sunting file konfigurasi modprobe untuk NVIDIA agar modeset diaktifkan. Contoh nama file yang umum digunakan:
` + "`" + `
/etc/modprobe.d/nvidia.conf.
` + "`" + `

Isi yang perlu ditambahkan (contoh):
` + "```bash" + `
options nvidia-drm modeset=1
` + "```" + `

Intinya: pastikan ada opsi yang mengaktifkan modeset (modeset=1) untuk modul nvidia-drm. Setelah menyimpan file ini, perubahan akan dipakai saat initramfs di-generate ulang dan saat modul dimuat.

### 2. Tambahkan modul NVIDIA dan i915 ke mkinitcpio

Buka /etc/mkinitcpio.conf dan tambahkan modul-modul berikut ke baris MODULES. Menyertakan modul i915 juga penting pada laptop hybrid (Intel + NVIDIA) agar aplikasi Electron (Discord, VS Code, dll) tidak mengalami delay atau masalah saat dibuka.

Contoh MODULES:
		MODULES=(i915 nvidia nvidia_modeset nvidia_uvm nvidia_drm)
Setelah menyimpan, rebuild initramfs dengan perintah:
		sudo mkinitcpio -P
Perintah ini akan menerapkan konfigurasi baru sehingga opsi modprobe dan modul yang ditentukan dimuat saat boot.

### 3. Tambahkan environment variables di konfigurasi Hyprland

Beberapa variabel lingkungan membantu integrasi driver NVIDIA dengan compositor Wayland. Edit file konfigurasi Hyprland-mu, misalnya ~/.config/hyperland/hyperland.conf, dan tambahkan environment variables di bagian atas konfigurasi atau di bagian environment jika struktur konfigurasi-mu berbeda.

Contoh variabel yang bisa ditambahkan (sesuaikan bila perlu):
- WLR_NO_HARDWARE_CURSORS=1 — kadang memperbaiki masalah kursor atau tearing pada beberapa setup NVIDIA/Wayland.
- __GLX_VENDOR_LIBRARY_NAME=nvidia — memastikan penggunaan vendor GLX NVIDIA jika diperlukan oleh beberapa aplikasi.
Jika konfigurasi Hyprland-mu memuat file lain atau struktur berbeda, pastikan variabel ini di-source di saat sesi Hyprland dimulai. Alternatifnya, tambahkan variabel ini di skrip startup sesi atau file lingkungan global (mis. ~/.profile) selama sesuai dengan cara kamu memulai Hyprland.

### 4. Reboot dan cek hasil

Setelah langkah di atas, reboot sistem. Jalankan alat pemeriksa FPS seperti glxgears atau aplikasi benchmarking untuk memantau stabilitas FPS pada kedua layar. Pada setup yang diuji, FPS menjadi stabil kembali dan tidak lagi drop drastis saat mengaktifkan monitor eksternal.

## Tips troubleshooting
- Pastikan driver NVIDIA terinstal dengan benar sesuai distro (paket kernel dan driver harus cocok).
- Jika muncul error saat rebuild initramfs, periksa sintaks /etc/mkinitcpio.conf dan pastikan modul yang ditambahkan memang tersedia.
- Periksa log kernel dan sistem untuk pesan terkait NVIDIA: dmesg dan journalctl -b.
- Jika masih bermasalah, coba nonaktifkan sementara variabel environment yang ditambahkan untuk melihat pengaruhnya satu per satu.
- Bila menggunakan distribusi yang berbeda, penempatan file konfigurasi atau perintah regenerasi initramfs bisa berbeda. Sesuaikan langkah dengan dokumentasi distro.

## Kesimpulan
Masalah FPS drop saat memakai Hyprland + NVIDIA pada konfigurasi dual monitor seringkali berkaitan dengan modeset dan urutan modul kernel. Mengaktifkan modeset untuk nvidia-drm, memastikan modul NVIDIA (dan i915 untuk setup hybrid) dimasukkan ke mkinitcpio, dan menambahkan beberapa environment variables di konfigurasi Hyprland biasanya menyelesaikan masalah dan mengembalikan FPS yang stabil.

Setup ini sudah dites pada sistem berbasis Arch (termasuk CachyOS) dan pada laptop Lenovo Ideapad Gaming dengan RTX 3050. Jika butuh bantuan lebih lanjut, periksa log sistem dan konfigurasi yang dibuat, atau coba sesuaikan variabel lingkungan sesuai kebutuhan aplikasi.
`,
		Tags:      []string{"Hyprland", "Linux", "Nvidia", "Tutorial"},
		CreatedAt: time.Date(2025, 12, 9, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:        2,
		Title:     "Cara Main Roblox di Linux 2025 TANPA LAG  | Tutorial Lengkap",
		Slug:      "cara-main-roblox-di-linux-2025-tanpa-lag-tutorial-lengkap",
		Tags:      []string{"Linux", "Gaming", "OpenSources"},
		CreatedAt: time.Date(2025, 12, 8, 0, 0, 0, 0, time.UTC),
		Content: `
## Video Guide
{{< youtube -7SqiyAvSJ8 >}}

---

Ingin migrasi ke Linux tapi takut tidak bisa main Roblox? Tenang, bagi kamu para *streamer* atau pemain hobi, sekarang sudah ada solusi praktis untuk menjalankan Roblox di berbagai distro Linux dengan performa yang lancar.

## Apa itu Sober?

**Sober** adalah aplikasi *third-party client* untuk Roblox yang dirancang khusus agar bisa berjalan di lingkungan Linux. 

> **Catatan Penting:** Sober bukan aplikasi resmi dari Roblox Corp. Berdasarkan forum diskusi, developer Roblox tidak melarang penggunaan Sober (tidak membuat akun di-*ban*), namun keamanan data tidak dijamin oleh pihak Roblox. *Do it with your own risk!*

---

## Langkah-Langkah Instalasi

### 1. Instalasi Flatpak
Sober didistribusikan melalui **Flathub**, jadi pastikan sistem kamu sudah memiliki ` + "`flatpak`" + `.

* **Untuk Arch Linux / CachyOS:**
    ` + "```bash" + `
    sudo pacman -S flatpak
    ` + "```" + `
* **Untuk Ubuntu/Debian:**
    ` + "```bash" + `
    sudo apt install flatpak
    ` + "```" + `

*Disarankan untuk restart PC setelah menginstal Flatpak untuk pertama kali.*

### 2. Menginstal Sober via Flathub
Buka terminal dan jalankan perintah berikut untuk mengunduh dan menginstal Sober:

` + "```bash" + `
flatpak install flathub org.sober.Sober
` + "```" + `

### 3. Setup Awal & Permission
Setelah terinstal, buka Sober melalui menu aplikasi atau terminal. Ikuti wizard yang muncul:
1.  Klik **Continue** pada jendela awal.
2.  Atur preferensi seperti *Controller Access* atau *Discord Rich Presence*.
3.  Tunggu proses download data Roblox selesai.
4.  **PENTING:** Jika diminta permission update, copy command yang muncul di jendela Sober, lalu paste dan jalankan di terminal kamu.

---

## Review Performa & Fitur



### Keunggulan Sober Versi Terbaru:
* **Support Game 18+:** Berbeda dengan versi lama, versi terbaru sekarang mendukung game dengan kategori *Maturity 18+* dan fitur *Voice Chat* selama akun kamu sudah terverifikasi.
* **High FPS:** Mendukung layar dengan *refresh rate* tinggi (misal: 165 FPS) tanpa kendala.
* **Stabilitas:** Performa hampir identik dengan menjalankan Roblox di Windows, bahkan dengan setting grafis "rata kanan".

### Pengalaman Bermain
Saat diuji coba pada game FPS, pergerakan terasa sangat *smooth* dan tidak ada *input lag* yang berarti. Ini menjadikannya solusi terbaik bagi pengguna Linux yang tetap ingin kompetitif di Roblox.

---

## Kesimpulan
Sober adalah *game changer* bagi komunitas Roblox di Linux. Meskipun ada peringatan mengenai keamanan data karena statusnya sebagai aplikasi pihak ketiga, secara fungsionalitas aplikasi ini sudah sangat matang.

**Langkah selanjutnya:** Apakah kamu tertarik mencoba distro Linux tertentu untuk gaming? Kamu juga bisa melihat tutorial main Genshin Impact di Linux pada artikel kami lainnya!`,
		Summary: "Sober hadir sebagai solusi utama bagi pengguna Linux (Ubuntu, Arch, Debian, dll.) untuk bermain Roblox secara lancar tanpa perlu *dual-boot* ke Windows.",
	},
	{
		ID:        3,
		Title:     "Genshin Impact di Linux 2025 | Tanpa Ribet & Anti-Gagal",
		Slug:      "genshin-impact-linux-2025-anti-gagal",
		Tags:      []string{"Linux", "Gaming", "Proton", "Wine"},
		CreatedAt: time.Date(2025, 8, 24, 0, 0, 0, 0, time.UTC),
		Content: `
## Video Guide
{{< youtube 7nJhoYS1XbY >}}

---

Banyak pengguna Linux yang ragu untuk migrasi karena takut tidak bisa bermain game populer seperti **Genshin Impact**. Secara resmi, HoYoverse memang tidak menyediakan *client* asli untuk Linux, namun dengan bantuan tool yang tepat, kita bisa memainkannya dengan performa yang sangat stabil.

## Mengapa Tidak Pakai Steam?

Genshin Impact tidak tersedia di Steam, sehingga kita tidak bisa langsung menggunakan fitur Proton bawaan Steam Deck/Steam. Kita membutuhkan aplikasi pihak ketiga untuk membuat *prefix* (lingkungan virtual) yang menyerupai Windows agar file ` + "`.exe`" + ` bisa dijalankan.

---

## Solusi: Menggunakan Port Proton

**Port Proton** adalah aplikasi yang memudahkan instalasi game Windows di Linux tanpa perlu pusing melakukan konfigurasi manual.

### Cara Instalasi Port Proton:
* **Arch Linux / CachyOS**: Gunakan AUR helper seperti ` + "`yay`" + ` atau ` + "`paru`" + `:
  ` + "```bash" + `
  yay -S port-proton
  ` + "```" + `
* **Ubuntu / Debian**: Kamu bisa mencarinya di format **Flatpak** melalui Flathub.

### Cara Instal HoyoPlay via Port Proton:
1. Buka aplikasi **Port Proton**.
2. Masuk ke tab **Auto Install**.
3. Pilih **HoyoPlay** (Launcher resmi terbaru dari HoYoverse).
4. Klik **Quick Install**. Aplikasi akan mengunduh file resmi langsung dari server HoYoverse.

---

## Tips: Pindah Data Game (Tanpa Download Ulang)

Jika kamu sudah punya file game Genshin dari Windows, kamu tidak perlu mengunduh 90GB lagi. Cukup arahkan atau pindahkan file ke folder prefix Port Proton:
1. Buka folder: ` + "`~/PortProton/prefixes/hoyo/drive_c/Program Files/HoyoPlay/Games`" + `.
2. Copy folder game Genshin kamu ke sana.
3. Di dalam HoyoPlay, klik **Locate Game**.

---

## Review Performa & Penggunaan RAM

Berdasarkan pengujian di **CachyOS (Hyprland)**:
* **Visual**: Sangat bersih, tidak ada bug visual atau *glitch* cahaya yang sering ditemui di launcher lama seperti Lootris.
* **FPS**: Stabil di 60 FPS pada resolusi 1080p.
* **Keunggulan RAM**: Linux jauh lebih hemat. Saat menjalankan Genshin sambil *recording*, penggunaan RAM hanya sekitar **7-8 GB**. Bandingkan dengan Windows yang biasanya memakan 4GB hanya untuk sistem dalam kondisi *idle*. Kamu punya sisa RAM lebih luas untuk *streaming* atau membuka browser.

## Kesimpulan

Bermain Genshin Impact di Linux tahun 2025 sudah sangat mungkin dan sangat lancar. Kamu tidak perlu lagi *dual-boot* ke Windows hanya untuk bermain game ini atau game HoYoverse lainnya seperti Honkai: Star Rail dan ZZZ.

---

Walaupun performanya luar biasa, ingatlah bahwa Linux tetap menggunakan lapisan translasi (Wine/Proton). Jika ada update besar pada sistem *Anti-Cheat* HoYoverse di masa depan, ada kemungkinan metode ini terganggu. Selalu pastikan Port Proton kamu tetap terupdate!

*Punya cara lain yang lebih efektif? Bagikan pengalamanmu di kolom komentar!*`,
		Summary: "Bermain Genshin Impact di Linux tahun 2025 sudah sangat mungkin dan sangat lancar. Kamu tidak perlu lagi *dual-boot* ke Windows hanya untuk bermain game ini atau game HoYoverse lainnya seperti Honkai: Star Rail dan ZZZ.",
	},
}

func ptrTime(t time.Time) *time.Time {
	return &t
}

var Experiences = []models.Experience{
	{
		ID:          1,
		Type:        "work",
		Title:       "Laboratory Assistant",
		Institution: "Universitas Mulawarman",
		Location:    "Samarinda, Indonesia",
		StartDate:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
		EndDate:     ptrTime(time.Date(2025, 11, 1, 0, 0, 0, 0, time.UTC)),
		Description: `As the Laboratory Coordinator for Operating Systems, I lead the instructional team in delivering practical sessions focused on Linux administration, virtualization, and network service configurations. I specialize in breaking down complex system concepts for students, ensuring they gain hands-on proficiency in managing server-side environments. Additionally, I serve as a Teaching Assistant for Data Structures and Algorithms, where I support students in mastering C++ and Python through direct mentoring. My focus is on helping students navigate the challenges of memory management and algorithmic logic, bridging the gap between theoretical programming and practical implementation`,
		CreatedAt:   time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          2,
		Type:        "education",
		Title:       "React & Backend with AI Cohort",
		Institution: "Asah led by Dicoding, Accenture",
		Location:    "Online",
		StartDate:   time.Date(2025, 8, 25, 0, 0, 0, 0, time.UTC),
		EndDate:     nil,
		Description: "Learn frontend fundamental concepts using React and backend development with Node.js, HapiJs, and Redis, integrating AI capabilities.",
		CreatedAt:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          3,
		Type:        "education",
		Title:       "Informatics Student",
		Institution: "Universitas Mulawarman",
		Location:    "Samarinda, Indonesia",
		StartDate:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		EndDate:     nil,
		Description: "Focus on web development, cloud computing, and software engineering principles.",
		CreatedAt:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
	},
	{
		ID:          3,
		Type:        "certification",
		Title:       "AWS re/Start Graduate",
		Institution: "Orbit Future Academy",
		Location:    "Online",
		StartDate:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
		EndDate:     nil,
		Description: "Earners of this badge have completed the AWS re/Start program. AWS re/Start is a skills development and job training program that prepares learners for careers in the cloud. Each cohort, supported by professional mentors and accredited trainers, completes training featuring real-world scenario-based learning, hands-on labs, and coursework.",
		CreatedAt:   time.Date(2023, 8, 1, 0, 0, 0, 0, time.UTC),
	},
	{},
}
