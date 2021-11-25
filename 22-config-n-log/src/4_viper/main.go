package main

import (
	"github.com/spf13/viper"
	"log"
)

//PORT=9000 go run ./src/4_viper
func main() {
	viper.SetDefault("port", 8080)

	/* 1 */
	// viper.BindEnv("port", "PORT")
	// fmt.Println(viper.Get("port"))

	/* 2 */
	// viper.SetEnvPrefix("MY_APP")
	// viper.AutomaticEnv()
	// port := viper.GetInt("port") // Becomes "MY_APP_PORT"
	// fmt.Println(port)

	/* 3 */
	 viper.SetConfigFile("./config/config.json")
	 if err := viper.ReadInConfig(); err != nil {
	 	log.Fatalf("failed to read config: %v", err)
	 }
	// fmt.Println(viper.ConfigFileUsed())
	// fmt.Println(viper.AllSettings())
	//
	// if err := viper.WriteConfigAs("config.json"); err != nil {
	// 	log.Fatalf("failed to write config: %v", err)
	// }

	// /* 4 */
	// replacer := strings.NewReplacer(".", "_")
	// viper.SetEnvKeyReplacer(replacer)
	// viper.AutomaticEnv()
	// port := viper.GetString("amqp.url") // Becomes AMQP_URL
	// fmt.Println(port)

	// /* 5 */
	// type config struct {
	//	Port        int
	//	ServiceName string
	//	Directory   string
	// }
	// var C config
	// err := viper.Unmarshal(&C)
	// if err != nil {
	// log.Fatalf("unable to decode into struct, %v", err)
	// }
	// fmt.Printf("%+v", C)

	// /* 6 */
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGINT)
	// viper.WatchConfig()
	// go func(){
	//	 for {
	//		 time.Sleep(time.Second)
	//		 viper.OnConfigChange(func(e fsnotify.Event) {
	//			 fmt.Println("Config file changed:", e.Name)
	//			 fmt.Println(viper.AllSettings())
	//		 })
	//	 }
	//}()
	//<-ch
}
