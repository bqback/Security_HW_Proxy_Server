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
	"sync"
)

func GenCert(host string, config *config.TLSConfig, logger logging.ILogger) (tls.Certificate, error) {
	certFileName := host + ".crt"
	certFilePath := filepath.Join(config.CertDir, certFileName)
	keyFileName := host + ".key"
	keyFilePath := filepath.Join(config.KeyDir, keyFileName)
	certificate := tls.Certificate{}

	mu := &sync.RWMutex{}

	if _, err := os.Stat(certFilePath); os.IsNotExist(err) {
		logger.Debug("Generating certificate")

		mu.Lock()

		serial, err := genSerial()
		if err != nil {
			logger.Error("Failed generating a serial for the certificate")
			return certificate, err
		}

		genCmd := exec.Command("/bin/sh", config.CertGenScript,
			config.CACertFile, config.CAKeyFile, config.X509Config,
			certFilePath, keyFilePath,
			host, fmt.Sprint(serial))
		logger.Debug(fmt.Sprintf("Command to run: %v", genCmd))

		if err := genCmd.Run(); err != nil {
			return certificate, err
		}

		mu.Unlock()
	}

	certF, err := os.ReadFile(certFilePath)
	if err != nil {
		logger.Error("Failed reading host certificate")
		return certificate, err
	}

	keyF, err := os.ReadFile(keyFilePath)
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

func genSerial() (uint64, error) {
	serial := make([]byte, 64)
	_, err := rand.Read(serial)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint64(serial), nil
}
