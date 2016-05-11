package gogt_sdk

import (
	"fmt"
	"net/http"
	"time"
	"strings"
	"crypto/md5"
	"math/rand"
	"strconv"
)

const (
	FN_CHALLENGE = "geetest_challenge"
	FN_VALIDATE = "geetest_validate"
	FN_SECCODE = "geetest_seccode"

	GT_STATUS_SESSION_KEY = "gt_server_status"

	API_URL = "http://api.geetest.com"
	REGISTER_HANDLER = "/register.php"
	VALIDATE_HANDLER = "/validate.php"

	VERSION = "3.2.0"
)

type GeetestLib struct {
	privateKey  string
	captchaID   string
	sdkVersion  string
	responseStr string
}

//CreateGeeTest 创建GeetestLib实例
func CreateGeeTest(privateKey, captchaID string) *GeetestLib {
	return &GeetestLib{
		privateKey:privateKey,
		captchaID:captchaID,
		sdkVersion:VERSION,
		responseStr:"",
	}
}

//PreProcess 验证初始化预处理.
func (gt *GeetestLib)PreProcess() {

}

func (gt *GeetestLib)register(userid ...string)(int,string) {
	challenge := gt.registerChallenge(userid...)
	if len(challenge) != 32{
		return 0,gt.makeFailChallenge()
	}
	return 2,gt.md5Encode(append(challenge,[]byte(gt.privateKey)))
}

func (gt *GeetestLib)GetResponseStr()string {
	return gt.responseStr
}

func (gt *GeetestLib)makeFailChallenge()string {
	rand.Seed(time.Now().Unix())
	rnd1 := rand.Intn(100)
	rnd2 := rand.Intn(100)
	md5_str1 := gt.md5Encode([]byte(strconv.Itoa(rnd1)))
	md5_str2 := gt.md5Encode([]byte(strconv.Itoa(rnd2)))
	challenge := md5_str1 + md5_str2[0:2]
	return challenge
}

func (gt *GeetestLib)makeResponseFormat() {

}

//registerChallenge
func (gt *GeetestLib)registerChallenge(userid ...string)(respbytes []byte){
	var registerURL string
	if len(userid) == 1 {
		registerURL = fmt.Sprintf("%s%s?gt=%s&user_id=%s", API_URL, REGISTER_HANDLER, gt.captchaID, userid[0])
	} else {
		registerURL = fmt.Sprintf("%s%s?gt=%s", API_URL, REGISTER_HANDLER, gt.captchaID)
	}
	client := http.Client{Timeout: 2 * time.Second }
	resp, err := client.Get(registerURL)
	if err != nil || resp.StatusCode != 200 {
		return
	}
	_, err = resp.Body.Read(respbytes)
	if err != nil {
		return
	}
	return
}

func (gt *GeetestLib)SuccessValidate() {

}

func (gt *GeetestLib)postValues() {

}

func (gt *GeetestLib)checkResult() {

}

func (gt *GeetestLib)FailbackValidate() {

}

func (gt *GeetestLib)checkPara() {

}

func (gt *GeetestLib)validateFailImage() {

}

func (gt *GeetestLib)md5Encode(values []byte)string {
	return fmt.Sprintf("%x",md5.Sum(values))
}

func (gt *GeetestLib)decodeRandBase() {

}

func (gt *GeetestLib)decodeResponse() {

}













