//go:build wireinject

package app

//go:generate wire
import (
	"github.com/google/wire"
	"github.com/jackc/pgx/v5/pgxpool"
	domainCompleteCourse "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteCourse"
	domainCompleteUser "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteUser"
	domainAcademicYear "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/academicYear"
	domainAllocation "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/allocation"
	domainCourse "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/course"
	domainCourseInstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/courseInstance"
	domainFacultyProfile "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	domainInstitute "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	domainLanguage "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	domainParsing "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/parsing"
	domainPosition "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/position"
	domainProfileCourseInstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileCourseInstance"
	domainProfileInstitute "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileInstitute"
	domainProfileLanguage "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileLanguage"
	domainProfileVersion "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	domainProgram "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/program"
	domainProgramCourseInstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/programCourseInstance"
	domainResponsibleInstitute "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/responsibleInstitute"
	domainSemester "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/semester"
	domainStaff "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/staff"
	domainTrack "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/track"
	domainTrackCourseInstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/trackCourseInstance"
	domainTransaction "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/transaction"
	domainWorkload "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/workload"
	allocationHandler "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/allocation"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/courses"
	userprofile2 "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/filters"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/handler/parse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/academicYear"
	allocationService "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/allocation"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/completeCourse"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/completeUser"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/course"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/courseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/language"
	Parsing "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/parsing"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/position"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileLanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/program"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/programCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/responsibleInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/semester"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/staff"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/track"
	trackcourseinstance "gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/trackCourseInstance"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/service/workload"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/storage/postgres"
	"go.uber.org/zap"
)

func InitializeApp(
	pool *pgxpool.Pool,
	logger *zap.Logger,
) (*App, error) {
	wire.Build(
		// Unit of Work
		postgres.NewPostgresUnitOfWork,
		wire.Bind(new(domainTransaction.UnitOfWork), new(*postgres.PostgresUnitOfWork)),
		// Repositories
		postgres.NewFacultyProfileRepo,
		wire.Bind(new(domainFacultyProfile.Repository), new(*postgres.FacultyProfileRepo)),

		postgres.NewUserLanguageRepo,
		wire.Bind(new(domainProfileLanguage.Repository), new(*postgres.UserLanguageRepo)),

		postgres.NewUserInstituteRepo,
		wire.Bind(new(domainProfileInstitute.Repository), new(*postgres.UserInstituteRepo)),

		postgres.NewUserCourseInstance,
		wire.Bind(new(domainProfileCourseInstance.Repository), new(*postgres.ProfileCourseInstanceRepo)),

		postgres.NewPositionRepo,
		wire.Bind(new(domainPosition.Repository), new(*postgres.PositionRepo)),

		postgres.NewLanguageRepo,
		wire.Bind(new(domainLanguage.Repository), new(*postgres.LanguageRepo)),

		postgres.NewInstituteRepo,
		wire.Bind(new(domainInstitute.Repository), new(*postgres.InstituteRepo)),

		postgres.NewUserProfileVersionRepo,
		wire.Bind(new(domainProfileVersion.Repository), new(*postgres.ProfileVersionRepo)),

		postgres.NewSemesterWorkloadRepo,
		wire.Bind(new(domainWorkload.Repository), new(*postgres.WorkloadRepo)),

		postgres.NewProgramRepo,
		wire.Bind(new(domainProgram.Repository), new(*postgres.ProgramRepo)),

		postgres.NewStaffRepo,
		wire.Bind(new(domainStaff.Repository), new(*postgres.StaffRepo)),

		postgres.NewTrackRepo,
		wire.Bind(new(domainTrack.Repository), new(*postgres.TrackRepo)),

		postgres.NewCourseRepo,
		wire.Bind(new(domainCourse.Repository), new(*postgres.CourseRepo)),

		postgres.NewSemesterRepo,
		wire.Bind(new(domainSemester.Repository), new(*postgres.SemesterRepo)),

		postgres.NewAcademicYearRepo,
		wire.Bind(new(domainAcademicYear.Repository), new(*postgres.AcademicYearRepo)),

		postgres.NewResponsibleInstituteRepo,
		wire.Bind(new(domainResponsibleInstitute.Repository), new(*postgres.ResponsibleInstituteRepo)),

		postgres.NewCourseInstanceRepo,
		wire.Bind(new(domainCourseInstance.Repository), new(*postgres.CourseInstanceRepo)),

		postgres.NewTrackCourseRepo,
		wire.Bind(new(domainTrackCourseInstance.Repository), new(*postgres.TrackCourseRepo)),

		postgres.NewProgramCourseRepo,
		wire.Bind(new(domainProgramCourseInstance.Repository), new(*postgres.ProgramCourseRepo)),

		// Services
		facultyProfile.NewService,
		wire.Bind(new(domainFacultyProfile.Service), new(*facultyProfile.Service)),

		profileLanguage.NewService,
		wire.Bind(new(domainProfileLanguage.Service), new(*profileLanguage.Service)),

		profileCourseInstance.NewService,
		wire.Bind(new(domainProfileCourseInstance.Service), new(*profileCourseInstance.Service)),

		position.NewService,
		wire.Bind(new(domainPosition.Service), new(*position.Service)),

		profileInstitute.NewService,
		wire.Bind(new(domainProfileInstitute.Service), new(*profileInstitute.Service)),

		language.NewService,
		wire.Bind(new(domainLanguage.Service), new(*language.Service)),

		institute.NewService,
		wire.Bind(new(domainInstitute.Service), new(*institute.Service)),

		responsibleInstitute.NewService,
		wire.Bind(new(domainResponsibleInstitute.Service), new(*responsibleInstitute.Service)),

		profileVersion.NewService,
		wire.Bind(new(domainProfileVersion.Service), new(*profileVersion.Service)),

		workload.NewService,
		wire.Bind(new(domainWorkload.Service), new(*workload.Service)),

		program.NewService,
		wire.Bind(new(domainProgram.Service), new(*program.Service)),

		track.NewService,
		wire.Bind(new(domainTrack.Service), new(*track.Service)),

		trackcourseinstance.NewService,
		wire.Bind(new(domainTrackCourseInstance.Service), new(*trackcourseinstance.Service)),

		programCourseInstance.NewService,
		wire.Bind(new(domainProgramCourseInstance.Service), new(*programCourseInstance.Service)),

		courseInstance.NewService,
		wire.Bind(new(domainCourseInstance.Service), new(*courseInstance.Service)),

		course.NewService,
		wire.Bind(new(domainCourse.Service), new(*course.Service)),

		completeCourse.NewService,
		wire.Bind(new(domainCompleteCourse.Service), new(*completeCourse.Service)),

		completeUser.NewService,
		wire.Bind(new(domainCompleteUser.Service), new(*completeUser.Service)),

		Parsing.NewService,
		wire.Bind(new(domainParsing.Service), new(*Parsing.Service)),

		staff.NewStaffService,
		wire.Bind(new(domainStaff.Service), new(*staff.Service)),

		academicYear.NewService,
		wire.Bind(new(domainAcademicYear.Service), new(*academicYear.Service)),

		semester.NewService,
		wire.Bind(new(domainSemester.Service), new(*semester.Service)),

		allocationService.NewService,
		wire.Bind(new(domainAllocation.Service), new(*allocationService.Service)),
		// Handlers
		userprofile2.NewHandler,
		courses.NewHandler,
		filters.NewHandler,
		parse.NewHandler,
		allocationHandler.NewHandler,

		// App
		NewApp,
	)
	return nil, nil
}
