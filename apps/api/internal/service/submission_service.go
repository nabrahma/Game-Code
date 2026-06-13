package service

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/cache"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/executor"
    "github.com/gc-platform/api/internal/repository"
)

type SubmissionService interface {
    Create(ctx context.Context, userID uuid.UUID, problemID uuid.UUID, language string, code string) (*domain.Submission, error)
    GetByID(ctx context.Context, id uuid.UUID) (*domain.Submission, error)
}

type submissionService struct {
    subRepo  repository.SubmissionRepo
    probRepo repository.ProblemRepo
    userRepo repository.UserRepo
    cache    cache.Cache
    judge0   executor.Judge0Client
}

func NewSubmissionService(subRepo repository.SubmissionRepo, probRepo repository.ProblemRepo, userRepo repository.UserRepo, cache cache.Cache) SubmissionService {
    return &submissionService{
        subRepo:  subRepo,
        probRepo: probRepo,
        userRepo: userRepo,
        cache:    cache,
        judge0:   executor.NewJudge0Client(),
    }
}

func (s *submissionService) Create(ctx context.Context, userID uuid.UUID, problemID uuid.UUID, language string, code string) (*domain.Submission, error) {
    sub := &domain.Submission{
        ID:        uuid.New(),
        UserID:    userID,
        ProblemID: problemID,
        Language:  domain.Language(language),
        Code:      code,
        Verdict:   domain.VerdictPending,
    }

    if err := s.subRepo.Create(ctx, sub); err != nil {
        return nil, err
    }

    // Fire and forget evaluation
    go s.evaluateSubmission(userID, sub.ID, problemID, language, code)

    return sub, nil
}

func (s *submissionService) evaluateSubmission(userID uuid.UUID, subID uuid.UUID, problemID uuid.UUID, language string, code string) {
    ctx := context.Background() // New context since HTTP request may finish early

    // 1. Fetch test cases (assuming a mock for now since we don't have a test case repository yet)
    // Normally: testCases, _ := s.probRepo.GetTestCases(ctx, problemID)
    testCases := []domain.TestCase{
        {Input: "1\n2\n", Output: "3\n"},
        {Input: "5\n10\n", Output: "15\n"},
    }

    // 2. Evaluate
    results, finalVerdict, err := s.judge0.EvaluateSubmission(ctx, language, code, testCases)

    errorMsg := ""
    if err != nil {
        finalVerdict = domain.VerdictInternalError
        errorMsg = err.Error()
    }

    // Calculate metrics
    passed := int32(0)
    for _, res := range results {
        if res.Passed {
            passed++
        }
    }

    // 3. Update Submission DB Row
    _ = s.subRepo.UpdateVerdict(ctx, subID, finalVerdict, 0, 0, errorMsg, passed, int32(len(testCases)))

    // 4. Upsert User Progress
    status := "attempted"
    if finalVerdict == domain.VerdictAccepted {
        status = "solved"
    }
    _ = s.userRepo.UpsertProgress(ctx, userID, problemID, status)
}

func (s *submissionService) GetByID(ctx context.Context, id uuid.UUID) (*domain.Submission, error) {
    return s.subRepo.GetByID(ctx, id)
}
