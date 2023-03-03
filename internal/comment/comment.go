package comment

import "github.com/jinzhu/gorm"

// Service - the struct for our comment service
type Service struct {
	DB *gorm.DB
}

// Comment- defines the comment structure
type Comment struct {
	gorm.Model
	Slug    string
	Body    string
	Author  string
	Created string
}

// CommentService - the interface for our comment service
type CommentService interface {
	GetComment(ID int) (Comment, error)
	GetCommentsBySlug(slug string) ([]Comment, error)
	PostComment(comment Comment) (Comment, error)
	UpdateComment(ID uint, newComment Comment) (Comment, error)
	DeleteComment(ID uint) error
	GetAllComment() ([]Comment, error)
}

// GetComment - Retrives comments by their ID from the database
func (s *Service) GetComment(ID uint) (Comment, error) {

	var comment Comment
	if result := s.DB.First(&comment, ID); result != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// GetCommentsBySlug - retrieve all comments by slug (path /article/name/)
func (s *Service) GetCommentsBySlug(slug string) ([]Comment, error) {
	var comments []Comment
	if result := s.DB.Find(&comments).Where("slug = ?", slug); result.Error != nil {
		return []Comment{}, result.Error
	}
	return comments, nil
}

// PostComment - adds new comment to database
func (s *Service) PostComment(comment Comment) (Comment, error) {
	if result := s.DB.Save(&comment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// UpdateComment - updates comments by ID with new comment
func (s *Service) UpdateComment(ID uint, newComment Comment) (Comment, error) {
	comment, err := s.GetComment(ID)
	if err != nil {
		return Comment{}, err
	}
	if result := s.DB.Model(&comment).Updates(newComment); result.Error != nil {
		return Comment{}, result.Error
	}
	return comment, nil
}

// DeleteComment- Deletes the comment with the ID
func (s *Service) DeleteComment(ID uint) error {
	if result := s.DB.Delete(&Comment{}, ID); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAllComments - return all the list of comments from the database
func (s *Service) GetAllComment() ([]Comment, error) {
	var comments []Comment

	if result := s.DB.Find(&comments); result.Error != nil {
		return comments, result.Error
	}
	return comments, nil
}

// NewService - Returns the new comment service
func NewService(db *gorm.DB) *Service {
	return &Service{
		DB: db,
	}
}
