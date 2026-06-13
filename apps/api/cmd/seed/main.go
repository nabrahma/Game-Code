package main

import (
    "context"
    "fmt"
    "log"

    "github.com/gc-platform/api/internal/config"
    "github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    cfg := config.Load()

    pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }
    defer pool.Close()

    ctx := context.Background()

    // 1. Insert Tags
    fmt.Println("Seeding Tags...")
    tagQueries := []string{
        `INSERT INTO tags (name, slug, category, order_index) VALUES ('Pathfinding', 'pathfinding', 'AI', 1) ON CONFLICT (name) DO NOTHING`,
        `INSERT INTO tags (name, slug, category, order_index) VALUES ('Math', 'math', 'Math', 2) ON CONFLICT (name) DO NOTHING`,
        `INSERT INTO tags (name, slug, category, order_index) VALUES ('Physics', 'physics', 'Engine', 3) ON CONFLICT (name) DO NOTHING`,
        `INSERT INTO tags (name, slug, category, order_index) VALUES ('Data Structures', 'data-structures', 'Computer Science', 4) ON CONFLICT (name) DO NOTHING`,
    }
    for _, q := range tagQueries {
        _, err = pool.Exec(ctx, q)
        if err != nil {
            log.Fatal("failed to insert tag:", err)
        }
    }

    // 2. Insert Problems
    fmt.Println("Seeding Problems...")
    problemQueries := []struct {
        Slug        string
        Title       string
        Difficulty  string
        Status      string
        Desc        string
        Constraints string
    }{
        {"a-star-pathfinding", "A* Pathfinding Implementation", "hard", "published", "Implement the A* algorithm on a grid.", "Grid max 100x100"},
        {"vector-normalization", "Vector Normalization", "easy", "published", "Normalize a 3D vector.", "Float precision 0.0001"},
        {"aabb-collision", "AABB Collision Detection", "easy", "published", "Detect if two Axis-Aligned Bounding Boxes intersect.", "Box coordinates are integers"},
        {"object-pool", "Object Pool Implementation", "medium", "published", "Implement a generic Object Pool to prevent garbage collection spikes.", "Pool size up to 10,000"},
        {"quadtree-insertion", "Quadtree Insertion", "medium", "published", "Insert an object into a spatial partitioning Quadtree.", "Max objects per node is 4"},
    }

    for _, p := range problemQueries {
        query := `INSERT INTO problems (slug, title, difficulty, status, description, constraints) 
                  VALUES ($1, $2, $3, $4, $5, $6) 
                  ON CONFLICT (slug) DO NOTHING`
        _, err = pool.Exec(ctx, query, p.Slug, p.Title, p.Difficulty, p.Status, p.Desc, p.Constraints)
        if err != nil {
            log.Fatal("failed to insert problem:", err)
        }
    }

    // Note: To fully map problem_tags we would query the inserted UUIDs, but for the sake of simplicity, 
    // the UI will still display these problems in the list view without tags or we can join them later.
    
    fmt.Println("Database seeding completed successfully!")
}
