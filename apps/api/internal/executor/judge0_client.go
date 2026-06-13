package executor

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "net/http"
    "time"
    "github.com/gc-platform/api/internal/domain"
    "github.com/google/uuid"
)

type Judge0Client interface {
    Submit(ctx context.Context, language string, code string, stdin string) (string, error)
    PollStatus(ctx context.Context, token string) (*Judge0Submission, error)
    EvaluateSubmission(ctx context.Context, language string, code string, testCases []domain.TestCase) ([]domain.TestResult, domain.SubmissionVerdict, error)
}

type judge0Client struct {
    baseURL    string
    httpClient *http.Client
}

type Judge0Submission struct {
    Stdout        string `json:"stdout"`
    Stderr        string `json:"stderr"`
    CompileOutput string `json:"compile_output"`
    Message       string `json:"message"`
    Time          string `json:"time"`
    Memory        int    `json:"memory"`
    Status        struct {
        ID          int    `json:"id"`
        Description string `json:"description"`
    } `json:"status"`
}

func NewJudge0Client() Judge0Client {
    return &judge0Client{
        baseURL: "https://ce.judge0.com",
        httpClient: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (c *judge0Client) getLanguageID(lang string) int {
    // Judge0 CE language IDs
    switch lang {
    case "cpp":
        return 54 // C++ (GCC 9.2.0)
    case "csharp":
        return 51 // C# (Mono 6.6.0.161)
    case "python": // fallback for gdscript
        return 71 // Python (3.8.1)
    default:
        return 71
    }
}

func (c *judge0Client) Submit(ctx context.Context, language string, code string, stdin string) (string, error) {
    langID := c.getLanguageID(language)
    
    payload := map[string]interface{}{
        "source_code": code,
        "language_id": langID,
        "stdin":       stdin,
    }
    
    body, _ := json.Marshal(payload)
    
    req, err := http.NewRequestWithContext(ctx, "POST", c.baseURL+"/submissions?base64_encoded=false&wait=false", bytes.NewBuffer(body))
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")
    
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != 201 {
        return "", fmt.Errorf("failed to submit to Judge0, status: %d", resp.StatusCode)
    }
    
    var result struct {
        Token string `json:"token"`
    }
    if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        return "", err
    }
    
    return result.Token, nil
}

func (c *judge0Client) PollStatus(ctx context.Context, token string) (*Judge0Submission, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", c.baseURL+"/submissions/"+token+"?base64_encoded=false", nil)
    if err != nil {
        return nil, err
    }
    
    resp, err := c.httpClient.Do(req)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()
    
    var sub Judge0Submission
    if err := json.NewDecoder(resp.Body).Decode(&sub); err != nil {
        return nil, err
    }
    
    return &sub, nil
}

func (c *judge0Client) EvaluateSubmission(ctx context.Context, language string, code string, testCases []domain.TestCase) ([]domain.TestResult, domain.SubmissionVerdict, error) {
    var results []domain.TestResult
    finalVerdict := domain.VerdictAccepted

    for _, tc := range testCases {
        // 1. Submit
        token, err := c.Submit(ctx, language, code, tc.Input)
        if err != nil {
            return results, domain.VerdictInternalError, err
        }

        // 2. Poll
        var sub *Judge0Submission
        for i := 0; i < 20; i++ { // Poll for max 20 seconds per testcase
            time.Sleep(1 * time.Second)
            sub, err = c.PollStatus(ctx, token)
            if err == nil && sub.Status.ID > 2 {
                break
            }
        }

        if sub == nil || sub.Status.ID <= 2 {
            finalVerdict = domain.VerdictTimeLimitExceeded
            break
        }

        result := domain.TestResult{
            ID:        uuid.New(),
            Input:     tc.Input,
            Expected:  tc.Output,
            Actual:    sub.Stdout,
            Passed:    sub.Stdout == tc.Output,
        }

        if sub.Status.ID == 3 && result.Passed {
            // Passed!
        } else if sub.Status.ID == 3 && !result.Passed {
            finalVerdict = domain.VerdictWrongAnswer
        } else if sub.Status.ID == 6 {
            finalVerdict = domain.VerdictCompileError
            result.Actual = sub.CompileOutput
        } else if sub.Status.ID == 5 {
            finalVerdict = domain.VerdictTimeLimitExceeded
            result.Actual = "Time Limit Exceeded"
        } else {
            finalVerdict = domain.VerdictRuntimeError
            result.Actual = sub.Stderr
        }

        results = append(results, result)

        // Break early on first failure to save execution time
        if finalVerdict != domain.VerdictAccepted {
            break
        }
    }

    return results, finalVerdict, nil
}
