package repoimpl

import (
	"context"
	"time"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-family/moon/cmd/server/palace/internal/biz/repository"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/pkg/util/types"
)

// NewCaptchaRepository 创建验证码操作
func NewCaptchaRepository(data *data.Data) repository.Captcha {
	return &captchaRepositoryImpl{
		data: data,
	}
}

type captchaRepositoryImpl struct {
	data *data.Data
}

func (l *captchaRepositoryImpl) CreateCaptcha(ctx context.Context, captcha *bo.ValidateCaptchaItem, duration time.Duration) error {
	bs, err := captcha.MarshalBinary()
	if !types.IsNil(err) {
		return err
	}
	return l.data.GetCacher().Client().Set(ctx, captcha.ID, string(bs), duration).Err()
}

func (l *captchaRepositoryImpl) GetCaptchaByID(ctx context.Context, id string) (*bo.ValidateCaptchaItem, error) {
	str, err := l.data.GetCacher().Client().Get(ctx, id).Result()
	if !types.IsNil(err) {
		return nil, err
	}
	var captcha bo.ValidateCaptchaItem
	if err = captcha.UnmarshalBinary([]byte(str)); !types.IsNil(err) {
		return nil, err
	}
	return &captcha, nil
}
