// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/5imili/reboot/server"
	"github.com/leopoldxx/go-utils/trace/glog"
	"time"
	"github.com/5imili/reboot/server/controller"
	"github.com/5imili/reboot/cmd/pkgs"
)

var cfgFile string

// This represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
// Uncomment the following line if your bare application
// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		defer glog.Flush()
		ticker := time.NewTimer(time.Second)
		defer ticker.Stop()
		go func(){
			for _ = range ticker.C{
				glog.Flush()
			}
		}()

		glog.Info(viper.GetString("server.listen"))

		s := server.New(server.Options{
			ListenAddr:viper.GetString("server.listen"),
			CtrlOpts:&controller.Options{
				DB:pkgs.GetDao(),
			},
		})

		schedManager := pkgs.GetScheduler(pkgs.GetDao())
		go schedManager.Schedule()
		defer schedManager.Stop()
		glog.Fatal(s.ListenAndServer())
		//Todo: graceful exit
	},
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "","config file (default is $HOME/.reboot.yaml)")
	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	viper.SetConfigFile(cfgFile)
	//if cfgFile != "" { // enable ability to specify config file via flag
	//	viper.SetConfigFile(cfgFile)
	//}
	//
	//viper.SetConfigName(".reboot") // name of config file (without extension)
	//viper.AddConfigPath("$HOME")  // adding home directory as first search path
	//viper.AutomaticEnv()          // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
