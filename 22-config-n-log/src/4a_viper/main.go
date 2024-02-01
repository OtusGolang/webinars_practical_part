package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetDefault("Url", "default.com")
	viper.SetDefault("WorkerCount", 4)
	viper.SetDefault("LaunchesCount", 0)

	// viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	// viper.AddConfigPath(".")

	// viper.SetEnvPrefix("MYAPP")
	// viper.AutomaticEnv()

	// pflag.String("url", "fromflag.ru", "url to connect")
	// pflag.Parse()
	// viper.BindPFlags(pflag.CommandLine)

	// err := viper.ReadInConfig()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Printf("Launch #%d\n", viper.GetInt("LaunchesCount"))
	fmt.Println("Url:", viper.Get("Url"))
	fmt.Println("Workers:", viper.Get("WorkerCount"))

	// viper.Set("LaunchesCount", viper.GetInt("LaunchesCount")+1)
	// err = viper.WriteConfig()
	// if err != nil {
	// 	fmt.Println("cant write config: ", err)
	// }

	// viper.Debug()
}

//  MYAPP_URL="fromenv.ru" go run main.go
// go run main.go --url="fromflag.ru"

// To show:
// 1. base idea, default values
// 2. config files
// 3. env variables
// 4. flags
// 5. writes
// 6. debug
// 7. watch
// 8. io.reader `viper.ReadConfig(bytes.NewBuffer(yamlExample))`
