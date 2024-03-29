package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"app/config"
	"app/storage"
)

type store struct {
	db     *pgxpool.Pool
	user   *userRepo
	role   *roleRepo
	course *courseRepo
	group  *groupRepo
	lesson *lessonRepo
	device *deviceRepo
}

func NewConnectionPostgres(cfg *config.Config) (storage.StorageI, error) {

	connect, err := pgxpool.ParseConfig(fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresUser,
		cfg.PostgresDatabase,
		cfg.PostgresPassword,
		cfg.PostgresPort,
	))

	if err != nil {
		return nil, err
	}
	connect.MaxConns = cfg.PostgresMaxConnection

	pgxpool, err := pgxpool.ConnectConfig(context.Background(), connect)
	if err != nil {
		return nil, err
	}

	return &store{
		db: pgxpool,
	}, nil
}

func (s *store) Close() {
	s.db.Close()
}
func (s *store) User() storage.UserRepoI {
	if s.user == nil {
		s.user = NewUserRepo(s.db)
	}

	return s.user
}

func (s *store) Role() storage.RoleRepoI {
	if s.role == nil {
		s.role = NewRoleRepo(s.db)
	}

	return s.role
}

func (s *store) Course() storage.CourseRepoI {
	if s.course == nil {
		s.course = NewCourseRepo(s.db)
	}

	return s.course
}

func (s *store) Group() storage.GroupRepoI {
	if s.group == nil {
		s.group = NewGroupRepo(s.db)
	}

	return s.group
}
func (s *store) Lesson() storage.LessonRepoI {
	if s.lesson == nil {
		s.lesson = NewLessonRepo(s.db)
	}

	return s.lesson
}

func (s *store) Device() storage.DeviceRepoI {
	if s.device == nil {
		s.device = NewDeviceRepo(s.db)
	}

	return s.device
}
