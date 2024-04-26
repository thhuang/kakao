package main

import (
	"github.com/thhuang/kakao/internal/saide"
	"github.com/thhuang/kakao/pkg/util/ctx"
)

func main() {
	ctx, cancel := ctx.WithCancel(ctx.Background())
	defer cancel()
	saide.New(ctx).Run(ctx)
}
