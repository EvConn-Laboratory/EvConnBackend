package ports

import (
	"context"
	"evconn/internal/core/domain/models"
	"io"
)

type CourseService interface {
	CreateCourse(ctx context.Context, course *models.Course) error
	GetCourse(ctx context.Context, id uint) (*models.Course, error)
	GetAllCourses(ctx context.Context) ([]*models.Course, error)
	GetCoursesByMentor(ctx context.Context, mentorID uint) ([]*models.Course, error)
	UpdateCourse(ctx context.Context, course *models.Course) error
	DeleteCourse(ctx context.Context, id uint) error
	SearchCourses(ctx context.Context, query string) ([]*models.Course, error)
}

type ModuleService interface {
	CreateModule(ctx context.Context, module *models.Module) error
	GetModule(ctx context.Context, id uint) (*models.Module, error)
	GetModulesByCourse(ctx context.Context, courseID uint) ([]models.Module, error)
	UpdateModule(ctx context.Context, module *models.Module) error
	DeleteModule(ctx context.Context, id uint) error
	AssignMentor(ctx context.Context, moduleID, mentorID uint) error
}

type SubmissionService interface {
	CreateSubmission(ctx context.Context, submission *models.Submission) error
	GetSubmission(ctx context.Context, id uint) (*models.Submission, error)
	GetSubmissionsByAssignment(ctx context.Context, assignmentID uint) ([]models.Submission, error)
	GetSubmissionsByUser(ctx context.Context, userID uint) ([]models.Submission, error)
	GradeSubmission(ctx context.Context, id uint, score float64) error
}

type FeedbackService interface {
	CreateFeedback(ctx context.Context, feedback *models.Feedback) error
	GetFeedbackByID(ctx context.Context, id uint) (*models.Feedback, error)
	GetFeedbackByModule(ctx context.Context, moduleID uint) ([]models.Feedback, error)
	GetFeedbackByUser(ctx context.Context, userID uint) ([]models.Feedback, error)
	GetFeedbackByModuleAndType(ctx context.Context, moduleID uint, feedbackType string) ([]models.Feedback, error)
	UpdateFeedback(ctx context.Context, feedback *models.Feedback) error
	DeleteFeedback(ctx context.Context, id uint) error
}

type AssignmentService interface {
	GetAssignment(ctx context.Context, id uint) (*models.Assignment, error)
	GetAssignmentsByModule(ctx context.Context, moduleID uint) ([]models.Assignment, error)
	GetPendingAssignments(ctx context.Context, userID uint) ([]*models.Assignment, error)
	CreateAssignment(ctx context.Context, assignment *models.Assignment) error
	UpdateAssignment(ctx context.Context, assignment *models.Assignment) error
	DeleteAssignment(ctx context.Context, id uint) error
}

type UserService interface {
	Create(ctx context.Context, user *models.User) error
	GetByID(ctx context.Context, id uint) (*models.User, error)
	GetByNIM(ctx context.Context, nim string) (*models.User, error)
	GetByLab(ctx context.Context, lab string) ([]*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	GetAllUsersPaginated(ctx context.Context, page, pageSize int) ([]*models.User, int64, error)
	ImportUsers(ctx context.Context, users []*models.User) ([]string, error)
}

type LabService interface {
	GetLabs(ctx context.Context) ([]*models.Lab, error)
	GetLabByID(ctx context.Context, id uint) (*models.Lab, error)
	CreateLab(ctx context.Context, lab *models.Lab) error
	UpdateLab(ctx context.Context, lab *models.Lab) error
	DeleteLab(ctx context.Context, id uint) error
}

type EnrollmentService interface {
	EnrollInCourse(ctx context.Context, studentID, courseID uint) (*models.Enrollment, error)
	GetEnrollment(ctx context.Context, id uint) (*models.Enrollment, error)
	GetEnrollmentsByStudent(ctx context.Context, studentID uint) ([]*models.Enrollment, error)
	GetEnrollmentsByCourse(ctx context.Context, courseID uint) ([]*models.Enrollment, error)
	UpdateEnrollmentStatus(ctx context.Context, id uint, status string, notes string) error
	CancelEnrollment(ctx context.Context, id uint) error
}

type FileService interface {
	UploadFile(ctx context.Context, file *models.File, reader io.Reader) error
	GetFileURL(ctx context.Context, id uint) (string, error)
	GetFileByID(ctx context.Context, id uint) (*models.File, error)
	GetFilesByEntity(ctx context.Context, entityType string, entityID uint) ([]*models.File, error)
	DeleteFile(ctx context.Context, id uint) error
}
