<div align="center">
  <img src="https://via.placeholder.com/150" alt="GameCode Logo" width="120" height="120" />
  <h1>GameCode</h1>
  <p><strong>The premier practice platform for Game Developers.</strong></p>
  
  <p>
    <a href="#-philosophy">Philosophy</a> •
    <a href="#-architecture--tech-stack">Tech Stack</a> •
    <a href="#-features">Features</a> •
    <a href="#-getting-started">Getting Started</a> •
    <a href="#-contributing">Contributing</a>
  </p>
</div>

<br/>

**GameCode** is an open-source, high-performance platform providing "LeetCode-density, LeetCode-speed" practice tailored specifically for game development. 

Whether you're mastering spatial math for Unity (C#), deep engine architecture for Unreal Engine (C++), procedural generation in Godot (GDScript), or networking in Roblox (Lua)—GameCode offers isolated, frictionless code execution without the bloat.

---

## 🧭 Philosophy

GameCode is built on a few core principles:
1. **Engine-Agnostic Concepts:** We focus on the math, algorithms, and architectural patterns underlying modern game engines, rather than engine-specific API memorization.
2. **Zero Gamification:** No streaks, no coins, no badges, no social feeds. We respect your time. Just you, the editor, and the compiler.
3. **Frictionless Practice:** Instant code execution, instant feedback, and zero setup.
4. **Performance Above All:** The entire backend and execution pipeline is written in Go to guarantee blazing fast runtimes and robust concurrency.

---

## 🛠 Architecture & Tech Stack

GameCode employs a highly optimized monorepo architecture, utilizing a decoupled Go REST API and a statically optimized Next.js frontend.

### Frontend (`apps/web`)
- **Framework:** Next.js 14 (App Router)
- **Styling:** Tailwind CSS + Radix UI Primitives
- **State Management:** TanStack React Query
- **Editor:** Monaco Editor (VS Code's core editor)

### Backend (`apps/api`)
- **Core:** Go 1.22+ with the Echo (v4) web framework
- **Database:** PostgreSQL via `pgx/v5` and strictly typed `sqlc` queries
- **Caching & Queues:** Redis with `asynq` for background code-execution jobs
- **Security:** JWT authentication (Access/Refresh tokens) alongside Magic Links and OAuth (GitHub/Google)

### Code Execution Engine
- **Sandboxing:** Code is executed in dynamically spun-up, ephemeral Docker containers.
- **Constraints:** Strict security parameters applied per execution (Network disabled, read-only filesystem, 256MB RAM cap, 2-second timeout).
- **Supported Languages:** C# (`dotnet`), C++ (`g++`), Lua (`lua5.4`), GDScript (`godot --headless`).

---

## ✨ Features

- ⚡ **Multi-Language Support**: Write and execute code natively in C#, C++, Lua, and GDScript.
- 🔍 **Full-Text Search**: Instantly find problems using PostgreSQL trigram indexing (`pg_trgm`).
- 📝 **Editorials**: Comprehensive write-ups with engine-specific variants (e.g., "How this applies to Unity vs. Unreal").
- 🔒 **Secure Authentication**: Robust session management with JWTs, Refresh tokens, OAuth, and email Magic Links.
- 🐳 **Isolated Sandboxes**: Highly secure code execution environment.
- 🌙 **Gorgeous Dark Mode**: A sleek, distraction-free interface built for late-night coding.

---

## 🚀 Getting Started

### Prerequisites
Before you begin, ensure you have the following installed:
- [Go 1.22+](https://go.dev/)
- [Node.js 20+](https://nodejs.org/) & `pnpm`/`npm`
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- PostgreSQL & Redis (Can be run locally via Docker Compose)

### 1. Local Infrastructure Setup
Clone the repository and start the underlying database and cache layers using the provided Docker Compose file:
```bash
git clone https://github.com/nabrahma/Game-Code.git
cd Game-Code
docker-compose -f docker-compose.dev.yml up -d
```

### 2. Backend Setup
Navigate to the API folder, configure your environment variables, and run the backend server alongside the execution worker:
```bash
cd apps/api
cp .env.example .env

# Start the API server (Port 8080)
go run cmd/api/main.go

# In a new terminal, start the Asynq worker for code execution
go run cmd/worker/main.go
```

### 3. Frontend Setup
Navigate to the web folder, install dependencies, and spin up the frontend:
```bash
cd apps/web
npm install
npm run dev
```
The application will be live at `http://localhost:3000`.

---

## 🤝 Contributing

GameCode is a community-driven, open-source project. We welcome contributions from developers of all skill levels! Whether you want to add new spatial math problems, optimize the execution engine, or refine the UI, we'd love your help.

### How to Contribute
1. **Fork the Repository**
2. **Create a Feature Branch:** `git checkout -b feature/AmazingFeature`
3. **Commit your Changes:** `git commit -m 'feat: Add some AmazingFeature'`
4. **Push to the Branch:** `git push origin feature/AmazingFeature`
5. **Open a Pull Request**

Please read our `CONTRIBUTING.md` (coming soon) for details on our code of conduct and the process for submitting pull requests.

---

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  <i>Built with ❤️ for Game Developers.</i>
</div>
