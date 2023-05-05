package utils

import (
	"context"
	"fmt"
	"time"
)

const CAPTCHA = "captcha:"

var ctx = context.Background()

type CaptchaRedisStore struct {
}

// Set 实现设置captcha的方法
func (r CaptchaRedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	//time.Minute*5：有效时间5分钟
	err := RedisDb.Set(ctx, key, value, time.Minute*5).Err()

	return err
}

// Get 实现获取captcha的方法
func (r CaptchaRedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		//clear为true，验证通过，删除这个验证码
		err := RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify 实现验证captcha的方法
func (r CaptchaRedisStore) Verify(id, answer string, clear bool) bool {
	v := CaptchaRedisStore{}.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}
