package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
)

func main() {
	settings := cli.New()

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(settings.RESTClientGetter(), settings.Namespace(), "secrets", log.Printf); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	chart, err := loader.Load("./gateway-chart")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	install := action.NewInstall(actionConfig)
	install.Namespace = "default"
	install.ReleaseName = "gateway"
	install.UseReleaseName = true

	vals := map[string]interface{}{}

	rawVals, err := os.ReadFile("./gateway-values.yaml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(rawVals, &vals)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	install.RunWithContext(context.Background(), chart, vals)
}
