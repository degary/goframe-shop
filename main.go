package main

import (
	_ "github.com/degary/goframe-shop/internal/logic"
	_ "github.com/degary/goframe-shop/internal/packed"
	"github.com/gogf/gf/v2/os/gctx"

	"github.com/degary/goframe-shop/internal/cmd"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.New())
}
