package executor

import (
    "bytes"
    "context"
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "strings"
    "time"
)

type DockerRunner struct{}

func NewDockerRunner() *DockerRunner {
    return &DockerRunner{}
}

func (r *DockerRunner) Run(ctx context.Context, language string, code string, input string) (string, error) {
    // Determine image based on language
    imageMap := map[string]string{
        "cpp":    "gc-runner-cpp",
        "csharp": "gc-runner-csharp",
    }
    
    image, ok := imageMap[language]
    if !ok {
        return "", fmt.Errorf("unsupported language: %s", language)
    }

    // Create a temporary directory for this execution
    tempDir, err := os.MkdirTemp("", "gc-run-*")
    if err != nil {
        return "", fmt.Errorf("failed to create temp directory: %v", err)
    }
    defer os.RemoveAll(tempDir) // cleanup

    // Write code to file
    codeFile := filepath.Join(tempDir, getFilename(language))
    if err := os.WriteFile(codeFile, []byte(code), 0644); err != nil {
        return "", fmt.Errorf("failed to write code file: %v", err)
    }

    // Command configuration for strict execution
    // Memory limit: 256m, No network, No privileges
    args := []string{
        "run", "--rm",
        "--network", "none",
        "--memory", "256m",
        "-v", fmt.Sprintf("%s:/sandbox/%s:ro", codeFile, getFilename(language)),
        image,
    }

    // The context will ensure the command is killed if it exceeds 3 seconds
    timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
    defer cancel()

    cmd := exec.CommandContext(timeoutCtx, "docker", args...)
    
    var outBuf bytes.Buffer
    cmd.Stdout = &outBuf
    cmd.Stderr = &outBuf
    
    if input != "" {
        cmd.Stdin = strings.NewReader(input)
    }

    err = cmd.Run()
    
    // Check if context timed out
    if timeoutCtx.Err() == context.DeadlineExceeded {
        return outBuf.String(), fmt.Errorf("execution timed out")
    }

    if err != nil {
        return outBuf.String(), fmt.Errorf("execution failed: %v", err)
    }

    return outBuf.String(), nil
}

func getFilename(lang string) string {
    switch lang {
    case "cpp":
        return "Solution.cpp"
    case "csharp":
        return "Solution.cs"
    default:
        return "Solution.txt"
    }
}
