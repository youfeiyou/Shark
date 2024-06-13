package file

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"hash"
	"io"
	"os"
	"time"
)

const (
	endpoint           = "oss-cn-shenzhen.aliyuncs.com"
	bucket             = "shark0611"
	ossAccessKeyId     = ""
	ossAccessKeySecret = ""
	expireTime         = 60
	callBackURL        = ""
	host               = "http://shark0611.oss-cn-shenzhen.aliyuncs.com"
)

func init() {
	_ = os.Setenv("OSS_ACCESS_KEY_ID", ossAccessKeyId)
	_ = os.Setenv("OSS_ACCESS_KEY_SECRET", ossAccessKeySecret)
}

// url for client download file
func generateURL(path string) (string, error) {
	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	client, err := oss.New(endpoint, "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	bucketName := bucket
	objectName := path
	// 将Object下载到本地文件，并保存到指定的本地路径中。如果指定的本地文件存在会覆盖，不存在则新建。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	// 生成用于下载的签名URL，并指定签名URL的有效时间为60秒。
	signedURL, err := bucket.SignURL(objectName, oss.HTTPGet, expireTime)
	if err != nil {
		return "", err
	}
	fmt.Printf("Sign Url:%s\n", signedURL)
	return signedURL, nil
}

func getGmtIso8601(expireEnd int64) string {
	var tokenExpire = time.Unix(expireEnd, 0).UTC().Format("2006-01-02T15:04:05Z")
	return tokenExpire
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

type PolicyToken struct {
	AccessKeyId string `json:"accessid"`
	Host        string `json:"host"`
	Expire      int64  `json:"expire"`
	Signature   string `json:"signature"`
	Policy      string `json:"policy"`
	Directory   string `json:"dir"`
	Callback    string `json:"callback"`
}

type CallbackParam struct {
	CallbackUrl      string `json:"callbackUrl"`
	CallbackBody     string `json:"callbackBody"`
	CallbackBodyType string `json:"callbackBodyType"`
}

func get_policy_token(dir string) string {
	expireEnd := time.Now().Unix() + expireTime
	var tokenExpire = getGmtIso8601(expireEnd)
	//create post policy json
	var config ConfigStruct
	config.Expiration = tokenExpire
	var condition []string
	condition = append(condition, "starts-with")
	condition = append(condition, "$key")
	condition = append(condition, dir)
	config.Conditions = append(config.Conditions, condition)

	//calucate signature
	result, err := json.Marshal(config)
	debyte := base64.StdEncoding.EncodeToString(result)
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(ossAccessKeySecret))
	io.WriteString(h, debyte)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	var callbackParam CallbackParam
	callbackParam.CallbackUrl = callBackURL
	callbackParam.CallbackBody = "filename=${object}&size=${size}&mimeType=${mimeType}&height=${imageInfo.height}&width=${imageInfo.width}"
	callbackParam.CallbackBodyType = "application/x-www-form-urlencoded"
	callbackStr, err := json.Marshal(callbackParam)
	if err != nil {
		fmt.Println("callback json err:", err)
	}
	callbackBase64 := base64.StdEncoding.EncodeToString(callbackStr)
	var policyToken PolicyToken
	policyToken.AccessKeyId = ossAccessKeyId
	policyToken.Host = host
	policyToken.Expire = expireEnd
	policyToken.Signature = string(signedStr)
	policyToken.Directory = dir
	policyToken.Policy = debyte
	policyToken.Callback = callbackBase64
	response, err := json.Marshal(policyToken)
	if err != nil {
		fmt.Println("json err:", err)
	}
	return string(response)
}
