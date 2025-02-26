package services

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
)

type courseService struct {
	*BaseService
	courseRepo ports.CourseRepository
}

func NewCourseService(courseRepo ports.CourseRepository) ports.CourseService {
	return &courseService{
		BaseService: NewBaseService(),
		courseRepo:  courseRepo,
	}
}

func (s *courseService) CreateCourse(ctx context.Context, course *models.Course) error {
	return s.courseRepo.Create(ctx, course)
}

func (s *courseService) GetCourse(ctx context.Context, id uint) (*models.Course, error) {
	return s.courseRepo.FindByID(ctx, id)
}

// Update the return type to match the interface
func (s *courseService) GetAllCourses(ctx context.Context) ([]*models.Course, error) {
	return s.courseRepo.FindAll(ctx)
}

// Update the return type to match the interface
func (s *courseService) GetCoursesByMentor(ctx context.Context, mentorID uint) ([]*models.Course, error) {
	return s.courseRepo.FindByMentorID(ctx, mentorID)
}

func (s *courseService) UpdateCourse(ctx context.Context, course *models.Course) error {
	return s.courseRepo.Update(ctx, course)
}

func (s *courseService) DeleteCourse(ctx context.Context, id uint) error {
	return s.courseRepo.Delete(ctx, id)
}

// Add this method to your courseService struct
func (s *courseService) SearchCourses(ctx context.Context, query string) ([]*models.Course, error) {
	return s.courseRepo.Search(ctx, query)
}
