package proxytls

import (
	"crypto/rand"
	"crypto/tls"
	"encoding/binary"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"proxy_server/internal/config"
	"proxy_server/internal/logging"
)

func GenCert(host string, config *config.TLSConfig, logger logging.ILogger) (tls.Certificate, error) {
	certFileName := host + ".crt"
	certFilePath := filepath.Join(config.CertDir, certFileName)
	certificate := tls.Certificate{}

	if _, err := os.Stat(certFilePath); os.IsNotExist(err) {
		logger.Debug("Generating certificate")
		serial := make([]byte, 64)
		_, err = rand.Read(serial)
		if err != nil {
			logger.Error("Error generating serial for the certificate")
			return certificate, err
		}
		genCmd := exec.Command("/bin/sh", filepath.Join(config.TLSDir, config.CertGenScript),
			config.CertKeyFile, config.CACertFile, config.CAKeyFile,
			certFilePath, host, fmt.Sprint(binary.LittleEndian.Uint64(serial)))
		genCmd.Dir = config.TLSDir
		logger.Debug(fmt.Sprintf("Command to run: %v", genCmd))
		if err := genCmd.Run(); err != nil {
			return certificate, err
		}
	}

	certF, err := os.ReadFile(certFilePath)
	if err != nil {
		logger.Error("Failed reading host certificate")
		return certificate, err
	}

	keyF, err := os.ReadFile(config.CertKeyFile)
	if err != nil {
		logger.Error("Failed reading private key")
		return certificate, err
	}

	certificate, err = tls.X509KeyPair(certF, keyF)
	if err != nil {
		logger.Error("Failed parsing certificate/key pair")
		return certificate, err
	}

	return certificate, nil
}
