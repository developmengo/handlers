package snap

import (
	"crypto"
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/logs"

	"reflect"
	"strings"
)

func SnapVerifySignature(ctx *context.Context) {
	var (
		res SnapValidationRes = SnapValidationRes{}
	)

	r := ctx.Input
	var (
		authorization string = r.Header("Authorization")
		timeStamp     string = strings.TrimSpace(r.Header("X-TIMESTAMP"))
	)

	if authorization == "" {
		ctx.Output.SetStatus(403)
		res.ResponseCode = "403"
		res.ResponseMessage = "Unauthorized. [403]"

		resBody, _ := json.Marshal(res)
		ctx.Output.Body(resBody)
		return
	}

	checkTimestamp := SnapXTimestamp(timeStamp)
	if checkTimestamp != nil {
		ctx.Output.SetStatus(403)
		resBody, _ := json.Marshal(checkTimestamp)
		ctx.Output.Body(resBody)
		return
	}

	headers := SymetricSignatureSnap{}
	headers.Url = strings.TrimSpace(r.URL())
	headers.Method = strings.TrimSpace(strings.ToUpper(r.Method()))
	headers.AccessToken = strings.TrimSpace(strings.Split(authorization, "Bearer ")[1])

	if headers.Method == "GET" {
		headers.Body = "{}"
	} else {
		headers.Body = Decrypt(r.Header("X-PARTNER"), beego.AppConfig.String("SECRET"))
	}
	headers.ClientSecret = strings.TrimSpace(beego.AppConfig.String("SECRET"))
	headers.TimeStamp = strings.TrimSpace(timeStamp)

	hmac512Body := headers.Method + ":" + headers.Url + ":svc-product:" + r.Header("X-PARTNER") + ":" + headers.TimeStamp
	logs.Info("hmacBody : ", hmac512Body)
	hmac512 := hmac.New(crypto.SHA512.New, []byte(headers.ClientSecret))
	hmac512.Write([]byte(strings.TrimSpace(hmac512Body)))

	signatureToken := base64.StdEncoding.EncodeToString(hmac512.Sum(nil))

	if ok := reflect.DeepEqual(headers.AccessToken, signatureToken); !ok {
		ctx.Output.SetStatus(403)
		res.ResponseCode = "403"
		res.ResponseMessage = "Invalid Signature"

		resBody, _ := json.Marshal(res)
		ctx.Output.Body(resBody)
		return
	}
}
