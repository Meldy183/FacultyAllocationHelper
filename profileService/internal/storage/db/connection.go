package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/config"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"
	"go.uber.org/zap"
)

type ConnectAndInit struct {
	logger *zap.Logger
}

func NewConnectAndInit(logger *zap.Logger) *ConnectAndInit {
	return &ConnectAndInit{logger: logger}
}
func (str *ConnectAndInit) NewPostgresPool(ctx context.Context, cfg config.Database) (*pgxpool.Pool, error) {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DatabaseName,
		cfg.SSLMode)
	poolConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		str.logger.Error("Error parsing PostgreSQL connection string",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogNewPostgresPool),
			zap.Error(err),
		)
		return nil, err
	}
	str.logger.Info("Connected to PostgreSQL",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogNewPostgresPool),
	)
	poolConfig.MaxConns = int32(cfg.MaxOpenConnections)
	poolConfig.MinConns = int32(cfg.MaxIdleConnections)
	poolConfig.MaxConnLifetime = cfg.ConnMaxLifetime
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		str.logger.Error("Error connecting to PostgreSQL",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogNewPostgresPool),
			zap.Error(err),
		)
		return nil, err
	}
	str.logger.Info("config sent successfully, end connection func",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogNewPostgresPool),
	)
	return pool, err
}
func (str *ConnectAndInit) InitSchema(ctx context.Context, pool *pgxpool.Pool) error {
	var conn *pgxpool.Conn
	var err error
	for i := 1; i < 10; i++ {
		conn, err = pool.Acquire(ctx)
		if err != nil {
			str.logger.Warn("Error acquiring PostgreSQL connection",
				zap.String("layer", logctx.LogDBInitLayer),
				zap.String("function", logctx.LogNewPostgresPool),
				zap.Error(err),
				zap.Int("attempt", i),
			)
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	if err != nil {
		str.logger.Error("Error acquiring connection",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	defer conn.Release()

	// Enable uuid-ossp extension for UUID generation
	query := `CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating uuid-ossp extension",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created uuid-ossp extension",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS position (
      position_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      name VARCHAR(255) UNIQUE NOT NULL
    )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating position_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created position_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS user_profile (
      profile_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      email VARCHAR(50) UNIQUE NOT NULL,
      english_name VARCHAR(255) NOT NULL,
      russian_name VARCHAR(255),
      alias VARCHAR(255) UNIQUE NOT NULL,
      start_date DATE,
      end_date DATE
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating user_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err))
		return err
	}
	str.logger.Info("created user_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS language (
      code VARCHAR(20) PRIMARY KEY,
      language_name VARCHAR(255) UNIQUE NOT NULL
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating language_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err))
		return err
	}
	str.logger.Info("created language_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS institute (
      institute_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      name VARCHAR(255) UNIQUE NOT NULL
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating institute_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created institute_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS lab (
      lab_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      name VARCHAR(255) UNIQUE NOT NULL,
      institute_id UUID NOT NULL,
      FOREIGN KEY (institute_id) REFERENCES institute (institute_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating lab_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created lab_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS user_language (
      user_language_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      profile_id UUID NOT NULL,
      code VARCHAR(255) NOT NULL,
      FOREIGN KEY (profile_id) REFERENCES user_profile (profile_id),
      FOREIGN KEY (code) REFERENCES language (code)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating user_language_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created user_language_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS user_institute (
      user_institute_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      profile_id UUID NOT NULL,
      institute_id UUID NOT NULL,
      FOREIGN KEY (profile_id) REFERENCES user_profile (profile_id),
      FOREIGN KEY (institute_id) REFERENCES institute (institute_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating user_institute_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}

	query = `CREATE TABLE IF NOT EXISTS responsible_institute (
    responsible_institute_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    responsible_institute_name VARCHAR
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating responsible_institute table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created responsible_institute table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS course (
    course_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR (100),
	official_name VARCHAR (100),
	responsible_institute_id UUID,
    lec_hours INTEGER,
    lab_hours INTEGER,
	is_elective BOOL,
	FOREIGN KEY (responsible_institute_id) REFERENCES responsible_institute (responsible_institute_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating course table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created course table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS semester (
    semester_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    semester_name VARCHAR(20)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating semester table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created semester table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS academic_year (
    academic_year_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    academic_year_name VARCHAR(20)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating academic_year table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created academic_year table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS course_instance (
    instance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    course_id UUID NOT NULL,
    semester_id UUID,
    year INT,
    mode VARCHAR(20),
    academic_year_id UUID,
	hardness_coefficient FLOAT,
    form VARCHAR(30),
    groups_needed INT,
    groups_taken INT,
    pi_allocation_status VARCHAR(20),
    ti_allocation_status VARCHAR(20),
    FOREIGN KEY (course_id) REFERENCES course (course_id),
    FOREIGN KEY (semester_id) REFERENCES semester (semester_id),
    FOREIGN KEY (academic_year_id) REFERENCES academic_year (academic_year_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating course_instance table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created course_instance table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS program (
    program_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(20)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating program table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created program table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS track (
    track_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(20),
    program_id UUID,
    FOREIGN KEY (program_id) REFERENCES program (program_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating track table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created track table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS track_course_instance (
    track_course_instance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    track_id UUID,
    instance_id UUID,
    FOREIGN KEY (track_id) REFERENCES track (track_id),
    FOREIGN KEY (instance_id) REFERENCES course_instance (instance_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating track_course_instance table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created track_course_instance table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS program_course_instance (
    program_course_instance_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    program_id UUID,
    instance_id UUID,
    FOREIGN KEY (program_id) REFERENCES program (program_id),
    FOREIGN KEY (instance_id) REFERENCES course_instance (instance_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating program_course_instance table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created program_course_instance table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS user_profile_version (
    profile_version_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    profile_id UUID,
    year INT,
	maxload INT,
	position_id UUID,
	employment_type VARCHAR(128),
	student_type VARCHAR(32),
	fsro VARCHAR(128),
	frontal_hours INT,
	extra_activities FLOAT,
	degree BOOLEAN,
	mode VARCHAR(16),
    FOREIGN KEY (profile_id) REFERENCES user_profile (profile_id),
	FOREIGN KEY (position_id) REFERENCES position (position_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating user_profile_version table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err))
		return err
	}
	str.logger.Info("created user_profile_version table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema))

	query = `CREATE TABLE IF NOT EXISTS staff (
    assignment_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    instance_id UUID,
    profile_version_id UUID,
    position_type VARCHAR(3),
    groups_assigned INT,
    is_confirmed BOOLEAN,
    labs_count INT,
    tutorials_count INT,
    lectures_count INT,
    FOREIGN KEY (instance_id) REFERENCES course_instance (instance_id),
    FOREIGN KEY (profile_version_id) REFERENCES user_profile_version (profile_version_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating staff table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created staff table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS profile_course_instance (
		profile_course_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
		profile_version_id UUID NOT NULL,
		instance_id UUID NOT NULL,
		FOREIGN KEY (profile_version_id) REFERENCES user_profile_version (profile_version_id),
		FOREIGN KEY (instance_id) REFERENCES course_instance (instance_id)
	)`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating profile_course_instance",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created profile_course_instance_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS log (
  log_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  user_id VARCHAR(255),
  action VARCHAR(50),
  timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  subject_id UUID
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating log table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err))
		return err
	}
	str.logger.Info("created log table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema))

	query = `CREATE TABLE IF NOT EXISTS workload (
      workload_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      profile_version_id UUID NOT NULL,
      semester_id UUID NOT NULL,
      lectures_count INT NOT NULL,
      tutorials_count INT NOT NULL,
      labs_count INT NOT NULL,
      electives_count INT NOT NULL,
	  rate FLOAT, 
      FOREIGN KEY (profile_version_id) REFERENCES user_profile_version (profile_version_id),
	  FOREIGN KEY (semester_id) REFERENCES semester (semester_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating workload_table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created workload_table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	query = `CREATE TABLE IF NOT EXISTS institute_course_link (
      profile_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
      email VARCHAR(50) UNIQUE NOT NULL,
      english_name VARCHAR(255) NOT NULL,
      russian_name VARCHAR(255),
      alias VARCHAR(255) UNIQUE NOT NULL,
      start_date DATE,
      end_date DATE
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating institute_course_link table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err))
		return err
	}
	str.logger.Info("created institute_course_link table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	tx, err := pool.Begin(ctx)
	if err != nil {
		str.logger.Error("Error starting transaction",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err))
		return err
	}
	str.logger.Info("started transaction",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `
    INSERT INTO language (code, language_name)
    VALUES ('en', 'English'), ('ru', 'Russian')
    ON CONFLICT (code) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding language",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return fmt.Errorf("failed to insert languages: %w", err)
	}
	str.logger.Info("added languages SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	_, err = tx.Exec(ctx, `
    INSERT INTO institute (institute_id, name)
    VALUES ('00000000-0000-0000-0000-000000000001', 'Институт анализа данных и Искусственного Интеллекта'),
           ('00000000-0000-0000-0000-000000000002', 'Институт разработки ПО и программной инженерии'),
           ('00000000-0000-0000-0000-000000000003', 'Институт робототехники и компьютерного зрения'),
           ('00000000-0000-0000-0000-000000000004', 'Институт информационной безопасности'),
           ('00000000-0000-0000-0000-000000000005', 'Институт гуманитарных и социальных наук')
    ON CONFLICT (institute_id) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding institute manually",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
	}
	str.logger.Info("added institutes SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	_, err = tx.Exec(ctx, `
    INSERT INTO responsible_institute (responsible_institute_id, responsible_institute_name)
    VALUES ('00000000-0000-0000-0001-000000000001', 'DS'),
           ('00000000-0000-0000-0001-000000000002', 'DS/Math'),
           ('00000000-0000-0000-0001-000000000003', 'DS/SDE'),
           ('00000000-0000-0000-0001-000000000004', 'GAMEDEV'),
           ('00000000-0000-0000-0001-000000000005', 'HUM'),
		   ('00000000-0000-0000-0001-000000000006', 'RO'),
		   ('00000000-0000-0000-0001-000000000007', 'SDE'),
		   ('00000000-0000-0000-0001-000000000008', 'SNE')
    ON CONFLICT (responsible_institute_id) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding responsible_institute manually",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
	}
	str.logger.Info("added responsible_institutes SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	_, err = tx.Exec(ctx, `
    INSERT INTO program (program_id, name)
    VALUES ('00000000-0000-0000-0002-000000000001', 'AI360'),
           ('00000000-0000-0000-0002-000000000002', 'МОИИ'),
           ('00000000-0000-0000-0002-000000000003', 'BS RO'),
           ('00000000-0000-0000-0002-000000000004', 'AIDE'),
           ('00000000-0000-0000-0002-000000000005', 'SE'),
		   ('00000000-0000-0000-0002-000000000006', 'SNE'),
		   ('00000000-0000-0000-0002-000000000007', 'ROCV'),
		   ('00000000-0000-0000-0002-000000000008', 'MSRO'),
		   ('00000000-0000-0000-0002-000000000009', 'TE'),
           ('00000000-0000-0000-0002-000000000010', 'УРКИ'),
           ('00000000-0000-0000-0002-000000000011', 'КБ'),
           ('00000000-0000-0000-0002-000000000012', 'УнОД'),
           ('00000000-0000-0000-0002-000000000013', 'УЦП'),
		   ('00000000-0000-0000-0002-000000000014', 'DS'),
		   ('00000000-0000-0000-0002-000000000015', 'R'),
		   ('00000000-0000-0000-0002-000000000016', 'ITE'),
           ('00000000-0000-0000-0002-000000000017', 'ИиВТ'),
           ('00000000-0000-0000-0002-000000000018', 'DSAI'),
           ('00000000-0000-0000-0002-000000000019', 'CSE')
    ON CONFLICT (program_id) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding program manually",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
	}
	str.logger.Info("added program SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	_, err = tx.Exec(ctx, `
    INSERT INTO track (track_id, name)
    VALUES ('00000000-0000-0000-0003-000000000001', 'AAI'),
           ('00000000-0000-0000-0003-000000000002', 'AAIR'),
           ('00000000-0000-0000-0003-000000000003', 'CS'),
           ('00000000-0000-0000-0003-000000000004', 'CSDS'),
           ('00000000-0000-0000-0003-000000000005', 'DS'),
		   ('00000000-0000-0000-0003-000000000006', 'GD'),
		   ('00000000-0000-0000-0003-000000000007', 'ITE'),
		   ('00000000-0000-0000-0003-000000000008', 'R'),
		   ('00000000-0000-0000-0003-000000000009', 'SD'),
           ('00000000-0000-0000-0003-000000000010', 'SE'),
           ('00000000-0000-0000-0003-000000000011', 'SNE')
    ON CONFLICT (track_id) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding tracks manually",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
	}
	str.logger.Info("added tracks SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	_, err = tx.Exec(ctx, `
    INSERT INTO position (position_id, name)
    VALUES ('00000000-0000-0000-0004-000000000001', 'Professor'),
           ('00000000-0000-0000-0004-000000000002', 'Docent'),
           ('00000000-0000-0000-0004-000000000003', 'Senior Instructor'),
           ('00000000-0000-0000-0004-000000000004', 'Instructor'),
           ('00000000-0000-0000-0004-000000000005', 'TA'),
		   ('00000000-0000-0000-0004-000000000006', 'TA intern'),
		   ('00000000-0000-0000-0004-000000000007', 'Visiting')
    ON CONFLICT (position_id) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding positions manually",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
	}
	str.logger.Info("added positions SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	_, err = tx.Exec(ctx, `
    INSERT INTO semester (semester_id, semester_name)
    VALUES ('00000000-0000-0000-0005-000000000001', 'T1'),
           ('00000000-0000-0000-0005-000000000002', 'T2'),
           ('00000000-0000-0000-0005-000000000003', 'T3')
    ON CONFLICT (semester_id) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding semesters manually",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
	}
	str.logger.Info("added semesters SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	_, err = tx.Exec(ctx, `
    INSERT INTO academic_year (academic_year_id, academic_year_name)
    VALUES ('00000000-0000-0000-0006-000000000001', 'BS1'),
           ('00000000-0000-0000-0006-000000000002', 'BS2'),
           ('00000000-0000-0000-0006-000000000003', 'BS3'),
		   ('00000000-0000-0000-0006-000000000004', 'BS4'),
           ('00000000-0000-0000-0006-000000000005', 'MS1'),
           ('00000000-0000-0000-0006-000000000006', 'MS2'),
		   ('00000000-0000-0000-0006-000000000007', 'PhD1'),
           ('00000000-0000-0000-0006-000000000008', 'PhD2')
		ON CONFLICT (academic_year_id) DO NOTHING;
  `)
	if err != nil {
		str.logger.Error("Error adding academic_years manually",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
	}
	str.logger.Info("added academic_years SUCCESS",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	if err := tx.Commit(ctx); err != nil {
		str.logger.Error("Error committing transaction",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}
	str.logger.Info("committed transaction",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)
	return nil
}
