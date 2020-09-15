package handler

import (
	"context"
	"omo-bla-acm/config"

	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"

	proto "github.com/xtech-cloud/omo-msp-account/proto/account"
)

type Auth struct {
	Service    proto.AuthService
}

type Profile struct {
	Service proto.ProfileService
}


func SetupAccountHandler(_server server.Server, _client client.Client) {
    authHandler := new(Auth)
    authHandler.Service = proto.NewAuthService(config.Schema.MSA.Account, _client)
	proto.RegisterAuthHandler(_server, authHandler)

    profileHandler := new(Profile)
    profileHandler.Service = proto.NewProfileService(config.Schema.MSA.Account, _client)
	proto.RegisterProfileHandler(_server, profileHandler)
}

func (this *Auth) Signup(_ctx context.Context, _req *proto.SignupRequest, _rsp *proto.SignupResponse) error {
	logger.Infof("Received Auth.Signup, username is %v", _req.Username)
	rsp, err := this.Service.Signup(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Auth) Signin(_ctx context.Context, _req *proto.SigninRequest, _rsp *proto.SigninResponse) error {
	logger.Infof("Received Auth.Signin, username is %v, strategy is %v", _req.Username, _req.Strategy)
	rsp, err := this.Service.Signin(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Auth) Signout(_ctx context.Context, _req *proto.SignoutRequest, _rsp *proto.SignoutResponse) error {
	logger.Infof("Received Auth.Signout, accessToken is %v", _req.AccessToken)
	rsp, err := this.Service.Signout(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Auth) ResetPasswd(_ctx context.Context, _req *proto.ResetPasswdRequest, _rsp *proto.ResetPasswdResponse) error {
	logger.Infof("Received Auth.ResetPasswd, accessToken is %v", _req.AccessToken)
	rsp, err := this.Service.ResetPasswd(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Profile) Query(_ctx context.Context, _req *proto.QueryProfileRequest, _rsp *proto.QueryProfileResponse) error {
	logger.Infof("Received Profile.Query, accessToken is %v", _req.AccessToken)
	rsp, err := this.Service.Query(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}

func (this *Profile) Update(_ctx context.Context, _req *proto.UpdateProfileRequest, _rsp *proto.UpdateProfileResponse) error {
	logger.Infof("Received Profile.Update, accessToken is %v", _req.AccessToken)
	rsp, err := this.Service.Update(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
	return err
}
