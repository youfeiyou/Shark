package file

import (
	"context"
	"github.com/pkg/errors"
	"log"
	pb "shark/pkg/proto/file"
	"strconv"
)

const (
	defaultImageID = "0"
)

type Server struct {
	pb.UnimplementedSigServiceServer
}

func (Server) UpLoad(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if err := checkReq(req); err != nil {
		return rsp, err
	}
	// to do: check uin Permissions
	typ := req.GetResType()
	dir := ""
	if typ == pb.ResourceType_IMAGE || typ == pb.ResourceType_PERSONAL_FILE {
		dir = strconv.FormatUint(req.GetUin(), 10) + "/"
	} else {
		dir = strconv.FormatUint(req.GetGroupId(), 10) + "/"
	}
	token, err := getPolicyToken(dir)
	if err != nil {
		log.Printf("getPolicyToken fail %v", err)
		return rsp, err
	}
	rsp.Result = token
	return rsp, nil

}

func (Server) Download(ctx context.Context, req *pb.Req) (*pb.Rsp, error) {
	rsp := &pb.Rsp{}
	if err := checkReq(req); err != nil {
		return rsp, err
	}
	url, err := generateURL(generateDownloadPath(req))
	if err != nil {
		log.Printf("generateURL fail %v", err)
		return rsp, err
	}
	rsp.Result = url
	return rsp, nil
}

func checkReq(req *pb.Req) error {
	if req.GetResType() <= pb.ResourceType_INVALID || req.GetResType() > pb.ResourceType_GROUP_FILE {
		return errors.New("invalid resource type")
	}
	// to do: check uin Permissions
	return nil
}
func generateDownloadPath(req *pb.Req) string {
	path := ""
	if req.GetResType() == pb.ResourceType_IMAGE {
		path = strconv.FormatUint(req.GetUin(), 10) + "/" + defaultImageID
	} else if req.GetResType() == pb.ResourceType_GROUP_FILE {
		path = strconv.FormatUint(req.GetGroupId(), 10) + "/" + strconv.FormatUint(uint64(req.GetFileId()), 10)
	} else if req.GetResType() == pb.ResourceType_PERSONAL_FILE {
		path = strconv.FormatUint(req.GetUin(), 10) + "/" + strconv.FormatUint(uint64(req.GetFileId()), 10)
	}
	return path
}
