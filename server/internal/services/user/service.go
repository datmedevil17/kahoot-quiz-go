package user

import (
	"errors"

	"github.com/datmedevil17/kahoot-quiz-go/internal/models"
	"github.com/datmedevil17/kahoot-quiz-go/internal/utils"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

func (s *Service) CreateUser(email, name, password, role string) (*models.User, error) {
	var existingUser models.User
	if err := s.db.Where("email=?", email).First(&existingUser).Error; err == nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
		Role:         models.UserRole(role),
	}

	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) AuthenticateUser(email, password string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email=?", email).First(&user).Error; err == nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Record Not Found")
		}
	}
	if err := utils.CheckPassword(user.PasswordHash, password); err != nil {
		return nil, errors.New("Invalid Password")
	}
	return &user, nil
}

func (s *Service) GetUserById(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) GetUserProfile(id uint) (*models.User, error) {
	var user models.User
	if err := s.db.Preload("Followers").Preload("Following").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *Service) FollowUser(followerId, followingId uint) error {
	if followerId == followingId {
		return errors.New("You cannot follow yourself")
	}
	var existingFollow models.Follows
	err := s.db.Where("follower_id=? AND following_id=?", followerId, followingId).First(&existingFollow).Error
	if err == nil {
		return errors.New("You are already following this user")
	}
	follow := &models.Follows{
		FollowerID:  followerId,
		FollowingID: followingId,
	}
	if err := s.db.Create(follow).Error; err != nil {
		return err
	}
	return nil
}
