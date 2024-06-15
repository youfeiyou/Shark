package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"shark/internal/file"
	pb "shark/pkg/proto/file"
)

func main() {
	conn, _ := grpc.Dial("127.0.0.1:16001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	cc := pb.NewSigServiceClient(conn)
	req := &pb.Req{
		Uin:     884322372,
		ResType: pb.ResourceType_IMAGE,
	}
	// 1. 获取上传文件需要的token
	rsp, err := cc.UpLoad(context.Background(), req)
	log.Printf("rsp %v err %v", rsp, err)
	jp := &file.PolicyToken{}
	err = json.Unmarshal([]byte(rsp.Result), jp)
	log.Printf("jp %+v", jp)

	// 2. 客户端利用拿到的token上传为文件到OSS服务器
	uploadToOSS(jp, "0")

	// 3. 客户端如果需要下文件，需要向后台获取签名
	rsp, err = cc.Download(context.Background(), req)
	log.Printf("rsp %+v err %v", rsp, err)
}

func uploadToOSS(jp *file.PolicyToken, filename string) {
	var b bytes.Buffer
	writer := multipart.NewWriter(&b)

	// 添加文本字段 k1=v1 和 k2=v2
	_ = writer.WriteField("key", jp.Directory+filename)
	_ = writer.WriteField("policy", jp.Policy)
	_ = writer.WriteField("OSSAccessKeyId", jp.AccessKeyId)
	_ = writer.WriteField("success_action_status", "200")
	_ = writer.WriteField("signature", jp.Signature)
	//_ = writer.WriteField("callback", jp.Callback)

	// 打开要上传的文件
	file, err := os.Open("./cmd/file/main.go")
	log.Println(os.ReadDir("./"))
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// 创建一个文件字段
	part, err := writer.CreateFormFile("file", "889myfile")
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return
	}

	// 将文件内容复制到文件字段
	_, err = io.Copy(part, file)
	if err != nil {
		fmt.Println("Error copying file:", err)
		return
	}
	// 关闭 writer 以结束 multipart 的部分
	err = writer.Close()
	if err != nil {
		fmt.Println("Error closing writer:", err)
		return
	}
	// 创建一个新的 POST 请求
	req, err := http.NewRequest("POST", jp.Host, &b)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// 设置 Content-Type 为 multipart/form-data 并包含 boundary
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// 发送请求并获取响应
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()
	log.Printf("resp %v", resp)
	// 读取并输出响应
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}
	fmt.Println("Response:", string(body))
}
