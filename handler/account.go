package handler

import (
	"context"

	"github.com/micro/go-micro/v2/logger"

	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

type Account struct {
	AuthService    proto.AuthService
	ProfileService proto.ProfileService
}

func (this *Account) Signup(_ctx context.Context, _req *proto.SignupRequest, _rsp *proto.SignupResponse) error {
	logger.Infof("Received Auth.Signup, username is %v", _req.Username)
	rsp, err := this.AuthService.Signup(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Account) Signin(_ctx context.Context, _req *proto.SigninRequest, _rsp *proto.SigninResponse) error {
	logger.Infof("Received Auth.Signin, username is %v, strategy is %v", _req.Username, _req.Strategy)
	rsp, err := this.AuthService.Signin(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Account) Signout(_ctx context.Context, _req *proto.SignoutRequest, _rsp *proto.SignoutResponse) error {
	logger.Infof("Received Auth.Signout, accessToken is %v", _req.AccessToken)
	rsp, err := this.AuthService.Signout(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Account) ResetPasswd(_ctx context.Context, _req *proto.ResetPasswdRequest, _rsp *proto.ResetPasswdResponse) error {
	logger.Infof("Received Auth.ResetPasswd, accessToken is %v", _req.AccessToken)
	rsp, err := this.AuthService.ResetPasswd(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Account) Query(_ctx context.Context, _req *proto.QueryProfileRequest, _rsp *proto.QueryProfileResponse) error {
	logger.Infof("Received Profile.Query, accessToken is %v", _req.AccessToken)
	rsp, err := this.ProfileService.Query(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Account) Update(_ctx context.Context, _req *proto.UpdateProfileRequest, _rsp *proto.UpdateProfileResponse) error {
	logger.Infof("Received Profile.Update, accessToken is %v", _req.AccessToken)
	rsp, err := this.ProfileService.Update(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}
