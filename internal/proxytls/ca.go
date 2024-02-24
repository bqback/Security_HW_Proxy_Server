package proxytls

import (
	"crypto/tls"
	"crypto/x509"
	"os"
	"path/filepath"
	"proxy_server/internal/config"
	"proxy_server/internal/logging"
)

func LoadCA(config *config.TLSConfig, logger *logging.LogrusLogger) (tls.Certificate, error) {
	keyFile := filepath.Join(config.TLSDir, config.KeyFile)
	certFile := filepath.Join(config.TLSDir, config.CertFile)

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if os.IsNotExist(err) {
		logger.Error("CA certificate and/or key don't exist")
		return tls.Certificate{}, err
	}
	cert.Leaf, err = x509.ParseCertificate(cert.Certificate[0])

	return cert, err
}
