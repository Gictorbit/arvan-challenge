package envconfig

import (
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"os"
)

type WalletEnvConfig struct {
	WalletDatabase string `env:"WALLET_DB_POSTGRES,notEmpty"`
	DebugMode      bool   `env:"DEBUG_MODE" envDefault:"false"`
	Host           string `env:"HOST" envDefault:"0.0.0.0"`
	Port           uint   `env:"PORT" envDefault:"3000"`
	Address        string `env:"ADDRESS,expand" envDefault:"$HOST:$PORT"`
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
	WalletDatabase string `env:"DISCOUNT_DB_POSTGRES,notEmpty"`
	DebugMode      bool   `env:"DEBUG_MODE" envDefault:"false"`
	Host           string `env:"HOST" envDefault:"0.0.0.0"`
	Port           uint   `env:"PORT" envDefault:"3000"`
	Address        string `env:"ADDRESS,expand" envDefault:"$HOST:$PORT"`
	LogRequests    bool   `env:"LOG_REQUESTS" envDefault:"false"`
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
	Host             string `env:"HOST" envDefault:"0.0.0.0"`
	Port             uint   `env:"PORT" envDefault:"3000"`
	Address          string `env:"ADDRESS,expand" envDefault:"$HOST:$PORT"`
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
