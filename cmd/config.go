// Copyright © 2018 Adron Hall <adron@thrashingcode.com>
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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For the custom example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Twitterers File: %s\n", viper.GetString("file"))
		fmt.Printf("Export File: %s\n", viper.GetString("fileExport"))
		fmt.Printf("Export Format: %s\n", viper.GetString("fileFormat"))
		fmt.Printf("Consumer API Key: %s\n", viper.GetString("consumer_api_key")[0:6])
		fmt.Printf("Consumer API Secret: %s\n", viper.GetString("consumer_api_secret")[0:6])
		fmt.Printf("Access Token: %s\n", viper.GetString("access_token")[0:6])
		fmt.Printf("Access Token Secret: %s\n", viper.GetString("access_token_secret")[0:6])
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
