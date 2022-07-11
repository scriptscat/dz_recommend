package es

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"

	"github.com/codfrm/cago"
	"github.com/codfrm/cago/config"
	"github.com/elastic/go-elasticsearch/v8"
)

type Config struct {
	Address  []string `yaml:"address"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Cert     string   `yaml:"cert"`
}

func Elasticsearch(ctx context.Context, cfg *config.Config) error {
	config := &Config{}
	if err := cfg.Scan("elasticsearch", config); err != nil {
		return err
	}
	ca, err := ioutil.ReadFile(config.Cert)
	if err != nil {
		return err
	}
	certs := x509.NewCertPool()
	tlsConfig := &tls.Config{}
	if ok := certs.AppendCertsFromPEM(ca); !ok {
		return err
	}
	tlsConfig.RootCAs = certs
	tlsConfig.InsecureSkipVerify = true
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: config.Address,
		Username:  config.Username,
		Password:  config.Password,
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	})
	if err != nil {
		return err
	}
	es = client
	return nil
}

var es *elasticsearch.Client

func Ctx(ctx cago.Context) *elasticsearch.Client {
	return es
}
