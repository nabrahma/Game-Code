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

Whether you're mastering spatial math for Unity (C#), deep engine architecture for Unreal Engine (C++), procedural generation in Godot (GDScript), or networking in Roblox (Lua)—GameCode offers frictionless code execution, detailed execution constraints, and professional code evaluations.

---

## Philosophy

GameCode is built on a few core principles:
1. **Engine-Agnostic Concepts:** We focus on the math, algorithms, and architectural patterns underlying modern game engines, rather than engine-specific API memorization.
2. **Zero Gamification:** No streaks, no coins, no badges, no social feeds. We respect your time. Just you, the editor, and the compiler.
3. **Frictionless Practice:** Instant code execution, instant feedback, and zero setup.
4. **Professional Standards:** The entire backend and execution pipeline is engineered using industry-standard tools and frameworks designed for scale.

---

## 🛠 Architecture & Tech Stack

GameCode employs a highly optimized monorepo architecture, utilizing a decoupled Go REST API and a statically optimized Next.js frontend.

### Frontend (`apps/web`)
- **Framework:** Next.js 14 (App Router)
- **Styling:** Tailwind CSS + Radix UI Primitives
- **State Management:** TanStack React Query
- **Authentication:** NextAuth.js (Auth.js) for OAuth and Credential management
- **Editor:** Monaco Editor (VS Code's core editor)

### Backend (`apps/api`)
- **Core:** Go 1.22+ with the Echo (v4) web framework
- **Database:** PostgreSQL managed via **GORM** for robust, expressive data access
- **Caching:** Redis for low-latency state and session storage
- **Concurrency:** Native Go routines for parallel execution handling

### Code Execution Engine
- **Sandboxing:** Code is executed securely using the open-source **Judge0** execution engine.
- **Constraints:** Strict security parameters applied per execution (Network disabled, read-only filesystem, strict RAM caps, and runtime limits).
- **Supported Languages:** C# (`dotnet`), C++ (`g++`), Lua (`lua5.4`), GDScript (`godot --headless`).

---

## Features

- **Multi-Language Support**: Write and execute code natively in C#, C++, Lua, and GDScript.
- **Advanced Submission Evaluation**: Batch testcase evaluation running seamlessly against Judge0 APIs.
- **Progress Tracking**: Track your attempts, accepted solutions, runtime performance, and memory footprint.
- **Secure Authentication**: Robust session management out-of-the-box with NextAuth.
- **Distraction-Free Environment**: A sleek, dark-themed interface built for extended focus.

---

## Getting Started

### Prerequisites
Before you begin, ensure you have the following installed:
- [Go 1.22+](https://go.dev/)
- [Node.js 20+](https://nodejs.org/) & `npm`
- PostgreSQL & Redis (Can be run locally or via Docker Compose)
- A running instance of [Judge0 API](https://github.com/judge0/judge0) (or point to the public rapid API)

### 1. Database Setup
Ensure PostgreSQL is running locally on port `5432` with a database named `gamecode`. GORM will automatically migrate and create the necessary tables.

### 2. Backend Setup
Navigate to the API folder, configure your environment variables, and start the backend server:
```bash
cd apps/api
cp .env.example .env

# Configure your DATABASE_URL in the .env file
# Format: postgres://username:password@localhost:5432/gamecode?sslmode=disable

# Start the API server (Port 8080)
go run cmd/api/main.go
```

### 3. Frontend Setup
Navigate to the web folder, install dependencies, and spin up the frontend:
```bash
cd apps/web
cp .env.example .env.local

# Install dependencies and run
npm install
npm run dev
```
The application will be live at `http://localhost:3000`.

---

## Contributing

GameCode is an open-source project designed for rigorous engineering standards. We welcome contributions from developers dedicated to advancing game development education. Whether you want to add new spatial math problems, optimize the execution engine, or refine the architecture, your expertise is valued.

### How to Contribute
1. **Fork the Repository**
2. **Create a Feature Branch:** `git checkout -b feature/AmazingFeature`
3. **Commit your Changes:** `git commit -m 'feat: Add some AmazingFeature'`
4. **Push to the Branch:** `git push origin feature/AmazingFeature`
5. **Open a Pull Request**

Please read our `CONTRIBUTING.md` (coming soon) for details on our code of conduct and the process for submitting pull requests.

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<div align="center">
  <i>Engineered for Game Developers.</i>
</div>
