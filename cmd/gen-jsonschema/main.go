package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/invopop/jsonschema"
	"github.com/sozercan/aikit/pkg/aikit/config"
)

func main() {
	var r jsonschema.Reflector
	if err := r.AddGoComments("github.com/sozercan/aikit", "./"); err != nil {
		panic(err)
	}

	for _, spec := range []interface{}{
		&config.InferenceConfig{},
		&config.FineTuneConfig{},
	} {
		schema := r.Reflect(spec)
		dt, err := json.MarshalIndent(schema, "", "\t")
		if err != nil {
			panic(err)
		}

		if len(os.Args) > 1 {
			if err := os.MkdirAll(filepath.Dir(os.Args[1]), 0o755); err != nil {
				panic(err)
			}
			if err := os.WriteFile(os.Args[1], dt, 0o600); err != nil {
				panic(err)
			}
			return
		}
		fmt.Println(string(dt))
	}
}
