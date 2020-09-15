package handler

import (
	"context"
	"omo-bla-acm/config"

	"github.com/micro/go-micro/v2/server"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"

	proto "github.com/xtech-cloud/omo-msp-group/proto/group"
)

type Collection struct {
	Service    proto.CollectionService
}

type Member struct {
	Service proto.MemberService
}

func SetupGroupHandler(_server server.Server, _client client.Client) {
    collectionHandler := new(Collection)
    collectionHandler.Service = proto.NewCollectionService(config.Schema.MSA.Group, _client)
	proto.RegisterCollectionHandler(_server, collectionHandler)

    memberHandler := new(Member)
    memberHandler.Service = proto.NewMemberService(config.Schema.MSA.Group, _client)
	proto.RegisterMemberHandler(_server, memberHandler)
}


func (this *Collection) Make(_ctx context.Context, _req *proto.CollectionMakeRequest, _rsp *proto.BlankResponse) error {
	return ErrNotImplemented
}

func (this *Collection) List(_ctx context.Context, _req *proto.CollectionListRequest, _rsp *proto.CollectionListResponse) error {
	return ErrNotImplemented
}

func (this *Collection) Remove(_ctx context.Context, _req *proto.CollectionRemoveRequest, _rsp *proto.BlankResponse) error {
	return ErrNotImplemented
}

func (this *Collection) Get(_ctx context.Context, _req *proto.CollectionGetRequest, _rsp *proto.CollectionGetResponse) error {
	return ErrNotImplemented
}

func (this *Member) Add(_ctx context.Context, _req *proto.MemberAddRequest, _rsp *proto.BlankResponse) error {
	return ErrNotImplemented
}

func (this *Member) List(_ctx context.Context, _req *proto.MemberListRequest, _rsp *proto.MemberListResponse) error {
	return ErrNotImplemented
}

func (this *Member) Remove(_ctx context.Context, _req *proto.MemberRemoveRequest, _rsp *proto.BlankResponse) error {
	return ErrNotImplemented
}

func (this *Member) Get(_ctx context.Context, _req *proto.MemberGetRequest, _rsp *proto.MemberGetResponse) error {
	return ErrNotImplemented
}

func (this *Member) Where(_ctx context.Context, _req *proto.MemberWhereRequest, _rsp *proto.MemberWhereResponse) error {
	logger.Infof("Received Member.Where, req is %v", _req)
	rsp, err := this.Service.Where(_ctx, _req)
	if nil == err {
		*_rsp = *rsp
	}
    return err
}

