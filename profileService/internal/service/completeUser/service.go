package completeUser

import (
	"context"

	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/CompleteUser"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/facultyProfile"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/institute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/language"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileInstitute"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileLanguage"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/domain/profileVersion"
	"gitlab.pg.innopolis.university/f.markin/fah/profileService/internal/logctx"

	"go.uber.org/zap"
)

var _ CompleteUser.Service = (*Service)(nil)

type Service struct {
	logger                 *zap.Logger
	profileService         facultyProfile.Service
	profileVersionService  profileVersion.Service
	languageService        language.Service
	profileLanguageService profileLanguage.Service
	instituteService       institute.Service
	userInstituteService   profileInstitute.Service
}

func NewService(logger *zap.Logger,
	profileService facultyProfile.Service,
	profileVersionService profileVersion.Service,
	languageService language.Service,
	profileLanguageService profileLanguage.Service,
	instituteService institute.Service,
	userInstituteService profileInstitute.Service) *Service {
	return &Service{
		logger:                 logger,
		profileService:         profileService,
		profileVersionService:  profileVersionService,
		languageService:        languageService,
		profileLanguageService: profileLanguageService,
		instituteService:       instituteService,
		userInstituteService:   userInstituteService,
	}

}
func (s *Service) AddFullUser(ctx context.Context, fulluser *CompleteUser.FullUser) error {
	err := s.profileService.AddProfile(ctx, &fulluser.UserProfile)
	if err != nil {
		s.logger.Error("Error adding user profile",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddFullUser),
			zap.Error(err),
		)
		return err
	}
	s.logger.Info("Added user successfully")
	err = s.profileVersionService.AddProfileVersion(ctx, &fulluser.UserProfileVersion)
	if err != nil {
		s.logger.Error("Error adding user profile version",
			zap.String("layer", logctx.LogServiceLayer),
			zap.String("function", logctx.LogAddFullUser),
			zap.Error(err),
		)
		return err
	}
	s.logger.Info("Added course instance successfully")
	for _, institute := range fulluser.Institutes {
		instituteID, err := s.instituteService.GetInstituteIDByName(ctx, *institute)
		if err != nil {
			s.logger.Error("error seeking institute by name",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddFullUser),
				zap.Error(err))
			continue
		}
		err = s.userInstituteService.AddUserInstitute(ctx, &profileInstitute.UserInstitute{
			InstituteID: *instituteID,
			ProfileID:   fulluser.UserProfile.ProfileID})
		if err != nil {
			s.logger.Error("error adding institute to user profile",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddFullUser),
				zap.Error(err))
		}
	}
	for _, language := range fulluser.Languages {
		languageCode, err := s.languageService.GetCodeByLanguageName(ctx, *language)
		if err != nil {
			s.logger.Error("error seeking language code by name",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddFullUser),
				zap.Error(err))
			continue
		}
		err = s.profileLanguageService.AddUserLanguage(ctx, &profileLanguage.ProfileLanguage{LanguageCode: *languageCode,
			ProfileID: fulluser.UserProfile.ProfileID})
		if err != nil {
			s.logger.Error("error adding lanuage to user profile",
				zap.String("layer", logctx.LogServiceLayer),
				zap.String("function", logctx.LogAddFullUser),
				zap.Error(err))
		}
	}
	return nil
}
