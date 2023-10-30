package main

import (
	"context"
	"github.com/wangjc0216/errlog"
	"github.com/wangjc0216/rulemachine"
)

func main() {
	rulemachine.Execute(context.TODO())
	errlog.Errorf(context.TODO(), "say something ,it doesn't matter")
}
