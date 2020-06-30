package model

type ApplicationConfig struct {
	NodeURL string `mapstructure:"nodeURL"`
	ContractAddress string `mapstructure:"contractAddress"`
	NodeKeyPath string `mapstructure:"nodeKeyPath"`
	NodeAddressPath string `mapstructure:"nodeAddressPath"`
	Key string `mapstructure:"key"`
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