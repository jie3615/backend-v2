package comer

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/data/model/comer"
	"metaLand/data/model/comereducation"
	"metaLand/data/model/comerlanguage"
	"metaLand/data/model/comerprofile"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateComerInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户信息
func NewUpdateComerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateComerInfoLogic {
	return &UpdateComerInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateComerInfoLogic) UpdateComerInfo(req *types.UpdateComerInfoRequest) (resp *types.ComerMessageResponse, err error) {
	if req.ComerId == 0 {
		return nil, errors.New("comer_id is required")
	}

	comer, err := comer.FindComer(l.svcCtx.DB, uint64(req.ComerId))
	if err != nil {
		return nil, err
	}
	if comer == nil {
		return nil, errors.New("comer not found")
	}

	comerProfileInfo, err := comerprofile.FindComerProfile(l.svcCtx.DB, uint64(req.ComerId))
	if err != nil {
		return nil, err
	}

	var languageIds []string
	var educationIds []string

	// 处理语言信息
	if len(req.Languages) > 0 {
		// 先删除该用户的所有语言记录
		err = comerlanguage.DeleteComerLanguagesByComerId(l.svcCtx.DB, uint64(req.ComerId))
		if err != nil {
			logx.Errorf("delete comer languages failed: %v", err)
			return nil, err
		}

		// 添加新的语言记录
		for _, lang := range req.Languages {
			comerLanguage := &comerlanguage.ComerLanguage{
				ComerId:  uint64(req.ComerId),
				Language: lang.Language,
				Code:     lang.Code,
				Level:    lang.Level,
				IsNative: lang.IsNative,
			}
			_, err = comerlanguage.CreateComerLanguage(l.svcCtx.DB, comerLanguage)
			if err != nil {
				logx.Errorf("insert comer language failed: %v", err)
				return nil, err
			}
			languageIds = append(languageIds, strconv.FormatUint(comerLanguage.ID, 10))
		}
	}

	// 处理教育信息
	if len(req.Educations) > 0 {
		// 先删除该用户的所有教育记录
		err = comereducation.DeleteComerEducationsByComerId(l.svcCtx.DB, uint64(req.ComerId))
		if err != nil {
			logx.Errorf("delete comer educations failed: %v", err)
			return nil, err
		}

		// 添加新的教育记录
		for _, edu := range req.Educations {
			comerEducation := &comereducation.ComerEducation{
				ComerId:     uint64(req.ComerId),
				School:      edu.School,
				Degree:      edu.Degree,
				Major:       edu.Major,
				StartDate:   edu.StartDate,
				EndDate:     edu.EndDate,
				Description: edu.Description,
			}
			_, err = comereducation.CreateComerEducation(l.svcCtx.DB, comerEducation)
			if err != nil {
				logx.Errorf("insert comer education failed: %v", err)
				return nil, err
			}
			educationIds = append(educationIds, strconv.FormatUint(comerEducation.ID, 10))
		}
	}

	// 更新 comerProfileInfo
	if comerProfileInfo == nil {
		comerProfileInfo = &comerprofile.ComerProfile{
			ComerId:    int64(req.ComerId),
			Name:       req.Name,
			Avatar:     req.Avatar,
			Cover:      req.Cover,
			Location:   req.Location,
			TimeZone:   req.TimeZone,
			Website:    req.Website,
			Email:      req.Email,
			Twitter:    req.Twitter,
			Discord:    req.Discord,
			Telegram:   req.Telegram,
			Medium:     req.Medium,
			Facebook:   req.Facebook,
			Linktree:   req.Linktree,
			Bio:        req.Bio,
			Languages:  strings.Join(languageIds, ","),
			Educations: strings.Join(educationIds, ","),
		}
		err = comerprofile.InsertComerProfile(l.svcCtx.DB, comerProfileInfo)
	} else {
		// 更新现有 profile
		updates := map[string]interface{}{
			"name":       req.Name,
			"avatar":     req.Avatar,
			"cover":      req.Cover,
			"location":   req.Location,
			"time_zone":  req.TimeZone,
			"website":    req.Website,
			"email":      req.Email,
			"twitter":    req.Twitter,
			"discord":    req.Discord,
			"telegram":   req.Telegram,
			"medium":     req.Medium,
			"facebook":   req.Facebook,
			"linktree":   req.Linktree,
			"bio":        req.Bio,
			"languages":  strings.Join(languageIds, ","),
			"educations": strings.Join(educationIds, ","),
			"updated_at": time.Now(),
		}
		err = comerprofile.UpdateComerProfile(l.svcCtx.DB, uint64(req.ComerId), updates)
	}

	if err != nil {
		logx.Errorf("update comer profile failed: %v", err)
		return nil, err
	}

	return &types.ComerMessageResponse{
		Message: "success",
	}, nil
}
