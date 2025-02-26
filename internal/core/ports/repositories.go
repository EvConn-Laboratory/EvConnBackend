package ports

import (
	"context"
	"evconn/internal/core/domain/models"
)

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindByID(ctx context.Context, id uint) (*models.User, error)
	FindByNIM(ctx context.Context, nim string) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
	FindByLab(ctx context.Context, lab string) ([]*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id uint) error
	FindAllPaginated(ctx context.Context, page, pageSize int) ([]*models.User, int64, error)
}

type AssignmentRepository interface {
	Create(ctx context.Context, assignment *models.Assignment) error
	FindByID(ctx context.Context, id uint) (*models.Assignment, error)
	FindByModuleID(ctx context.Context, moduleID uint) ([]models.Assignment, error)
	Update(ctx context.Context, assignment *models.Assignment) error
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]models.Assignment, error)
}

type CourseRepository interface {
	Create(ctx context.Context, course *models.Course) error
	FindByID(ctx context.Context, id uint) (*models.Course, error)
	FindAll(ctx context.Context) ([]*models.Course, error)
	FindByMentorID(ctx context.Context, mentorID uint) ([]*models.Course, error)
	Update(ctx context.Context, course *models.Course) error
	Delete(ctx context.Context, id uint) error
	Search(ctx context.Context, query string) ([]*models.Course, error) // Add this method
}

type ModuleRepository interface {
	Create(ctx context.Context, module *models.Module) error
	FindByID(ctx context.Context, id uint) (*models.Module, error)
	FindByCourseID(ctx context.Context, courseID uint) ([]models.Module, error)
	Update(ctx context.Context, module *models.Module) error
	Delete(ctx context.Context, id uint) error
}

type ModuleMentorRepository interface {
	Create(ctx context.Context, moduleMentor *models.ModuleMentor) error
	FindByID(ctx context.Context, id uint) (*models.ModuleMentor, error)
	FindByModuleID(ctx context.Context, moduleID uint) ([]models.ModuleMentor, error)
	FindByMentorID(ctx context.Context, mentorID uint) ([]models.ModuleMentor, error)
	Delete(ctx context.Context, id uint) error
}

type SubmissionRepository interface {
	Create(ctx context.Context, submission *models.Submission) error
	FindByID(ctx context.Context, id uint) (*models.Submission, error)
	FindByAssignmentID(ctx context.Context, assignmentID uint) ([]models.Submission, error)
	FindByUserID(ctx context.Context, userID uint) ([]models.Submission, error)
	Update(ctx context.Context, submission *models.Submission) error
	Delete(ctx context.Context, id uint) error
}

type FeedbackRepository interface {
	Create(ctx context.Context, feedback *models.Feedback) error
	FindByID(ctx context.Context, id uint) (*models.Feedback, error)
	FindByUserID(ctx context.Context, userID uint) ([]models.Feedback, error)
	FindByModuleID(ctx context.Context, moduleID uint) ([]models.Feedback, error)
	FindByModuleAndType(ctx context.Context, moduleID uint, feedbackType string) ([]models.Feedback, error)
	Update(ctx context.Context, feedback *models.Feedback) error
	Delete(ctx context.Context, id uint) error
}

type LabRepository interface {
	Create(ctx context.Context, lab *models.Lab) error
	FindAll(ctx context.Context) ([]*models.Lab, error)
	FindByID(ctx context.Context, id uint) (*models.Lab, error)
	Update(ctx context.Context, lab *models.Lab) error
	Delete(ctx context.Context, id uint) error
}

type EnrollmentRepository interface {
	Create(ctx context.Context, enrollment *models.Enrollment) error
	FindByID(ctx context.Context, id uint) (*models.Enrollment, error)
	FindByCourse(ctx context.Context, courseID uint) ([]*models.Enrollment, error)
	FindByStudent(ctx context.Context, studentID uint) ([]*models.Enrollment, error)
	FindByStudentAndCourse(ctx context.Context, studentID, courseID uint) (*models.Enrollment, error)
	Update(ctx context.Context, enrollment *models.Enrollment) error
	Delete(ctx context.Context, id uint) error
}

// Update the FileRepository interface definition

type FileRepository interface {
	Create(ctx context.Context, file *models.File) error
	FindByID(ctx context.Context, id uint) (*models.File, error)
	FindByEntityTypeAndID(ctx context.Context, entityType string, entityID uint) ([]*models.File, error) // This method name
	FindByUserID(ctx context.Context, userID uint) ([]*models.File, error)
	Update(ctx context.Context, file *models.File) error
	Delete(ctx context.Context, id uint) error
}
