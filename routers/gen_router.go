package routers

import (
	"github.com/xxjwxc/ginrpc"
)

func init() {
	ginrpc.SetVersion(1611675919)
	ginrpc.AddGenOne("HongdongController.Detail", "/hongdong/detail", []string{"get"})
	ginrpc.AddGenOne("HongdongController.GetDetail", "/hongdong/getDetail", []string{"get"})
	ginrpc.AddGenOne("HongdongController.List", "/hongdong/list", []string{"get"})
}
