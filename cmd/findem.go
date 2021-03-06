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
	"github.com/Adron/twitz/coreTwitz"
	"github.com/Adron/twitz/helpers"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/spf13/cobra"
)

var findemCmd = &cobra.Command{
	Use:   "findem",
	Short: "The 'twitz findem' command goes through the parsed list of Twitter accounts and retrieves some basic data for each account.",
	Long: `The 'twitz findem' command goes through the parsed list of Twitter accounts and retrieves some basic data for each account.

That data includes:
 * Something
 * Another`,
	Run: func(cmd *cobra.Command, args []string) {
		// Step 1: Get list of Twitter accounts to query for.
		completedTwittererList := parseTwittererList()
		// Step 2: Get the Twitter client setup.
		twitterClient := coreTwitz.SetupConnection()
		// Step 3: Setup the parameters for the Twitter query for the Twitter accounts.
		userLookupParams := &twitter.UserLookupParams{ScreenName: completedTwittererList}
		// Step 4: Query the Twitter API for the account information.
		users, _, err := twitterClient.Users.Lookup(userLookupParams)
		helpers.Check(err)
		// Step 5: Print out the results to configured and pertinent outputs.
		var p = coreTwitz.TwitterDerived{TwitterAccounts: users}
		coreTwitz.ProcessTwitterAccounts(p)
		coreTwitz.PrintUsersToConsole(users)
		// Profit. Or ya know, be done with it.
	},
}

func parseTwittererList() []string {
	fmt.Println("Starting Twitter Information Retrieval.")
	completedTwittererList := coreTwitz.BuildTwitterList(true)
	fmt.Printf("Getting Twitter details for: \n%s", completedTwittererList)
	return completedTwittererList
}

func init() {
	rootCmd.AddCommand(findemCmd)
}
