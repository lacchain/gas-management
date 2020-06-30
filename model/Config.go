package model

type ApplicationConfig struct {
	NodeURL string `mapstructure:"nodeURL"`
	ContractAddress string `mapstructure:"contractAddress"`
}

type KeyStoreConfig struct {
	RelayNode string `mapstructure:"feeAgent"`
}

type PassphraseConfig struct {
	RelayNode string `mapstructure:"feeAgent"`
}

type Config struct {
	Application  ApplicationConfig `mapstructure:"application"`
	KeyStore  KeyStoreConfig `mapstructure:"keystore"`
	Passphrase PassphraseConfig   `mapstructure:"passphrase"`
}