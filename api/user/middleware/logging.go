package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// Logging - ログの出力
func Logging() gin.HandlerFunc {
	return func(c *gin.Context) {
		buf, _ := ioutil.ReadAll(c.Request.Body)
		rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf)) // ログ出力で使用する用
		rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf)) // ログ出力以降で使用する用

		requestLogging(rdr1)

		c.Request.Body = rdr2
		c.Next()
	}
}

func requestLogging(reader io.Reader) {
	// json ファイルの読み込み -> 値がなければ return
	data, _ := ioutil.ReadAll(reader)
	if len(data) == 0 {
		return
	}

	// json を データ型に変換する
	params := make(map[string]interface{})
	if err := json.Unmarshal(data, &params); err != nil {
		log.Printf("error: JSONの整形に失敗しました")
		return
	}

	// password のみフィルターにかける
	if params["password"] != nil {
		params["password"] = "********"
	}

	if params["passwordConfirmation"] != nil {
		params["passwordConfirmation"] = "********"
	}

	// ログの整形・出力
	message := ""
	for key, value := range params {
		message += fmt.Sprintf("%s:'%s' ", key, value)
	}

	log.Printf("info: params: { %s}", message)
}
