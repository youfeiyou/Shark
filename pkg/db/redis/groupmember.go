package db

import "context"

type groupMemberAPI struct {
}

var GroupMemeberAPI groupMemberAPI

func (groupMemberAPI) addGroupMember(ctx context.Context, uin, group uint64) error {

	return nil
}

func (groupMemberAPI) delGroupMember(ctx context.Context, uin, group uint64) error {

	return nil
}

func (groupMemberAPI) updateGroupMember(ctx context.Context) {

}
