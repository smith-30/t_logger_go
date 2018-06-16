package sample_test

import (
	"reflect"
	"testing"

	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"

	"github.com/smith-30/t_logger/zap_logger/handler/sample"
)

func Test_handler_Do(t *testing.T) {
	type field struct {
		msg    string
		params []zapcore.Field
	}
	type args struct {
		param interface{}
	}
	tests := []struct {
		name    string
		fields  []field
		args    args
		wantErr bool
	}{
		{
			name: "int",
			fields: []field{
				field{
					msg: "Did Int",
					params: []zapcore.Field{
						zapcore.Field{
							Key:     "param",
							Type:    zapcore.Int64Type,
							Integer: 1,
						},
					},
				},
			},
			args: args{
				param: 1,
			},
			wantErr: false,
		},
		{
			name: "Did String",
			fields: []field{
				field{
					msg: "Did String",
					params: []zapcore.Field{
						zapcore.Field{
							Key:    "param",
							Type:   zapcore.StringType,
							String: "string",
						},
					},
				},
			},
			args: args{
				param: "string",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			core, recorded := observer.New(zapcore.InfoLevel)
			zl := zap.New(core)
			h := sample.NewHandler(sample.Config{Logger: zl})
			if err := h.Do(tt.args.param); (err != nil) != tt.wantErr {
				t.Errorf("handler.Do() error = %v, wantErr %v", err, tt.wantErr)
			}

			for idx, logs := range recorded.All() {
				t.Logf("%v\n", logs.Message)

				for i, log := range logs.Context {
					if !reflect.DeepEqual(log, tt.fields[idx].params[i]) {
						t.Errorf("\n exp %#v\nact %#v\n", tt.fields[idx].params[i], log)
					}
				}
			}
		})
	}
}
