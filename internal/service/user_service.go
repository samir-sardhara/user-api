package service

import (
	"context"
	"time"
	"user-api/db/sqlc"
	"user-api/internal/models"
	"user-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func calculateAge(dob time.Time, now time.Time) int {
	age := now.Year() - dob.Year()
	if now.Month() < dob.Month() || (now.Month() == dob.Month() && now.Day() < dob.Day()) {
		age--
	}
	return age
}

func (s *UserService) CreateUser(ctx context.Context, req models.CreateUserRequest) (models.UserResponse, error) {
	dob, _ := time.Parse("2006-01-02", req.DOB)

	result, err := s.repo.Queries.CreateUser(ctx, sqlc.CreateUserParams{
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{ID: int32(id), Name: req.Name, DOB: req.DOB}, nil
}

func (s *UserService) GetUser(ctx context.Context, id int32) (models.UserResponse, error) {
	user, err := s.repo.Queries.GetUser(ctx, id)
	if err != nil {
		return models.UserResponse{}, err
	}
	age := calculateAge(user.Dob, time.Now())
	return models.UserResponse{ID: user.ID, Name: user.Name, DOB: user.Dob.Format("2006-01-02"), Age: &age}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id int32, req models.UpdateUserRequest) (models.UserResponse, error) {
	dob, _ := time.Parse("2006-01-02", req.DOB)

	err := s.repo.Queries.UpdateUser(ctx, sqlc.UpdateUserParams{
		ID:   id,
		Name: req.Name,
		Dob:  dob,
	})
	if err != nil {
		return models.UserResponse{}, err
	}

	return models.UserResponse{ID: id, Name: req.Name, DOB: req.DOB}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.Queries.DeleteUser(ctx, id)
}

func (s *UserService) ListUsers(ctx context.Context, page int32, limit int32) ([]models.UserResponse, error) {
	offset := (page - 1) * limit
	users, err := s.repo.Queries.ListUsers(ctx, sqlc.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	var res []models.UserResponse
	for _, user := range users {
		age := calculateAge(user.Dob, time.Now()) // Notice we pass time.Now() here now for better testing
		res = append(res, models.UserResponse{ID: user.ID, Name: user.Name, DOB: user.Dob.Format("2006-01-02"), Age: &age})
	}
	if res == nil {
		res = []models.UserResponse{}
	}
	return res, nil
}
