package authorization

import (
	"context"

	pb "github.com/aide-cloud/moon/api/admin/authorization"
	"github.com/aide-cloud/moon/api/merr"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/biz/bo"
	"github.com/aide-cloud/moon/cmd/server/palace/internal/service/build"
	"github.com/aide-cloud/moon/pkg/helper/middleware"
	"github.com/aide-cloud/moon/pkg/utils/captcha"
)

type Service struct {
	pb.UnimplementedAuthorizationServer

	captchaBiz       *biz.CaptchaBiz
	authorizationBiz *biz.AuthorizationBiz
}

func NewAuthorizationService(
	captchaBiz *biz.CaptchaBiz,
	authorizationBiz *biz.AuthorizationBiz,
) *Service {
	return &Service{
		captchaBiz:       captchaBiz,
		authorizationBiz: authorizationBiz,
	}
}

func (s *Service) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginReply, error) {
	captchaInfo := req.GetCaptcha()
	// 校验验证码
	if err := s.captchaBiz.VerifyCaptcha(ctx, &bo.ValidateCaptchaParams{
		Id:    captchaInfo.GetId(),
		Value: captchaInfo.GetCode(),
	}); err != nil {
		return nil, err
	}

	// 执行登录逻辑
	loginJwtClaims, err := s.authorizationBiz.Login(ctx, &bo.LoginParams{
		Username:   req.GetUsername(),
		EnPassword: req.GetPassword(),
		Team:       req.GetTeamId(),
		TeamRole:   req.GetTeamRoleId(),
	})
	if err != nil {
		return nil, err
	}

	token, err := loginJwtClaims.JwtClaims.GetToken()
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		User:     build.NewUserBuild(loginJwtClaims.User).ToApi(),
		Token:    token,
		Redirect: "/",
	}, nil
}

func (s *Service) Logout(ctx context.Context, _ *pb.LogoutRequest) (*pb.LogoutReply, error) {
	jwtClaims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, bo.UnLoginErr
	}

	if err := s.authorizationBiz.Logout(ctx, &bo.LogoutParams{
		JwtClaims: jwtClaims,
	}); err != nil {
		return nil, merr.ErrorNotification("系统错误")
	}

	return &pb.LogoutReply{
		Redirect: "/login",
	}, nil
}

func (s *Service) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	jwtClaims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, bo.UnLoginErr
	}
	tokenRes, err := s.authorizationBiz.RefreshToken(ctx, &bo.RefreshTokenParams{
		JwtClaims: jwtClaims,
		Team:      req.GetTeamId(),
		TeamRole:  req.GetTeamRoleId(),
	})
	if err != nil {
		return nil, err
	}

	token, err := tokenRes.JwtClaims.GetToken()
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokenReply{
		Token: token,
		User:  build.NewUserBuild(tokenRes.User).ToApi(),
	}, nil
}

func (s *Service) Captcha(ctx context.Context, req *pb.CaptchaReq) (*pb.CaptchaReply, error) {
	generateCaptcha, err := s.captchaBiz.GenerateCaptcha(ctx, &bo.GenerateCaptchaParams{
		Type:  captcha.Type(req.GetCaptchaType()),
		Theme: captcha.Theme(req.GetTheme()),
		Size:  []int{int(req.GetHeight()), int(req.GetWidth())},
	})
	if err != nil {
		return nil, err
	}
	return &pb.CaptchaReply{
		Captcha:     generateCaptcha.Base64s,
		CaptchaType: req.GetCaptchaType(),
		Id:          generateCaptcha.Id,
	}, nil
}

// CheckPermission 权限检查
func (s *Service) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, bo.UnLoginErr
	}
	if claims.IsAdminRole() {
		return &pb.CheckPermissionReply{HasPermission: true}, nil
	}
	if err := s.authorizationBiz.CheckPermission(ctx, &bo.CheckPermissionParams{
		JwtClaims: claims,
		Operation: req.GetOperation(),
	}); err != nil {
		return nil, err
	}
	return &pb.CheckPermissionReply{HasPermission: true}, nil
}

// CheckToken 检查token
func (s *Service) CheckToken(ctx context.Context, _ *pb.CheckTokenRequest) (*pb.CheckTokenReply, error) {
	claims, ok := middleware.ParseJwtClaims(ctx)
	if !ok {
		return nil, bo.UnLoginErr
	}
	if err := s.authorizationBiz.CheckToken(ctx, &bo.CheckTokenParams{JwtClaims: claims}); err != nil {
		return nil, err
	}
	return &pb.CheckTokenReply{IsLogin: true}, nil
}