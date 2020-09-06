/*
@Time : 2020/9/6 下午10:10
@Author : zhanghf
@File : test
@Software: GoLand
*/
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Cfg struct {
	Cid               int      `json:"cid"`
	CidName           string   `json:"cidName"`
	MainSearch        []string `json:"main_search"`
	AdverbSearch      []string `json:"adverb_search"`
	TbSpiderCount     int      `json:"tb_spider_count"`
	TbSpiderMainCount int      `json:"tb_spider_main_count"`
	JdSpiderCount     int      `json:"jd_spider_count"`
	JdSpiderMainCount int      `json:"jd_spider_min_count"`
}

func main() {
	fs, err := os.Open("./分类采集词.csv")
	if err != nil {
		fmt.Println("打开文件错误", err)
		return
	}
	defer fs.Close()

	br := bufio.NewReader(fs)
	list := make([]Cfg, 0)
	isFirst := true
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		if isFirst {
			isFirst = false
			continue
		}
		arr := strings.Split(string(a), ",")
		if len(arr) < 8 {
			fmt.Println("当行文件内容小于8")
			continue
		}
		cid, _ := strconv.Atoi(arr[0])
		cidName := arr[1]

		tbSpiderCount, _ := strconv.Atoi(arr[2])
		tbAdverbSpiderCount, _ := strconv.Atoi(arr[3])
		jdSpiderCount, _ := strconv.Atoi(arr[3])
		jdAdverbSpiderCount, _ := strconv.Atoi(arr[4])

		mainSearchText := arr[6]
		mainSearchArr := strings.Split(mainSearchText, "、")
		adverbText := arr[7]
		adverbArr := strings.Split(adverbText, "、")
		list = append(list, Cfg{
			Cid:               cid,
			CidName:           cidName,
			MainSearch:        mainSearchArr,
			AdverbSearch:      adverbArr,
			TbSpiderMainCount: tbSpiderCount,
			TbSpiderCount:     tbAdverbSpiderCount,
			JdSpiderMainCount: jdSpiderCount,
			JdSpiderCount:     jdAdverbSpiderCount,
		})
	}

	content, _ := json.Marshal(list)
	ioutil.WriteFile("./分类配置.json", content, 0777)
}
