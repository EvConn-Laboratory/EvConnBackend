package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type enrollmentRepository struct {
	db *gorm.DB
}

func NewEnrollmentRepository(db *gorm.DB) ports.EnrollmentRepository {
	return &enrollmentRepository{db: db}
}

func (r *enrollmentRepository) Create(ctx context.Context, enrollment *models.Enrollment) error {
	return r.db.WithContext(ctx).Create(enrollment).Error
}

func (r *enrollmentRepository) FindByID(ctx context.Context, id uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.WithContext(ctx).
		Preload("Student").
		Preload("Course").
		First(&enrollment, id).Error
	return &enrollment, err
}

func (r *enrollmentRepository) FindByStudentAndCourse(ctx context.Context, studentID, courseID uint) (*models.Enrollment, error) {
	var enrollment models.Enrollment
	err := r.db.WithContext(ctx).
		Where("student_id = ? AND course_id = ?", studentID, courseID).
		Preload("Student").
		Preload("Course").
		First(&enrollment).Error
	return &enrollment, err
}

func (r *enrollmentRepository) FindByCourse(ctx context.Context, courseID uint) ([]*models.Enrollment, error) {
	var enrollments []*models.Enrollment
	err := r.db.WithContext(ctx).
		Where("course_id = ?", courseID).
		Preload("Student").
		Find(&enrollments).Error
	return enrollments, err
}

func (r *enrollmentRepository) FindByStudent(ctx context.Context, studentID uint) ([]*models.Enrollment, error) {
	var enrollments []*models.Enrollment
	err := r.db.WithContext(ctx).
		Where("student_id = ?", studentID).
		Preload("Course").
		Find(&enrollments).Error
	return enrollments, err
}

func (r *enrollmentRepository) Update(ctx context.Context, enrollment *models.Enrollment) error {
	return r.db.WithContext(ctx).Save(enrollment).Error
}

func (r *enrollmentRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Enrollment{}, id).Error
}
