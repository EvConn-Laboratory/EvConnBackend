package repositories

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	db *gorm.DB
}

func NewBaseRepository[T any](db *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: db,
	}
}

func (r *BaseRepository[T]) DB(ctx context.Context) *gorm.DB {
	return r.db.WithContext(ctx)
}

// Create inserts a new record into the database
func (r *BaseRepository[T]) Create(ctx context.Context, entity *T) error {
	return r.DB(ctx).Create(entity).Error
}

// FindByID retrieves a record by its ID
func (r *BaseRepository[T]) FindByID(ctx context.Context, id uint) (*T, error) {
	var entity T
	err := r.DB(ctx).First(&entity, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &entity, nil
}

// FindAll retrieves all records
func (r *BaseRepository[T]) FindAll(ctx context.Context) ([]T, error) {
	var entities []T
	err := r.DB(ctx).Find(&entities).Error
	return entities, err
}

// Update updates an existing record
func (r *BaseRepository[T]) Update(ctx context.Context, entity *T) error {
	return r.DB(ctx).Save(entity).Error
}

// Delete removes a record by its ID
func (r *BaseRepository[T]) Delete(ctx context.Context, id uint) error {
	return r.DB(ctx).Delete(new(T), id).Error
}

// FindOneBy retrieves a single record by a condition
func (r *BaseRepository[T]) FindOneBy(ctx context.Context, condition interface{}) (*T, error) {
	var entity T
	err := r.DB(ctx).Where(condition).First(&entity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("record not found")
		}
		return nil, err
	}
	return &entity, nil
}

// FindAllBy retrieves all records matching a condition
func (r *BaseRepository[T]) FindAllBy(ctx context.Context, condition interface{}) ([]T, error) {
	var entities []T
	err := r.DB(ctx).Where(condition).Find(&entities).Error
	return entities, err
}

// Transaction executes operations within a database transaction
func (r *BaseRepository[T]) Transaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.DB(ctx).Transaction(fn)
}

func (r *BaseRepository[T]) WithTransaction(tx *gorm.DB) *BaseRepository[T] {
	return &BaseRepository[T]{
		db: tx,
	}
}

// Add to your repository methods
func (r *BaseRepository[T]) FindAllPaginated(ctx context.Context, model interface{}, page, pageSize int) (interface{}, int64, error) {
	var total int64
	r.db.WithContext(ctx).Model(model).Count(&total)

	offset := (page - 1) * pageSize

	err := r.db.WithContext(ctx).Offset(offset).Limit(pageSize).Find(model).Error
	return model, total, err
}
