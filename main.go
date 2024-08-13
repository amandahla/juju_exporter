package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	apiclient "github.com/juju/juju/api/client/client"
	"gopkg.in/yaml.v2"
)

var (
	addr       = ":9970"
	configFile = "juju_exporter.yaml"
)

func check(status string, accepted []string) float64 {
	if len(accepted) == 0 {
		return 1
	}
	for _, ok := range accepted {
		if status == ok {
			return 1
		}
	}
	return 0
}

func main() {
	conf := &config{}

	configFile = os.Getenv("JUJU_EXPORTER_CONFIG")
	if configFile == "" {
		configFile = "juju_exporter.yaml"
	}

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	if err := yaml.UnmarshalStrict(b, conf); err != nil {
		panic(err)
	}

	http.Handle("/metrics", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		model := req.URL.Query().Get("model")
		if model == "" {
			model = conf.Default
		}

		modelConf, ok := conf.Models[model]
		if !ok {
			http.Error(rw, fmt.Sprintf("error: model '%s' not found", model), http.StatusNotFound)
			return
		}
		registry := newRegistry(model, modelConf.ModelUUID)

		conn, err := modelConf.newClient()
		if err != nil {
			log.Printf("Error connecting to model %s: %v", model, err)
			http.Error(rw, "error: could not connect to Juju", http.StatusExpectationFailed)
			return
		}
		defer func() {
			if err := conn.Close(); err != nil {
				log.Printf("Error terminating Juju client %s: %v", model, err)
			}
		}()

		client := apiclient.NewClient(conn, nil)

		status, err := client.Status(nil)
		if err != nil {
			log.Fatalf("Error requesting status: %s", err)
		}
		if err != nil {
			log.Printf("Error retrieving Juju status for model %s: %v", model, err)
			http.Error(rw, "error: could not retrieve Juju status", http.StatusExpectationFailed)
			return
		}
		registry.parseStatus(status)

		promhttp.HandlerFor(
			registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
			},
		).ServeHTTP(rw, req)
	}))

	log.Printf("Listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
