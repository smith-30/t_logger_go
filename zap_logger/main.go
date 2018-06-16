package main

import (
	domain_sample "github.com/smith-30/t_logger/zap_logger/domain/sample"
	"github.com/smith-30/t_logger/zap_logger/handler/sample"
	"go.uber.org/zap"
)

func main() {
	l, _ := zap.NewProduction()
	h := sample.NewHandler(sample.Config{Logger: l})
	h.Do(1111)
	h.Do("ok")

	o := domain_sample.Sample{
		ID:   1,
		Name: "Obj",
	}
	h.Do(o)
}
