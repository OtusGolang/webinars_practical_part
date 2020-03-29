package main

import (
	"fmt"
	"github.com/spf13/viper"
)

//EXAMPLE_PORT=9000 go run viper/main.go
func main() {
	///* 1 */
	//viper.SetDefault("port", 8080)
	////viper.BindEnv("port", "EXAMPLE_PORT")
	////fmt.Println(viper.Get("port"))
	///*2*/
	//viper.SetEnvPrefix("foo")     // Becomes "FOO_"
	//os.Setenv("FO_PORT", "1313") // typically done outside of the app
	//viper.AutomaticEnv()
	//port := viper.GetInt("port")
	//fmt.Println(port)
	/* 3 */
	viper.SetConfigFile("/Users/a.zheltak/GolandProjects/awesomeProject/config/config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	fmt.Println(viper.ConfigFileUsed())
	fmt.Println(viper.AllSettings())
	if err := viper.WriteConfigAs("config.yml"); err != nil {
		panic(err)
	}
}
