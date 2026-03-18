package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

// 定义请求体结构
type reqBody struct {
	GoodsId int64 `json:"goodsId"`
	UserId  int64 `json:"userId"`
}

func main() {
	var wg sync.WaitGroup
	var successCount int32 // 记录抢购成功的次数
	var failCount int32    // 记录失败的次数

	totalRequests := 1000 // 模拟 1000 个并发用户
	targetUrl := "http://127.0.0.1:8000/api/app/seckill/do"

	fmt.Printf("🚀 开始向接口发起 %d 个并发抢购请求...\n", totalRequests)
	startTime := time.Now()

	// 瞬间拉起 1000 个协程
	for i := 1; i <= totalRequests; i++ {
		wg.Add(1)

		// 注意：把 i 当作 userId 传进去，模拟 1000 个不同的用户，绕过我们的防刷防重复购买机制
		go func(userId int) {
			defer wg.Done()

			body := reqBody{
				GoodsId: 1,
				UserId:  int64(userId),
			}
			jsonBytes, _ := json.Marshal(body)

			// 发起 HTTP POST 请求
			resp, err := http.Post(targetUrl, "application/json", bytes.NewBuffer(jsonBytes))
			if err != nil {
				atomic.AddInt32(&failCount, 1)
				return
			}
			defer resp.Body.Close()

			// 读取响应内容
			respData, _ := io.ReadAll(resp.Body)

			// 根据我们接口的返回格式判断是否成功 (GoFrame 默认 code:0 为成功)
			if bytes.Contains(respData, []byte(`"code":0`)) {
				atomic.AddInt32(&successCount, 1)
			} else {
				atomic.AddInt32(&failCount, 1)
			}
		}(i)
	}

	// 等待所有协程执行完毕
	wg.Wait()

	fmt.Printf("✅ 压测结束！耗时: %v\n", time.Since(startTime))
	fmt.Printf("🎉 抢购成功人数: %d\n", successCount)
	fmt.Printf("❌ 抢购失败人数 (手慢了/超卖报错): %d\n", failCount)
}
