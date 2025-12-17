package service

import (
	"context"
	"time"

	"user-api/internal/models"
	"user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	parsedDob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.CreateUser(ctx, req.Name, parsedDob)
	if err != nil {
		return models.UserResponse{}, err
	}

	return toUserResponse(user), nil
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}

	return toUserResponse(user), nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]models.UserResponse, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var res []models.UserResponse
	for _, u := range users {
		res = append(res, toUserResponse(u))
	}

	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	parsedDob, err := time.Parse("2006-01-02", req.Dob)
	if err != nil {
		return models.UserResponse{}, err
	}

	user, err := s.repo.UpdateUser(ctx, id, req.Name, parsedDob)
	if err != nil {
		return models.UserResponse{}, err
	}

	return toUserResponse(user), nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}

// helper function to convert User to UserResponse with calculated age
func toUserResponse(u models.User) models.UserResponse {
	return models.UserResponse{
		ID:   u.ID,
		Name: u.Name,
		Dob:  u.Dob.Format("2006-01-02"),
		Age:  calculateAge(u.Dob),
	}
}

// calculateAge calculates age from date of birth
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Check if birthday hasn't occurred this year yet
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}

	return age
}
