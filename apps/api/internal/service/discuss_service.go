package service

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "github.com/gc-platform/api/internal/repository"
)

type DiscussService interface {
    ListPostsByProblem(ctx context.Context, problemSlug string) ([]domain.DiscussPost, error)
    GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.DiscussPost, error)
    CreatePost(ctx context.Context, problemSlug string, userID uuid.UUID, title, content string) (*domain.DiscussPost, error)
    CreateComment(ctx context.Context, postID, userID uuid.UUID, content string) (*domain.DiscussComment, error)
    ListCommentsForPost(ctx context.Context, postID uuid.UUID) ([]domain.DiscussComment, error)
    TogglePostUpvote(ctx context.Context, postID, userID uuid.UUID) (bool, error)
    ToggleCommentUpvote(ctx context.Context, commentID, userID uuid.UUID) (bool, error)
}

type discussService struct {
    discussRepo repository.DiscussRepo
    problemRepo repository.ProblemRepo
}

func NewDiscussService(discussRepo repository.DiscussRepo, problemRepo repository.ProblemRepo) DiscussService {
    return &discussService{
        discussRepo: discussRepo,
        problemRepo: problemRepo,
    }
}

func (s *discussService) ListPostsByProblem(ctx context.Context, problemSlug string) ([]domain.DiscussPost, error) {
    problem, err := s.problemRepo.GetBySlug(ctx, problemSlug)
    if err != nil {
        return nil, err
    }
    return s.discussRepo.ListPostsByProblem(ctx, problem.ID)
}

func (s *discussService) GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.DiscussPost, error) {
    return s.discussRepo.GetPostByID(ctx, postID)
}

func (s *discussService) CreatePost(ctx context.Context, problemSlug string, userID uuid.UUID, title, content string) (*domain.DiscussPost, error) {
    problem, err := s.problemRepo.GetBySlug(ctx, problemSlug)
    if err != nil {
        return nil, err
    }

    post := &domain.DiscussPost{
        ProblemID: problem.ID,
        UserID:    userID,
        Title:     title,
        Content:   content,
    }

    if err := s.discussRepo.CreatePost(ctx, post); err != nil {
        return nil, err
    }
    return post, nil
}

func (s *discussService) CreateComment(ctx context.Context, postID, userID uuid.UUID, content string) (*domain.DiscussComment, error) {
    comment := &domain.DiscussComment{
        PostID:  postID,
        UserID:  userID,
        Content: content,
    }

    if err := s.discussRepo.CreateComment(ctx, comment); err != nil {
        return nil, err
    }
    return comment, nil
}

func (s *discussService) ListCommentsForPost(ctx context.Context, postID uuid.UUID) ([]domain.DiscussComment, error) {
    return s.discussRepo.ListCommentsForPost(ctx, postID)
}

func (s *discussService) TogglePostUpvote(ctx context.Context, postID, userID uuid.UUID) (bool, error) {
    return s.discussRepo.TogglePostUpvote(ctx, postID, userID)
}

func (s *discussService) ToggleCommentUpvote(ctx context.Context, commentID, userID uuid.UUID) (bool, error) {
    return s.discussRepo.ToggleCommentUpvote(ctx, commentID, userID)
}
