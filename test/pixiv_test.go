package example

import (
	"context"
	"fmt"
	"github.com/NateScarlet/pixiv/pkg/artwork"
	"github.com/NateScarlet/pixiv/pkg/client"
	"testing"
)


func TestPixiv(t *testing.T) {


	// 使用 PHPSESSID Cookie 登录 (推荐)。
	c := &client.Client{}
	c.SetDefaultHeader("User-Agent", client.DefaultUserAgent)
	//c.Login("1020224260@qq.com", "qq1020224260")
	c.SetPHPSESSID("11517896_d8PaKCTeRqKDc2SpRx835nVtFdHEnjd3")

	// 所有查询从 context 获取客户端设置, 如未设置将使用默认客户端。
	var ctx = context.Background()
	ctx = client.With(ctx, c)


	rank := &artwork.Rank{Mode: "weekly_r18"}
	rank.Fetch(ctx)

	fmt.Printf("len:%v",len(rank.Items))
	if len(rank.Items)!=0{
		fmt.Printf("rank.Items[0].Rank:%v",rank.Items[0].Rank)
		fmt.Printf("rank.Items[0].PreviousRank:%v",rank.Items[0].PreviousRank)
		fmt.Printf("rank.Items[0].Artwork%v",rank.Items[0].Artwork)
	}

	// 搜索画作
	result, _:= artwork.Search(ctx, "パチュリー・ノーレッジ")
	//result.JSON // json return data.
	result.Artworks() // []artwork.Artwork，只有部分数据，通过 `Fetch` `FetchPages` 方法获取完整数据。
	artwork.Search(ctx, "パチュリー・ノーレッジ", artwork.SearchOptionPage(2)) // 获取第二页
	fmt.Printf("result.JSON:%v",result.JSON)
}
