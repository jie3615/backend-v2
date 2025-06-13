package bounty

import (
	"context"
	"metaLand/data/model/bounty"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailBountyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询bounty详情
func NewDetailBountyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailBountyLogic {
	return &DetailBountyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailBountyLogic) DetailBounty(req *types.DetailRequest) (resp *types.BountyInfoResponse, err error) {
	var response types.BountyInfoResponse
	b := bounty.Bounty{}
	err = bounty.GetBounty(l.svcCtx.DB, req.Id, &b)
	if err != nil {
		return nil, err
	}
	response = types.BountyInfoResponse{
		ApplicantDeposit: b.ApplicantDeposit,
		//ApplicantMinDeposit: b.ApplicantMinDeposit,
		//ApplyDeadline:
		// BountyApplicant:  types.BountyApplicant{
		//	ApplyAt:     b.BountyApplicant.ApplyAt,
		//	ApproveAt:   b.BountyApplicant.ApproveAt,
		//	BountyID:    b.BountyApplicant.BountyID,
		//	ComerID:     b.BountyApplicant.ComerID,
		//	Description: b.BountyApplicant.Description,
		//},
		ChainId: b.ChainId,
		ComerId: b.ComerId,
		//ContractAddress:
		CreatedAt:              b.CreatedAt.Format("2006-01-02 15:04:05"),
		DepositContractAddress: b.DepositContract,
		//DepositContractTokenDecimal:
		DepositContractTokenSymbol: b.DepositTokenSymbol,
		Description:                b.Description,
		DiscussionLink:             b.DiscussionLink,
		ExpiredTime:                b.ApplyCutoffDate.Format("2006-01-02 15:04:05"),
		//  Founder: types.BountyComer{
		//	ComerID:     b.Founder.ComerID,
		//	Description: b.Founder.Description,
		//	Email:       b.Founder.Email,
		//	Name:        b.Founder.Name,
		//	Phone:       b.Founder.Phone,
		//	TelegramID:  b.Founder.TelegramID,
		//},
		FounderDeposit: b.FounderDeposit,
		Id:             b.ID,
		//IsLock:         b.IsDeleted,
		Title:  b.Title,
		TxHash: b.TxHash,
		Period: types.BountyPaymentPeriod{
			BountyId:     b.BountyPaymentPeriod.BountyId,
			PeriodType:   b.BountyPaymentPeriod.PeriodType,
			PeriodAmount: b.BountyPaymentPeriod.PeriodAmount,
			HoursPerDay:  b.BountyPaymentPeriod.HoursPerDay,
			Token1Symbol: b.BountyPaymentPeriod.Token1Symbol,
			Token1Amount: b.BountyPaymentPeriod.Token1Amount,
			Token2Symbol: b.BountyPaymentPeriod.Token2Symbol,
			Token2Amount: b.BountyPaymentPeriod.Token2Amount,
			Target:       b.BountyPaymentPeriod.Target,
		},
	}
	for _, applicant := range b.BountyApplicants {
		response.Applicants = append(response.Applicants, types.BountyApplicant{
			ComerID:     applicant.ComerId,
			ApplyAt:     applicant.ApplyAt.Format("2006-01-02 15:04:05"),
			RevokeAt:    applicant.RevokeAt.Format("2006-01-02 15:04:05"),
			ApproveAt:   applicant.ApproveAt.Format("2006-01-02 15:04:05"),
			QuitAt:      applicant.QuitAt.Format("2006-01-02 15:04:05"),
			SubmitAt:    applicant.SubmitAt.Format("2006-01-02 15:04:05"),
			Status:      applicant.Status,
			Description: applicant.Description,
			BountyId:    applicant.BountyId,
		})
	}
	for _, contact := range b.BountyContacts {
		response.Contacts = append(response.Contacts, types.BountyContact{
			BountyId:       contact.BountyId,
			ContactAddress: contact.ContactAddress,
			ContactType:    contact.ContactType,
		})
	}
	for _, deposit := range b.BountyDeposits {
		response.DepositRecords = append(response.DepositRecords, types.BountyDepositRecord{
			Amount:    deposit.TokenAmount,
			BountyId:  deposit.BountyId,
			ComerId:   deposit.ComerId,
			CreatedAt: deposit.CreatedAt.Format("2006-01-02 15:04:05"),
			Mode:      deposit.Access,
			Status:    deposit.Status,
			TxHash:    deposit.TxHash,
		})
	}

	for _, terms := range b.BountyPaymentTerms {
		response.Terms = append(response.Terms, types.BountyPaymentTerms{
			BountyId:     terms.BountyId,
			PaymentMode:  terms.PaymentMode,
			SeqNum:       terms.SeqNum,
			Status:       terms.Status,
			Terms:        terms.Terms,
			Token1Amount: terms.Token1Amount,
			Token1Symbol: terms.Token1Symbol,
			Token2Amount: terms.Token2Amount,
			Token2Symbol: terms.Token2Symbol,
		})
	}
	return &response, nil
}
