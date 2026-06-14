package repository

import (
    "context"

    "github.com/google/uuid"
    "github.com/gc-platform/api/internal/domain"
    "gorm.io/gorm"
)

type DiscussRepo interface {
    ListPostsByProblem(ctx context.Context, problemID uuid.UUID) ([]domain.DiscussPost, error)
    GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.DiscussPost, error)
    CreatePost(ctx context.Context, post *domain.DiscussPost) error
    CreateComment(ctx context.Context, comment *domain.DiscussComment) error
    ListCommentsForPost(ctx context.Context, postID uuid.UUID) ([]domain.DiscussComment, error)
    TogglePostUpvote(ctx context.Context, postID, userID uuid.UUID) (bool, error)
    ToggleCommentUpvote(ctx context.Context, commentID, userID uuid.UUID) (bool, error)
}

type discussRepo struct {
    db *gorm.DB
}

func NewDiscussRepo(db *gorm.DB) DiscussRepo {
    return &discussRepo{db: db}
}

func (r *discussRepo) ListPostsByProblem(ctx context.Context, problemID uuid.UUID) ([]domain.DiscussPost, error) {
    var posts []domain.DiscussPost
    err := r.db.WithContext(ctx).
        Preload("User").
        Where("problem_id = ?", problemID).
        Order("created_at desc").
        Find(&posts).Error
    return posts, err
}

func (r *discussRepo) GetPostByID(ctx context.Context, postID uuid.UUID) (*domain.DiscussPost, error) {
    var post domain.DiscussPost
    err := r.db.WithContext(ctx).
        Preload("User").
        Where("id = ?", postID).
        First(&post).Error
    if err != nil {
        return nil, err
    }
    return &post, nil
}

func (r *discussRepo) CreatePost(ctx context.Context, post *domain.DiscussPost) error {
    return r.db.WithContext(ctx).Create(post).Error
}

func (r *discussRepo) CreateComment(ctx context.Context, comment *domain.DiscussComment) error {
    return r.db.WithContext(ctx).Create(comment).Error
}

func (r *discussRepo) ListCommentsForPost(ctx context.Context, postID uuid.UUID) ([]domain.DiscussComment, error) {
    var comments []domain.DiscussComment
    err := r.db.WithContext(ctx).
        Preload("User").
        Where("post_id = ?", postID).
        Order("created_at asc").
        Find(&comments).Error
    return comments, err
}

func (r *discussRepo) TogglePostUpvote(ctx context.Context, postID, userID uuid.UUID) (bool, error) {
    var count int64
    r.db.WithContext(ctx).Model(&domain.PostUpvote{}).
        Where("post_id = ? AND user_id = ?", postID, userID).Count(&count)

    if count > 0 {
        // Remove upvote
        err := r.db.WithContext(ctx).
            Where("post_id = ? AND user_id = ?", postID, userID).
            Delete(&domain.PostUpvote{}).Error
        if err == nil {
            r.db.WithContext(ctx).Model(&domain.DiscussPost{}).Where("id = ?", postID).UpdateColumn("upvotes", gorm.Expr("upvotes - 1"))
        }
        return false, err
    } else {
        // Add upvote
        err := r.db.WithContext(ctx).Create(&domain.PostUpvote{
            PostID: postID,
            UserID: userID,
        }).Error
        if err == nil {
            r.db.WithContext(ctx).Model(&domain.DiscussPost{}).Where("id = ?", postID).UpdateColumn("upvotes", gorm.Expr("upvotes + 1"))
        }
        return true, err
    }
}

func (r *discussRepo) ToggleCommentUpvote(ctx context.Context, commentID, userID uuid.UUID) (bool, error) {
    var count int64
    r.db.WithContext(ctx).Model(&domain.CommentUpvote{}).
        Where("comment_id = ? AND user_id = ?", commentID, userID).Count(&count)

    if count > 0 {
        // Remove upvote
        err := r.db.WithContext(ctx).
            Where("comment_id = ? AND user_id = ?", commentID, userID).
            Delete(&domain.CommentUpvote{}).Error
        if err == nil {
            r.db.WithContext(ctx).Model(&domain.DiscussComment{}).Where("id = ?", commentID).UpdateColumn("upvotes", gorm.Expr("upvotes - 1"))
        }
        return false, err
    } else {
        // Add upvote
        err := r.db.WithContext(ctx).Create(&domain.CommentUpvote{
            CommentID: commentID,
            UserID:    userID,
        }).Error
        if err == nil {
            r.db.WithContext(ctx).Model(&domain.DiscussComment{}).Where("id = ?", commentID).UpdateColumn("upvotes", gorm.Expr("upvotes + 1"))
        }
        return true, err
    }
}
