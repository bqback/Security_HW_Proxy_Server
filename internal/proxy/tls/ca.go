package proxytls

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"proxy_server/internal/config"
	"proxy_server/internal/logging"
)

func LoadCA(config *config.TLSConfig, logger *logging.LogrusLogger) (tls.Certificate, error) {
	cert, err := tls.LoadX509KeyPair(config.CACertFile, config.CAKeyFile)
	if os.IsNotExist(err) {
		logger.Error("CA certificate and/or key don't exist")
		return tls.Certificate{}, err
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])

	return cert, err
}
