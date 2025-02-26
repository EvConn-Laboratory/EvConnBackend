package services

import (
	"context"
	"errors"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"
	"time"
)

type enrollmentService struct {
	*BaseService
	enrollmentRepo ports.EnrollmentRepository
	courseRepo     ports.CourseRepository
	userRepo       ports.UserRepository
}

func NewEnrollmentService(enrollmentRepo ports.EnrollmentRepository, courseRepo ports.CourseRepository, userRepo ports.UserRepository) ports.EnrollmentService {
	return &enrollmentService{
		BaseService:    NewBaseService(),
		enrollmentRepo: enrollmentRepo,
		courseRepo:     courseRepo,
		userRepo:       userRepo,
	}
}

func (s *enrollmentService) EnrollInCourse(ctx context.Context, studentID, courseID uint) (*models.Enrollment, error) {
	// Check if user exists and is a student
	user, err := s.userRepo.FindByID(ctx, studentID)
	if err != nil {
		return nil, errors.New("student not found")
	}

	if user.Role != "student" {
		return nil, errors.New("only students can enroll in courses")
	}

	// Check if course exists
	_, err = s.courseRepo.FindByID(ctx, courseID)
	if err != nil {
		return nil, errors.New("course not found")
	}

	// Check if already enrolled
	existing, err := s.enrollmentRepo.FindByStudentAndCourse(ctx, studentID, courseID)
	if err == nil && existing.ID > 0 {
		return existing, errors.New("already enrolled in this course")
	}

	// Create enrollment
	enrollment := &models.Enrollment{
		StudentID: studentID,
		CourseID:  courseID,
		Status:    models.EnrollmentStatusPending,
	}

	err = s.enrollmentRepo.Create(ctx, enrollment)
	if err != nil {
		return nil, err
	}

	// Reload to get relationships
	return s.enrollmentRepo.FindByID(ctx, enrollment.ID)
}

func (s *enrollmentService) GetEnrollment(ctx context.Context, id uint) (*models.Enrollment, error) {
	return s.enrollmentRepo.FindByID(ctx, id)
}

func (s *enrollmentService) GetEnrollmentsByStudent(ctx context.Context, studentID uint) ([]*models.Enrollment, error) {
	return s.enrollmentRepo.FindByStudent(ctx, studentID)
}

func (s *enrollmentService) GetEnrollmentsByCourse(ctx context.Context, courseID uint) ([]*models.Enrollment, error) {
	return s.enrollmentRepo.FindByCourse(ctx, courseID)
}

func (s *enrollmentService) UpdateEnrollmentStatus(ctx context.Context, id uint, status string, notes string) error {
	enrollment, err := s.enrollmentRepo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	enrollment.Status = status
	enrollment.Notes = notes

	if status == models.EnrollmentStatusApproved {
		now := time.Now()
		enrollment.ApprovedAt = &now
	}

	return s.enrollmentRepo.Update(ctx, enrollment)
}

func (s *enrollmentService) CancelEnrollment(ctx context.Context, id uint) error {
	return s.enrollmentRepo.Delete(ctx, id)
}
