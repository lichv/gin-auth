package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"math/rand"
	url2 "net/url"
	"reflect"
	"strconv"
	"time"
)

func In(haystack interface{}, needle interface{}) (bool, error) {
	sVal := reflect.ValueOf(haystack)
	kind := sVal.Kind()
	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < sVal.Len(); i++ {
			if sVal.Index(i).Interface() == needle {
				return true, nil
			}
		}

		return false, nil
	}

	return false, errors.New("ErrUnSupportHaystack")
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}

func GeneTimeUUID() string {
	now := time.Now().UnixNano()/1000
	return strconv.FormatUint(uint64(now),36)+strconv.Itoa(rand.New(rand.NewSource(now)).Intn(90)+10)
}

func URLAppendParams(url string, key ,value string) (string,error) {
	l, err := url2.Parse(url)
	if err != nil {
		return url,err
	}

	query := l.Query()
	query.Set(key,value)
	//u, err := url2.Parse(query.Encode())
	//if err != nil {
	//	return url,err
	//}
	//return fmt.Sprintf("%v",u),nil
	encodeurl := l.Scheme + "://" + l.Host + "?" + query.Encode()
	return encodeurl,nil
}
func GetMapFromJson(c *gin.Context) (m map[string]interface{},err error) {
	buf := make([]byte,1024)
	n,_ := c.Request.Body.Read(buf)
	c.Request.Body = ioutil.NopCloser(bytes.NewReader(buf[:n]))
	j := buf[0:n]
	err = json.Unmarshal(j, &m)
	if err != nil {
		return map[string]interface{}{},err
	}
	return m,err
}

// Strval 获取变量的字符串值
// 浮点型 3.0将会转换成字符串3, "3"
// 非数值或字符类型的变量将会被转换成JSON格式字符串
func Strval(value interface{}) string {
	// interface 转 string
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}

func BoolVal(value interface{}) bool {
	var key = false
	if value == nil {
		return key
	}
	switch value.(type) {
	case float64:
		ft := value.(float64)
		it := strconv.FormatFloat(ft, 'f', -1, 64)
		key = it == "0.0"
	case float32:
		ft := value.(float32)
		it := strconv.FormatFloat(float64(ft), 'f', -1, 64)
		key = it == "0.0"
	case int:
		it := value.(int)
		key = it != 0
	case uint:
		it := value.(uint)
		key = it != 0
	case int8:
		it := value.(int8)
		key = it != 0
	case uint8:
		it := value.(uint8)
		key = it != 0
	case int16:
		it := value.(int16)
		key = it != 0
	case uint16:
		it := value.(uint16)
		key = it != 0
	case int32:
		it := value.(int32)
		key = it != 0
	case uint32:
		it := value.(uint32)
		key = it != 0
	case int64:
		it := value.(int64)
		key = it != 0
	case uint64:
		it := value.(uint64)
		key = it != 0
	case string:
		it := value.(string)
		key = it != "" && it != "false"
	case []byte:
		it := value.([]byte)
		key = len(it) > 0
	default:
		newValue, _ := json.Marshal(value)
		it := string(newValue)
		key = len(it) > 0
	}
	return key
}

