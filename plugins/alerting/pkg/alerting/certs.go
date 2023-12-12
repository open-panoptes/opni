package alerting

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

func (p *Plugin) loadCerts() *tls.Config {
	certs := p.clusterDriver.Get().GetAlertingServiceConfig().Certs

	p.logger.With(
		"alertingServerCa", certs.ServerCA,
		"alertingClientCa", certs.ClientCA,
		"alertingClientCert", certs.ClientCert,
		"alertingClientKey", certs.ClientKey,
	).Debug("loading certs")

	clientCert, err := tls.LoadX509KeyPair(certs.ClientCert, certs.ClientKey)
	if err != nil {
		p.logger.Error(fmt.Sprintf("failed to load alerting client key id : %s", err))
		os.Exit(1)
	}

	serverCaPool := x509.NewCertPool()
	serverCaData, err := os.ReadFile(certs.ServerCA)
	if err != nil {
		p.logger.Error(fmt.Sprintf("failed to read alerting server CA %s", err))
		os.Exit(1)
	}

	if ok := serverCaPool.AppendCertsFromPEM(serverCaData); !ok {
		p.logger.Error(fmt.Sprintf("failed to load alerting server CA %s", err))
		os.Exit(1)
	}

	clientCaPool := x509.NewCertPool()
	clientCaData, err := os.ReadFile(certs.ClientCA)
	if err != nil {
		p.logger.Error(fmt.Sprintf("failed to load alerting client CA : %s", err))
		os.Exit(1)
	}

	if ok := clientCaPool.AppendCertsFromPEM(clientCaData); !ok {
		p.logger.Error("failed to load alerting client Ca")
		os.Exit(1)
	}

	return &tls.Config{
		MinVersion:   tls.VersionTLS12,
		Certificates: []tls.Certificate{clientCert},
		ClientCAs:    clientCaPool,
		RootCAs:      serverCaPool,
	}
}
