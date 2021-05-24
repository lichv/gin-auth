package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
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
