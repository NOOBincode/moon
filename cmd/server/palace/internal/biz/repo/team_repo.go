package repo

import (
	"context"

	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/do/model"
)

type TeamRepo interface {
	// GetUserTeamByID 查询用户指定团队信息
	GetUserTeamByID(ctx context.Context, userID, teamID uint32) (*model.SysTeamMember, error)

	// GetTeamRoleByUserID 查询用户指定团队角色
	GetTeamRoleByUserID(ctx context.Context, userID, teamID uint32) ([]*model.SysTeamMemberRole, error)

	// GetUserTeamRole 查询用户指定团队角色
	GetUserTeamRole(ctx context.Context, userID, teamID, roleID uint32) (*model.SysTeamMemberRole, error)
}