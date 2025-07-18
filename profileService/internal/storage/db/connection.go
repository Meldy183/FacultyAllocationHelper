package db

import (
	"context"
	"fmt"

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
	conn, err := pool.Acquire(ctx)
	if err != nil {
		str.logger.Error("Error acquiring connection",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	defer conn.Release()
	query := `CREATE TABLE IF NOT EXISTS position (
      position_id SERIAL PRIMARY KEY,
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
      profile_id SERIAL PRIMARY KEY,
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
      institute_id SERIAL PRIMARY KEY,
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
      lab_id SERIAL PRIMARY KEY,
      name VARCHAR(255) UNIQUE NOT NULL,
      institute_id INT NOT NULL,
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
      user_language_id SERIAL PRIMARY KEY,
      profile_id INT NOT NULL,
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
      user_institute_id SERIAL PRIMARY KEY,
      profile_id INT NOT NULL,
      institute_id INT NOT NULL,
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
	query = `CREATE TABLE IF NOT EXISTS course (
    course_id SERIAL PRIMARY KEY,
    code VARCHAR (50),
    name VARCHAR (50),
	is_elective BOOL,
    officialName VARCHAR (100),
    institute_id INTEGER,
    lec_hours INTEGER,
    lab_hours INTEGER
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

	query = `CREATE TABLE IF NOT EXISTS responsible_institute (
    responsible_institute_id SERIAL PRIMARY KEY,
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
	query = `CREATE TABLE IF NOT EXISTS institute_course_link (
    institute_course_id SERIAL PRIMARY KEY,
    course_id INT,
    responsible_institute_id INT,
    FOREIGN KEY (course_id) REFERENCES course (course_id),
    FOREIGN KEY (responsible_institute_id) REFERENCES responsible_institute (responsible_institute_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {
		str.logger.Error("Error creating institute_course_link table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created institute_course_link table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)
	query = `CREATE TABLE IF NOT EXISTS semester (
    semester_id SERIAL PRIMARY KEY,
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
    academic_year_id SERIAL PRIMARY KEY,
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
    instance_id SERIAL PRIMARY KEY,
    course_id INT,
    semester VARCHAR(2),
    year INT,
    mode VARCHAR(20),
    academic_year_id INT,
	hardness_coefficient DECIMAL,
    form VARCHAR(30),
    groups_needed INT,
    groups_taken INT,
    pi_allocation_status VARCHAR(20),
    ti_allocation_status VARCHAR(20),
    FOREIGN KEY (course_id) REFERENCES course (course_id)
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
    program_id SERIAL PRIMARY KEY,
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
    track_id SERIAL PRIMARY KEY,
    name VARCHAR(20),
    program_id INT,
    FOREIGN KEY (program_id) REFERENCES program (program_id)
  )` // name это типо CS AI пон???
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
    track_course_instance_id SERIAL PRIMARY KEY,
    track_id INT,
    instance_id INT,
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
    program_course_instance_id SERIAL PRIMARY KEY,
    program_id INT,
    instance_id INT,
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
	query = `CREATE TABLE IF NOT EXISTS course_staff (
    assignment_id SERIAL PRIMARY KEY,
    instance_id INT,
    profile_id INT,
    position_type VARCHAR(3),
    contribution_coefficient DECIMAL,
    groups_assigned INT,
    is_confirmed BOOLEAN,
    labs_count INT,
    tutorials_count INT,
    lectures_count INT,
    FOREIGN KEY (instance_id) REFERENCES course_instance (instance_id)
  )`
	_, err = conn.Exec(ctx, query)
	if err != nil {

		str.logger.Error("Error creating course_staff table",
			zap.String("layer", logctx.LogDBInitLayer),
			zap.String("function", logctx.LogInitSchema),
			zap.Error(err),
		)
		return err
	}
	str.logger.Info("created course_staff table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema),
	)

	str.logger.Info("created profile_staff table",
		zap.String("layer", logctx.LogDBInitLayer),
		zap.String("function", logctx.LogInitSchema))
	query = `CREATE TABLE IF NOT EXISTS user_profile_version (
    profile_version_id SERIAL PRIMARY KEY,
    profile_id INT,
    year INT,
	maxload INT,
	position_id INT,
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
	query = `CREATE TABLE IF NOT EXISTS profile_course_instance (
    profile_course_id SERIAL PRIMARY KEY,
    profile_version_id INT NOT NULL,
    instance_id INT NOT NULL,
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
  log_id SERIAL PRIMARY KEY,
  user_id VARCHAR(255),
  action VARCHAR(50),
  timestamp TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  subject_id INT
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
      workload_id SERIAL PRIMARY KEY,
      profile_version_id INT NOT NULL,
      semester_id INT NOT NULL,
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
    VALUES (1, 'Институт анализа данных и Искусственного Интеллекта'),
           (2, 'Институт разработки ПО и программной инженерии'),
           (3, 'Институт робототехники и компьютерного зрения'),
           (4, 'Институт информационной безопасности'),
           (5, 'Институт гуманитарных и социальных наук')
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
    VALUES (1, 'DS'),
           (2, 'DS/Math'),
           (3, 'DS/SDE'),
           (4, 'GAMEDEV'),
           (5, 'HUM'),
		   (6, 'RO'),
		   (7, 'SDE'),
		   (8, 'SNE')
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
    INSERT INTO position (position_id, name)
    VALUES (1, 'Professor'),
           (2, 'Docent'),
           (3, 'Senior Instructor'),
           (4, 'Instructor'),
           (5, 'TA'),
		   (6, 'TA intern'),
		   (7, 'Visiting')
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
    VALUES (1, 'T1'),
           (2, 'T2'),
           (3, 'T3')
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
    VALUES (1, 'BS1'),
           (2, 'BS2'),
           (3, 'BS3'),
		   (4, 'BS4'),
           (5, 'MS1'),
           (6, 'MS2'),
		   (7, 'PhD1'),
           (8, 'PhD2')
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
