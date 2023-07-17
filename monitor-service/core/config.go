package core

var Config = struct {
	Environment string `required:"true" env:"ENVIRONMENT"`

	APP struct {
		BindPort    int    `required:"true" env:"APP_BIND_PORT"`
		RpcUrl      string `required:"true" env:"RPC_URL"`
		MongoDbUrl  string `required:"true" env:"MONGO_DB_URL"`
		MongoDbName string `required:"true" env:"MONGO_DB_NAME"`
	}
}{}
