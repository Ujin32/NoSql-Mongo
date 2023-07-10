package config

type Config struct {
	MongoDBURI     string
	DBName         string
	CollectionName string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		MongoDBURI:     "mongodb://localhost:27017/",
		DBName:         "RebrainTask",
		CollectionName: "users",
	}
	return cfg, nil
}
