package utils

import (
	"math/rand"
	"time"
)

func GetRand(left int,right int) (result int){
	// 随机种子
	rand.Seed(time.Now().Unix())
	// 生成 20 个 [0, 100) 范围的伪随机数。
	result= rand.Intn(right-left)+left
	return result
}