package libs

import (
	"os"

	"github.com/spf13/viper"
)
func init(){
	logger:= GetLogger()
	workingdir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(workingdir + "/internal")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		logger.Fatal().Msg("ERROR WHILE LOADING CONFIG " + err.Error())
	}
}