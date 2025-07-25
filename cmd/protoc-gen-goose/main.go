package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	http "github.com/go-leo/goose/cmd/protoc-gen-goose/generator"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Fprintf(os.Stdout, "%v %v\n", filepath.Base(os.Args[0]), "v1.0.0")
		os.Exit(0)
	}

	var flags flag.FlagSet
	options := &protogen.Options{ParamFunc: flags.Set}
	options.Run(func(plugin *protogen.Plugin) error {
		plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
		return generate(plugin)
	})
}

func generate(plugin *protogen.Plugin) error {
	for _, file := range plugin.Files {
		if !file.Generate {
			continue
		}

		// http代码生成
		httpGenerator, err := http.NewGenerator(plugin, file)
		if err != nil {
			return err
		}
		if err := httpGenerator.Generate(); err != nil {
			return err
		}
	}
	return nil
}
