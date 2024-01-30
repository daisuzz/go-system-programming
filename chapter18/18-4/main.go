package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// 5秒
	_ = 5 * time.Second

	// 10ミリ秒
	_ = 10 * time.Millisecond

	// 10分30秒
	_, _ = time.ParseDuration("10m30s")

	// 現在時刻
	_ = time.Now()

	// 現在日時
	_ = time.Date(2017, time.August, 26, 11, 50, 30, 0, time.Local)

	// フォーマットを指定してパース
	_, _ = time.Parse(time.Kitchen, "11:30AM")

	// Epochタイムから作成
	_ = time.Unix(1503673200, 0)

	// 3時間後の時間
	fmt.Println(time.Now().Add(3 * time.Hour))

	// ファイル変更日時が何日前か知る
	fileInfo, _ := os.Stat(".vimrc")
	fmt.Printf("%v前\n", time.Now().Sub(fileInfo.ModTime()))

	// 時間を1時間ごとに丸める
	fmt.Println(time.Now().Round(time.Hour))

	// スリープ
	fmt.Println("waiting 5 seconds")
	time.Sleep(5 * time.Second)
	fmt.Println("done")

	// チャネルを使ったタイマー
	fmt.Println("waiting 5 seconds")
	after := time.After(5 * time.Second)
	<-after
	fmt.Println("done")

	// tick
	fmt.Println("waiting 5 seconds")
	for now := range time.Tick(5 * time.Second) {
		fmt.Println("now:", now)
	}
}
