package authorizations

import (
	"context"
	"errors"
	"fmt"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"
	"metaLand/app/api/utility/jwt"
	"metaLand/data/model/comer"
	"metaLand/data/model/comerprofile"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginByWalletAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 钱包地址登录
func NewLoginByWalletAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginByWalletAddressLogic {
	return &LoginByWalletAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginByWalletAddressLogic) LoginByWalletAddress(req *types.EthLoginRequest) (resp *types.ComerLoginResponse, err error) {
	// todo: add your logic here and delete this line
	// 1. 获取 nonce
	nonce, err := l.svcCtx.RedisClient.Get(context.TODO(), req.Address).Result()
	if err != nil {
		// redis: nil 表示 key 不存在
		if err.Error() == "redis: nil" {
			return nil, errors.New("nonce not found")
		}
		logx.Errorf("redis get nonce error: %v", err)
		return nil, err
	}

	// 2. 校验签名
	if err = VerifyEthWallet(req.Address, nonce, req.Signature); err != nil {
		return nil, err
	}

	// 3. 查询用户
	comerInfo, err := comer.FindComerByAddress(l.svcCtx.DB, req.Address)
	if err != nil {
		logx.Errorf("get comer by address failed: %v", err)
		return nil, err
	}

	var (
		isProfiled bool
		profile    *comerprofile.ComerProfile
		firstLogin = false
	)

	if comerInfo == nil {
		// 4. 新用户，创建
		newComer := &comer.Comer{
			Address: req.Address,
		}

		err = comer.InsertComer(l.svcCtx.DB, newComer)
		if err != nil {
			return nil, err
		}
		comerInfo = newComer
		isProfiled = false
		firstLogin = true

	} else {
		profile, err = comerprofile.FindComerProfile(l.svcCtx.DB, uint64(comerInfo.ID))
		if err != nil {
			logx.Errorf("get comer profile failed: %v", err)
			return nil, err
		}
		if profile != nil {
			isProfiled = true
		}
	}

	// 6. 删除 nonce
	_, err = l.svcCtx.RedisClient.Del(context.TODO(), req.Address).Result()
	if err != nil {
		logx.Errorf("redis remove nonce failed: %v", err)
	}

	// 7. 生成 jwt
	token, err := jwt.Sign(uint64(comerInfo.ID), l.svcCtx.Config.JWT.Secret, l.svcCtx.Config.JWT.Expired)
	if err != nil {
		logx.Errorf("generate jwt token failed: %v", err)
		return nil, err
	}
	// 8. 返回登录响应
	resp = &types.ComerLoginResponse{
		IsProfiled: isProfiled,
		Avatar:     "",
		Nick:       "",
		Address:    req.Address,
		Token:      token,
		ComerID:    uint64(comerInfo.ID),
		FirstLogin: firstLogin,
	}
	if profile != nil {
		resp.Avatar = profile.Avatar
		resp.Nick = profile.Name
	}
	return

}

func VerifyEthWallet(address, nonce, signature string) error {
	logx.Infof("address: %v, nonce: %v, signature: %v", address, nonce, signature)
	addrKey := common.HexToAddress(address)
	sig := hexutil.MustDecode(signature)
	if sig[64] == 27 || sig[64] == 28 {
		sig[64] -= 27
	}
	msg := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonce), nonce)
	msg256 := crypto.Keccak256([]byte(msg))
	pubKey, err := crypto.SigToPub(msg256, sig)
	if err != nil {
		return err
	}
	recoverAddr := crypto.PubkeyToAddress(*pubKey)
	if recoverAddr != addrKey {
		return errors.New("Address mismatch")
	}
	return nil
}
