package config

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App struct {
		Env string `yaml:"env" env-default:"dev"`
	} `yaml:"app" env-required:"true"`
	Http struct {
		HttpsEnabled   bool     `yaml:"https_enabled" env-default:"false"`
		Port           uint16   `yaml:"port" env-default:"8080"`
		Host           string   `yaml:"host"`
		KeyFilePath    string   `yaml:"key_file_path"`
		CertFilePath   string   `yaml:"cert_file_path"`
		AllowedOrigins []string `yaml:"allowed_origins" env-default:"http://localhost:*,http://127.0.0.1:*"`
	} `yaml:"http"`
	Database struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Database string `yaml:"database"`
	} `yaml:"database" env-required:"true"`
}

var (
	cfg  Config
	once sync.Once

	k     = "6368616e676520746869732070617373776f726420746f206120736563726574"
	Block cipher.Block // Block is needed to encode/decode sensitive data.
)

func Read(configPath string) {
	once.Do(func() {
		err := cleanenv.ReadConfig(configPath, &cfg)
		if err != nil {
			log.Fatalf("Failed to read config: %v", err)
		}

		decodedKey, err := hex.DecodeString(k)
		if err != nil {
			log.Fatalf("Failed to decode app key: %v", err)
		}

		Block, err = aes.NewCipher(decodedKey)
		if err != nil {
			log.Fatalf("Failed to create cipher block: %v", err)
		}
	})
}

func Get() *Config {
	return &cfg
}
