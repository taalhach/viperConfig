package main

import (
	"github.com/spf13/viper"
	"log"
	"fmt"
)
func main() {
	viper.SetConfigFile("./viperConfig/config/websockets.json")
	// or use names and paths
	//viper.AddConfigPath("./viperConfig/config")
	//viper.SetConfigName("websockets")
	//viper.SetConfigType("json")
	if err:=viper.ReadInConfig(); err!=nil{
		log.Println("Error in reading file: ",err)
		return
	}
	fmt.Println("is file used: ",viper.ConfigFileUsed())
	server:=viper.Get("bootstrapServer")
	fmt.Println(server)
}
