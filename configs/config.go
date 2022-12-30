package configs

import "github.com/spf13/viper"

// cria uma variável com um pointer para uma alocação na memória
var cfg *config

// cria um struct (tipo uma classe no Dart) que contém toda a configuração de comunicação com o DB
type config struct {
	// configuração da API
	API APIConfig
	// configuração do DB
	DB  DBConfig
}

// cria um struct com a configuração de acesso à API, que é basicamente a porta
type APIConfig struct {
	Port string
}

// cria o struct com a configuração de acesso ao DB
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

// função inicializadora do GO. Pode ser comparada ao initState do Flutter, ou seja, vai ser executado primeiro
func init() {
	// utiliza o viper para gravar informações padrão de acesso à API, caso não consiga ela pelo arquivo de configuração
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432")
}

// cria uma função de Load que buscará as configurações do DB e da API em um arquivo de configuração
func Load() error {
	// nome do arquivo de configuração
	viper.SetConfigName("config")
	// tipo do arquivo de configuração (parecido com o yaml)
	viper.SetConfigType("toml")
	// o arquivo de configuração sempre estará na pasta raiz que o binário estiver. É isso que o ponto significa
	viper.AddConfigPath(".")
	// cria uma variável err que irá ler o arquivo de configuração e salvar as informações em si mesmo
	err := viper.ReadInConfig()
	// verifica se a variável está vazia. Se estiver, verifica se o erro é que o arquivo de configuração não foi encontrado e retorna err
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	// inicializa a variável cfg criada lá em cima, alocando espaço na memória
	cfg = new(config)

	// grava na variável cfg, no struct API.port a porta que está definida no arquivo de configuração, recuperado pelo viper
	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}
	// grava na variável cfg, no struct DB, todas as informações necessárias para acesso ao banco de dados, recuperado pelo viper
	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	// retorna nil (tipo null no Dart, considerando que é um tipo) para dizer que não há erros e que tudo correu bem
	return nil
}

// cria função para retornar as configurações de acesso do banco de dados
func GetDB() DBConfig {
	return cfg.DB
}

// cria função para retornar configuração da porta para acesso à API
func GetServerPort() string {
	return cfg.API.Port
}