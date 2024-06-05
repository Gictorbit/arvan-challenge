package envconfig

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"os"
)

type WalletEnvConfig struct {
	WalletDatabase string `env:"WALLET_DB_POSTGRES,notEmpty"`
	DebugMode      bool   `env:"DEBUG_MODE" envDefault:"false"`
	Host           string `env:"WALLET_HOST" envDefault:"0.0.0.0"`
	Port           uint   `env:"WALLET_PORT" envDefault:"3000"`
	Address        string `env:"WALLET_ADDRESS,expand" envDefault:"$HOST:$PORT"`
	LogRequests    bool   `env:"LOG_REQUESTS" envDefault:"false"`
}

func ReadWalletEnvironment() (*WalletEnvConfig, error) {
	envFilePath := "wallet.env"
	if CheckFileExists(envFilePath) {
		if err := godotenv.Load(envFilePath); err != nil {
			return nil, err
		}
	}
	cfg := &WalletEnvConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type DiscountEnvConfig struct {
	DiscountDatabase string `env:"DISCOUNT_DB_POSTGRES,notEmpty"`
	DebugMode        bool   `env:"DEBUG_MODE" envDefault:"false"`
	Host             string `env:"DISCOUNT_HOST" envDefault:"0.0.0.0"`
	Port             uint   `env:"DISCOUNT_PORT" envDefault:"3000"`
	Address          string `env:"DISCOUNT_ADDRESS,expand" envDefault:"$HOST:$PORT"`
	LogRequests      bool   `env:"LOG_REQUESTS" envDefault:"false"`
}

func ReadDiscountEnvironment() (*DiscountEnvConfig, error) {
	envFilePath := "discount.env"
	if CheckFileExists(envFilePath) {
		if err := godotenv.Load(envFilePath); err != nil {
			return nil, err
		}
	}
	cfg := &DiscountEnvConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

type GatewayEnvConfig struct {
	DebugMode        bool   `env:"DEBUG_MODE" envDefault:"false"`
	Host             string `env:"GATEWAY_HOST" envDefault:"0.0.0.0"`
	Port             uint   `env:"GATEWAY_PORT" envDefault:"3000"`
	Address          string `env:"GATEWAY_ADDRESS,expand" envDefault:"$HOST:$PORT"`
	LogRequests      bool   `env:"LOG_REQUESTS" envDefault:"false"`
	WalletEndPoint   string `env:"WALLET_ENDPOINT"`
	DiscountEndPoint string `env:"DISCOUNT_ENDPOINT"`
}

func ReadGatewayEnvConfig() (*GatewayEnvConfig, error) {
	envFilePath := "gateway.env"
	if CheckFileExists(envFilePath) {
		if err := godotenv.Load(envFilePath); err != nil {
			return nil, err
		}
	}
	cfg := &GatewayEnvConfig{}
	err := env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func CheckFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
