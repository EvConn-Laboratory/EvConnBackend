package repositories

import (
	"context"
	"evconn/internal/core/domain/models"
	"evconn/internal/core/ports"

	"gorm.io/gorm"
)

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) ports.CourseRepository {
	return &courseRepository{db: db}
}

func (r *courseRepository) Create(ctx context.Context, course *models.Course) error {
	return r.db.WithContext(ctx).Create(course).Error
}

func (r *courseRepository) FindByID(ctx context.Context, id uint) (*models.Course, error) {
	var course models.Course
	err := r.db.WithContext(ctx).Preload("Mentor").Preload("Modules").First(&course, id).Error
	return &course, err
}

func (r *courseRepository) FindAll(ctx context.Context) ([]*models.Course, error) {
	var courses []*models.Course
	err := r.db.WithContext(ctx).Preload("Mentor").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindByMentorID(ctx context.Context, mentorID uint) ([]*models.Course, error) {
	var courses []*models.Course
	err := r.db.WithContext(ctx).Where("mentor_id = ?", mentorID).Find(&courses).Error
	return courses, err
}

func (r *courseRepository) Update(ctx context.Context, course *models.Course) error {
	return r.db.WithContext(ctx).Save(course).Error
}

func (r *courseRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Course{}, id).Error
}

func (r *courseRepository) Search(ctx context.Context, query string) ([]*models.Course, error) {
	var courses []*models.Course
	err := r.db.WithContext(ctx).Where("name LIKE ? OR description LIKE ?", "%"+query+"%", "%"+query+"%").Find(&courses).Error
	return courses, err
}
