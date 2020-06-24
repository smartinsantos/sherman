package config

type (
	// AppConfig type definition
	AppConfig struct {
		Env   string
		Debug bool
		Port  int
		Addr  string
	}
	// DBConfig type definition
	DBConfig struct {
		Driver      string
		Name        string
		User        string
		Pass        string
		Host        string
		Port        string
		ExposedPort string
	}
	// JwtConfig type definition
	JwtConfig struct {
		Secret string
	}
	// GlobalConfig type definition
	GlobalConfig struct {
		App AppConfig
		DB  DBConfig
		Jwt JwtConfig
	}
	// Config type definition
	Config interface {
		Load(cfg GlobalConfig)
		Get() GlobalConfig
	}

	service struct {
		Config GlobalConfig
	}
)

var (
	// DefaultConfig contains default values of global config
	DefaultConfig = &GlobalConfig{
		App: AppConfig{
			Env:   "local",
			Debug: false,
			Port:  5000,
			Addr:  ":5000",
		},
		DB: DBConfig{
			Driver:      "mysql",
			Name:        "sherman",
			User:        "db_user",
			Pass:        "db_password",
			Host:        "app-mysql",
			Port:        "3306",
			ExposedPort: "5001",
		},
		Jwt: JwtConfig{
			Secret: "jwt_secret",
		},
	}
	// TestConfig contains values of global config for testing
	TestConfig = GlobalConfig{
		App: AppConfig{
			Env:   "test",
			Debug: false,
			Port:  5000,
			Addr:  ":5000",
		},
		DB: DBConfig{
			Driver:      "mysql",
			Name:        "sherman",
			User:        "db_user",
			Pass:        "db_password",
			Host:        "localhost",
			Port:        "5001",
			ExposedPort: "5001",
		},
		Jwt: JwtConfig{
			Secret: "jwt_secret",
		},
	}
)

// New returns an instance of config.Config
func New() Config {
	return &service{
		Config: loadFromEnv(),
	}
}

// Load loads a provided GlobalConfig
func (s *service) Load(cfg GlobalConfig) {
	s.Config = cfg
}

// Get returns the loaded GlobalConfig instance
func (s *service) Get() GlobalConfig {
	return s.Config
}
