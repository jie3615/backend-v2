package authorizations

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"metaLand/app/api/internal/svc"
	"metaLand/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNonceByAddressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取钱包地址登录用的 nonce
func NewGetNonceByAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNonceByAddressLogic {
	return &GetNonceByAddressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetNonceByAddressLogic) GetNonceByAddress(req *types.GetNonceByAddressRequest) (resp *types.WalletNonceResponse, err error) {
	nonce, err := l.svcCtx.RedisClient.Get(context.TODO(), req.WalletAddress).Result()
	if err != nil {
		if err.Error() != "redis: nil" {
			return nil, err
		}
	}

	if nonce == "" {
		id, err := l.svcCtx.SF.NextID()
		if err != nil {
			return nil, err
		}
		nonce, err = createNonce(req.WalletAddress, id)
		if err != nil {
			return nil, err
		}
		err = l.svcCtx.RedisClient.Set(context.TODO(), req.WalletAddress, nonce, time.Hour*24).Err()
		if err != nil {
			return nil, err
		}
	}
	logx.Infof("get nonce by address: %s, nonce: %s", req.WalletAddress, nonce)
	return &types.WalletNonceResponse{
		Nonce: nonce,
	}, nil
}

// 创建nonce
func createNonce(address string, id uint64) (string, error) {
	// 1. 将id转为36进制字符串（包含数字和小写字母）
	idStr := strconv.FormatUint(id, 36)
	// 2. 将address去掉0x前缀，取后8位
	addrPart := address
	if len(addrPart) > 2 && addrPart[:2] == "0x" {
		addrPart = addrPart[2:]
	}
	if len(addrPart) > 8 {
		addrPart = addrPart[len(addrPart)-8:]
	}
	// 3. 拼接后打乱顺序，取前10位
	raw := idStr + addrPart
	// 使用局部随机数生成器，避免全局rand.Seed
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	runes := []rune(raw)
	r.Shuffle(len(runes), func(i, j int) { runes[i], runes[j] = runes[j], runes[i] })
	// 取前10位，不足补0
	nonce := string(runes)
	if len(nonce) > 10 {
		nonce = nonce[:10]
	} else if len(nonce) < 10 {
		nonce = leftPad(nonce, "0", 10)
	}
	return nonce, nil
}

func leftPad(s, pad string, length int) string {
	for len(s) < length {
		s = pad + s
	}
	return s
}
