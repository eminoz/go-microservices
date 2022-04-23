package config

var Config *Configuration

type Configuration struct {
	MongoDb string `mapstructure:"MONGODB_URI"`
	Port    string `mapstructure:"PORT"`
}

func SetupConfig() (err error) {
	//	config := godotenv.Load("./pkg/config/.env")

	/*if config != nil {
		return nil
	}*/
	configuration := &Configuration{
		MongoDb: "mongodb+srv://eminoz:Eminemin.07@cluster0.cvbx9.mongodb.net/myFirstDatabase?retryWrites=true&w=majority",
		Port:    "3000",
	}
	Config = configuration
	return
}

func GetConfig() *Configuration {
	return Config
}
