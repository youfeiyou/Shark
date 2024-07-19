package db

import (
	"context"
	pb "shark/pkg/proto/groupInfo"
)

type groupInfo struct {
}

var GroupInfoAPI groupInfo

const (
	tableGroupInfo = "gi_"
)

func (groupInfo) GetGroupInfo(ctx context.Context, groupCode uint64) (*pb.DbGroupInfo, error) {
	return nil, nil
}

func (groupInfo) UpdateGroupInfo(ctx context.Context, info *pb.DbGroupInfo) error {
	return nil
}
