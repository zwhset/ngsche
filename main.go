package main

import (
	"fmt"
	"github.com/ngsche/types"
	"strings"
)

func main() {
	t()
}

func t() {
	var Upstream1 []types.Node

	baidu1 := types.Node{
		Server:      "192.168.1.2",
		Port:        80,
		Weight:      1,
		FailTimeout: 3,
		SLowStart:   10,
		IsResolve:   true,
		IsBackup:    false,
		IsDown:      false,
	}

	baidu2 := types.Node{
		Server: "192.168.1.3",
		Port:   80,
		Weight: 2,
	}

	Upstream1 = append(Upstream1, baidu1, baidu2)

	upstream1 := UpStreamToString("www_baidu_com_us", Upstream1)
	fmt.Println(upstream1)

	var Upstream2 []types.Node

	weibo1 := types.Node{
		Server:      "192.168.2.2",
		Port:        80,
		Weight:      1,
		FailTimeout: 3,
		SLowStart:   10,
		IsResolve:   false,
		IsBackup:    true,
		IsDown:      true,
	}

	weibo2 := types.Node{
		Server:   "192.168.2.3",
		Port:     80,
		Weight:   2,
		IsBackup: true,
	}

	Upstream2 = append(Upstream2, weibo1, weibo2)

	upstream2 := UpStreamToString("www_weibo_com_us", Upstream2)
	fmt.Println()
	fmt.Println(upstream2)
}

func UpStreamToString(name string, ups []types.Node) (upstream string) {

	var nodes []string

	for _, v := range ups {
		var nodeInfo string
		var resolve string
		var backup string
		var down string

		if v.IsResolve {
			resolve = "resolve"
		}
		if v.IsBackup {
			backup = "backup"
		}
		if v.IsDown {
			down = "down"
		}

		if v.Weight == 0 {
			v.Weight = 1
		}
		if v.Port == 0 {
			v.Port = 80
		}
		if v.SLowStart == 0 {
			v.SLowStart = 10
		}
		if v.FailTimeout == 0 {
			v.FailTimeout = 3
		}

		nodeInfo = fmt.Sprintf(
			"server %s:%d weight=%d fail_timeout=%ds slow_start=%ds %s %s %s",
			v.Server, v.Port, v.Weight, v.FailTimeout, v.SLowStart, resolve, backup, down)
		nodes = append(nodes, nodeInfo)
	}

	upstream = fmt.Sprintf("upstream %s {\n", name)
	for _, node := range nodes {
		upstream = fmt.Sprintf("%s\t%s\n", upstream, strings.TrimSpace(node))
	}
	upstream += "}"
	return upstream
}
