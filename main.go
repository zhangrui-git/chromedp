package main

import (
	"context"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
	"time"
)

func main() {
	ctx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		//chromedp.ExecPath(),
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", true),
		//chromedp.Flag("blink-settings", "imagesEnabled=false"),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36"),
	)
	ctx, cancel = chromedp.NewContext(ctx, chromedp.WithDebugf(log.Printf))
	defer cancel()
	chromedp.Run(ctx, chromedp.Sleep(time.Second))
	kw := "天气"
	//html := ""
	title := ""
	//var nodes []*cdp.Node
	err := chromedp.Run(
		ctx,
		chromedp.Tasks{
			chromedp.Navigate(`https://www.baidu.com`),
			chromedp.WaitVisible(`kw`, chromedp.ByID),
			chromedp.SetValue(`kw`, kw, chromedp.ByID),
			chromedp.Sleep(time.Second),
			chromedp.Click(`su`, chromedp.ByID),
			chromedp.Sleep(time.Second * 3),
			//chromedp.Evaluate(`document.getElementsByTagName("title")[0].text;`, &title),
			chromedp.Text(`document.querySelector("#\\31  > div.op_weather4_twoicon_container_div > div.op_weather4_twoicon > a.op_weather4_twoicon_today.OP_LOG_LINK > p.op_weather4_twoicon_date")`, &title, chromedp.ByJSPath),
			// chromedp.OuterHTML(`div.op_weather4_twoicon_container_div > div.op_weather4_twoicon`, &html, chromedp.NodeVisible),
			// chromedp.Nodes(`div.op_weather4_twoicon_container_div > div.op_weather4_twoicon > a.op_weather4_twoicon_day.OP_LOG_LINK`, &nodes),
		},
		)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(strings.TrimSpace(html))
	log.Println(strings.TrimSpace(title))
	//for _, node := range nodes {
	//	jsonStr,_ := node.MarshalJSON()
	//	fmt.Println(string(jsonStr))
	//}

	//chromedp.ByFunc
	//chromedp.ByID
	//chromedp.ByJSPath
	//chromedp.ByNodeID
	//chromedp.ByQuery
	//chromedp.ByQueryAll
	//chromedp.BySearch
}
