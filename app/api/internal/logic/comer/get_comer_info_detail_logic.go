package comer

import (
	"context"
	"errors"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comeraccount"
	"metaLand/data/model/comereducation"
	"metaLand/data/model/comerlanguage"
	"metaLand/data/model/comerprofile"
	"metaLand/data/model/comerskill"
	"metaLand/data/model/comersocial"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetComerInfoDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户详情
func NewGetComerInfoDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetComerInfoDetailLogic {
	return &GetComerInfoDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetComerInfoDetailLogic) GetComerInfoDetail() (resp *types.ComerInfoDetailResponse, err error) {
	comerInfo, ok := l.ctx.Value("comerInfo").(*comer.Comer)
	if !ok {
		return nil, errors.New("user not found")
	}

	// 获取用户基本信息
	comerProfile, err := comerprofile.FindComerProfile(l.svcCtx.DB, uint64(comerInfo.ID))
	if err != nil {
		return nil, err
	}

	// 获取用户账户信息
	accounts, err := l.GetAccountsResponse(int(comerInfo.ID))
	if err != nil {
		return nil, err
	}

	// 获取用户技能
	skillResponses, err := l.GetSkillsResponse(int(comerInfo.ID))
	if err != nil {
		return nil, err
	}

	// 获取用户连接总数
	connectedTotalResponse, err := l.GetConnectedTotalResponse(int(comerInfo.ID))
	if err != nil {
		return nil, err
	}

	// 获取用户教育经历
	educations, err := l.GetEducationsResponse(int(comerInfo.ID))
	if err != nil {
		return nil, err
	}

	// 获取用户语言
	// todo 待确认返回格式或者取另外的表数据
	languages, err := l.GetLanguagesResponse(int(comerInfo.ID))
	if err != nil {
		return nil, err
	}

	// 获取用户社会关系
	socialRelations, err := l.GetSocialRelationsResponse(int(comerInfo.ID))
	if err != nil {
		return nil, err
	}

	// 构建完整响应
	resp = &types.ComerInfoDetailResponse{
		Accounts:       accounts,
		Activation:     true,
		Address:        comerInfo.Address,
		Avatar:         map[bool]string{true: comerProfile.Avatar, false: ""}[comerProfile != nil],
		Banner:         map[bool]string{true: comerProfile.Cover, false: ""}[comerProfile != nil],
		ConnectedTotal: connectedTotalResponse,
		CustomDomain:   map[bool]string{true: comerProfile.Website, false: ""}[comerProfile != nil],
		Educations:     educations,
		Id:             int(comerInfo.ID),
		Info: types.ComerInfo{
			ComerId: int(comerInfo.ID),
			Id:      int(comerInfo.ID),
			Bio:     map[bool]string{true: comerProfile.Bio, false: ""}[comerProfile != nil],
		},
		InvitationCode: "",
		IsConnected:    false,
		Languages:      languages,
		Location:       map[bool]string{true: comerProfile.Location, false: ""}[comerProfile != nil],
		Name:           map[bool]string{true: comerProfile.Name, false: ""}[comerProfile != nil],
		Skills:         skillResponses,
		Socials:        socialRelations,
		TimeZone:       map[bool]string{true: comerProfile.TimeZone, false: ""}[comerProfile != nil],
	}

	return resp, nil
}

// 获取用户账户信息
func (l *GetComerInfoDetailLogic) GetAccountsResponse(comerId int) ([]types.ComerAccountResponse, error) {
	accounts, err := comeraccount.ListComerAccounts(l.svcCtx.DB, uint64(comerId))
	if err != nil {
		return nil, err
	}

	accountResponses := make([]types.ComerAccountResponse, 0, len(accounts))
	for _, account := range accounts {
		accountResponses = append(accountResponses, types.ComerAccountResponse{
			Id:        int(account.ID),
			IsLinked:  account.IsLinked,
			IsPrimary: account.IsPrimary,
			Nickname:  account.Nick,
			Oin:       account.Oin,
			Type:      int(account.Type),
		})
	}
	return accountResponses, nil
}

// // 获取skills
// func (l *GetComerInfoDetailLogic) getSkillsResponse() ([]types.ComerSkillResponse, error) {
// 	tagList := make([]tag.Tag, 0)
// 	err := tag.ListTags(l.svcCtx.DB, &tagList)
// 	if err != nil {
// 		return nil, err
// 	}

// 	tagsMap := make(map[uint64]tag.Tag)
// 	for _, tag := range tagList {
// 		tagsMap[uint64(tag.ID)] = tag
// 	}

// 	tagTargetRelList := make([]tag.TagTargetRel, 0)
// 	err = tag.ListTagRelations(l.svcCtx.DB, &tagTargetRelList)
// 	if err != nil {
// 		return nil, err
// 	}

// 	skillResponses := make([]types.TagRelationResponse, 0, len(tagTargetRelList))
// 	for _, tagTargetRel := range tagTargetRelList {
// 		tag := tagsMap[tagTargetRel.TagID]
// 		skillResponses = append(skillResponses, types.TagRelationResponse{
// 			Id: int(tagTargetRel.TagID),
// 			Tag: types.TagResponse{
// 				Id:       int(tag.ID),
// 				Name:     tag.Name,
// 				Category: string(tag.Category),
// 			},
// 			TagId:    int(tag.ID),
// 			TargetId: int(tagTargetRel.TargetID),
// 			Type:     CategoryToType(string(tag.Category)),
// 		})
// 	}

// 	return skillResponses, nil
// }

// 获取用户技能
func (l *GetComerInfoDetailLogic) GetSkillsResponse(comerId int) ([]types.ComerSkillResponse, error) {
	comerSkills, err := comerskill.ListComerSkills(l.svcCtx.DB, uint64(comerId))
	if err != nil {
		return nil, err
	}
	skillResponses := make([]types.ComerSkillResponse, 0, len(comerSkills))
	for _, skill := range comerSkills {
		skillResponses = append(skillResponses, types.ComerSkillResponse{
			Id:          int(skill.ID),
			SkillName:   skill.SkillName,
			Level:       skill.Level,
			Years:       skill.Years,
			Description: skill.Description,
		})
	}
	return skillResponses, nil
}

// 获取用户连接总数
func (l *GetComerInfoDetailLogic) GetConnectedTotalResponse(comerId int) (types.ComerConnectedTotalResponse, error) {
	connectedTotalResponse := types.ComerConnectedTotalResponse{
		BeConnectComerTotal: 0,
		ConnectComerTotal:   0,
		ConnectStartupTotal: 0,
	}

	return connectedTotalResponse, nil
}

// 获取用户教育经历
func (l *GetComerInfoDetailLogic) GetEducationsResponse(comerId int) ([]types.ComerEducationResponse, error) {
	comerEducations, err := comereducation.ListComerEducations(l.svcCtx.DB, uint64(comerId))
	if err != nil {
		return nil, err
	}

	educationResponses := make([]types.ComerEducationResponse, 0, len(comerEducations))
	for _, education := range comerEducations {
		educationResponses = append(educationResponses, types.ComerEducationResponse{
			Id:          int(education.ID),
			School:      education.School,
			Degree:      education.Degree,
			StartDate:   education.StartDate,
			EndDate:     education.EndDate,
			Description: education.Description,
			Level:       education.Level,
			Major:       education.Major,
		})
	}
	return educationResponses, nil
}

// 获取用户语言
func (l *GetComerInfoDetailLogic) GetLanguagesResponse(comerId int) ([]types.ComerLanguageResponse, error) {
	comerLanguages, err := comerlanguage.ListComerLanguages(l.svcCtx.DB, uint64(comerId))
	if err != nil {
		return nil, err
	}
	languageResponses := make([]types.ComerLanguageResponse, 0, len(comerLanguages))
	for _, language := range comerLanguages {
		languageResponses = append(languageResponses, types.ComerLanguageResponse{
			Id:       int(language.ID),
			Language: language.Language,
			Code:     language.Code,
			Level:    language.Level,
			IsNative: language.IsNative,
		})
	}
	return languageResponses, nil
}

// 获取用户社会关系
func (l *GetComerInfoDetailLogic) GetSocialRelationsResponse(comerId int) ([]types.ComerSocialResponse, error) {
	comerSocials, err := comersocial.ListComerSocials(l.svcCtx.DB, uint64(comerId))
	if err != nil {
		return nil, err
	}
	socialResponses := make([]types.ComerSocialResponse, 0, len(comerSocials))
	for _, social := range comerSocials {
		socialResponses = append(socialResponses, types.ComerSocialResponse{
			Id:           int(social.ID),
			PlatformName: social.Platform,
			UserName:     social.Username,
			PlatformId:   "",
			Url:          social.Url,
			IsVerified:   social.IsVerified,
		})
	}
	return socialResponses, nil
}

// CategoryToType 将 Category 转换为对应的 type 值
func CategoryToType(category string) int {
	switch category {
	case "comerSkill":
		return 1
	case "startup":
		return 2
	case "bounty":
		return 3
	default:
		return 0
	}
}
