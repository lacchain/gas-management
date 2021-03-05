package model

type ApplicationConfig struct {
	NodeURL string `mapstructure:"nodeURL"`
	ContractAddress string `mapstructure:"contractAddress"`
	AccountContractAddress string `mapstructure:"accountContractAddress"`
	NodeKeyPath string `mapstructure:"nodeKeyPath"`
	NodeAddressPath string `mapstructure:"nodeAddressPath"`
	Key string `mapstructure:"key"`
	Port string `mapstructure:"port"`
}

type KeyStoreConfig struct {
	Agent string `mapstructure:"agent"`
}

type PassphraseConfig struct {
	Agent string `mapstructure:"agent"`
}

type Config struct {
	Application  ApplicationConfig `mapstructure:"application"`
	KeyStore  KeyStoreConfig `mapstructure:"keystore"`
	Passphrase PassphraseConfig   `mapstructure:"passphrase"`
}