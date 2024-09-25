package repository

import (
	"context"
	"github.com/aide-family/moon/pkg/palace/model/bizmodel"

	"github.com/aide-family/moon/cmd/server/palace/internal/biz/bo"
)

type TeamInvite interface {
	// InviteUser 邀请用户加入团队
	InviteUser(ctx context.Context, params *bo.InviteUserParams) error
	// UpdateInviteStatus 更新邀请状态
	UpdateInviteStatus(ctx context.Context, params *bo.UpdateInviteStatusParams) error
	// InviteList 受邀请列表
	InviteList(ctx context.Context, params *bo.QueryInviteListParams)
	// GetInviteUserByUserIdAndType 获取邀请用户信息
	GetInviteUserByUserIdAndType(ctx context.Context, params *bo.InviteUserParams) (*bizmodel.SysTeamInvite, error)
}