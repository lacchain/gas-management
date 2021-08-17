package model

type ApplicationConfig struct {
	NodeURL string `mapstructure:"nodeURL"`
	WSURL string `mapstructure:"wsURL"`
	ContractAddress string `mapstructure:"contractAddress"`
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

type SecurityConfig struct {
	PermissionsEnabled bool `mapstructure:"permissionsEnabled"`
	AccountContractAddress string `mapstructure:"accountContractAddress"`
}

type Config struct {
	Application  ApplicationConfig `mapstructure:"application"`
	KeyStore  KeyStoreConfig `mapstructure:"keystore"`
	Passphrase PassphraseConfig   `mapstructure:"passphrase"`
	Security SecurityConfig `mapstructure:"security"`
}