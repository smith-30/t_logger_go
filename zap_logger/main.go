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

	// Output is ...

	// {"level":"info","ts":1529129195.434518,"caller":"sample/main.go:27","msg":"Did Int","param":1111}
	// {"level":"info","ts":1529129195.434567,"caller":"sample/main.go:29","msg":"Did String","param":"ok"}
	// {"level":"info","ts":1529129195.4345996,"caller":"sample/main.go:31","msg":"Did Any","param":{"id":1,"name":"Obj"}}
}
