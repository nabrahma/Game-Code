# GameCode 🎮

**LeetCode for Game Developers.**

GameCode is an open-source practice platform designed specifically for game developers. It provides high-density, real-world problems tailored for game engine contexts (Unity, Unreal Engine, Godot, and Roblox). Master spatial math, engine architecture, pathfinding, and optimization through frictionless, single-purpose code execution.

![GameCode Architecture](https://via.placeholder.com/800x400?text=GameCode+Architecture)

## 🌟 Philosophy
- **Frictionless Practice**: No streaks, no coins, no gamification nonsense. Just you and the code.
- **Engine-Agnostic Concepts**: Focus on the core math and architecture underlying the engines.
- **Blazing Fast**: Built entirely in Go for maximum performance and minimal overhead.
- **Sandboxed Execution**: Safe, secure, and isolated code execution using Docker for multiple languages (C#, C++, Lua, GDScript).

## 🛠 Tech Stack
- **Frontend**: Next.js 14 (App Router), Tailwind CSS, React Query, Monaco Editor
- **Backend API**: Go (Echo v4), PostgreSQL (pgx, sqlc), Redis (Upstash)
- **Execution Engine**: Asynq Workers + Docker Runner (Memory/CPU limits, network disabled)
- **Infrastructure**: Vercel (Web), Railway/Render (API)

## 🚀 Getting Started

### Prerequisites
- [Go 1.22+](https://go.dev/)
- [Node.js 20+ & npm/pnpm](https://nodejs.org/)
- [Docker Desktop](https://www.docker.com/products/docker-desktop)
- PostgreSQL & Redis (Running locally via Docker Compose or hosted)

### Local Development Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/nabrahma/Game-Code.git
   cd Game-Code
   ```

2. **Start Local Database & Cache**
   We use `docker-compose` to spin up PostgreSQL and Redis.
   ```bash
   docker-compose -f docker-compose.dev.yml up -d
   ```

3. **Backend API (Go)**
   ```bash
   cd apps/api
   
   # Set up environment variables (copy .env.example to .env)
   cp .env.example .env
   
   # Run the API server
   go run cmd/api/main.go
   
   # In a separate terminal, run the execution worker
   go run cmd/worker/main.go
   ```

4. **Frontend Web (Next.js)**
   ```bash
   cd apps/web
   
   # Install dependencies
   npm install
   
   # Start the development server
   npm run dev
   ```
   Open `http://localhost:3000` in your browser.

## 🏗 Repository Structure (Monorepo)

```text
GameCode/
├── apps/
│   ├── api/          # Go backend API & workers
│   └── web/          # Next.js 14 frontend
├── docker/
│   └── runners/      # Sandboxed Dockerfile environments for code execution
│       ├── cpp/
│       ├── csharp/
│       ├── gdscript/
│       └── lua/
└── scripts/          # Shared CI/CD and deployment scripts
```

## 📜 Database Migrations
We use `sqlc` for type-safe database queries and standard SQL migration files. All SQL definitions and schemas are located in `apps/api/internal/db`.

*(Note: On Windows, it is highly recommended to run `sqlc generate` via Docker to avoid CGO cross-compilation errors).*

## 🤝 Contributing
GameCode is an open-source initiative to help the game development community master technical skills. We welcome contributions, from adding new problems and test cases to improving the execution engine and frontend UI.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License
Distributed under the MIT License. See `LICENSE` for more information.
