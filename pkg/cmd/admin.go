// +build optional admin

package cmd

import (
	"github.com/geoffmore/esctl/pkg/admin"
	//"github.com/geoffmore/esctl/pkg/escmd"
	"errors"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func init() {
	rootCmd.AddCommand(adminCmd)
	//adminCmd.AddCommand(adminTest)
	//adminCmd.AddCommand(adminInt)
	// Help gets its own subcommand because it probably makes sense
	adminCmd.AddCommand(listNodes)
	adminCmd.AddCommand(listNodesStorage)
	adminCmd.AddCommand(showBigShards)
	adminCmd.AddCommand(showSmallShards)
	//adminCmd.AddCommand(helpCmd)
	//helpCmd.AddCommand(helpCat)
	//helpCat.AddCommand(helpCatIndices)
}

// I should add a --help flag that adds the pointer help field to a request. Not
// sure how to wrap that for the admin package. Maybe I don't need to?
var adminCmd = &cobra.Command{
	// esctl admin
	Use:   "admin",
	Short: "No description",
}

var listNodes = &cobra.Command{
	// esctl admin list-nodes
	Use:   "list-nodes",
	Short: "No description",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = admin.ListNodes(client, outputFmt, help)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var listNodesStorage = &cobra.Command{
	// esctl admin list-nodes
	Use:   "list-nodes-storage",
	Short: "No description",
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		err = admin.ListNodesStorage(client, outputFmt, help)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var showBigShards = &cobra.Command{
	// esctl admin list-nodes
	Use:   "show-big-shards",
	Short: "No description",
	// first arg int
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an integer argument")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("provided argument was not a an integer")
		}
		if i < 1 {
			return errors.New("provided integer needs to be greater than 0")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		// Error checking is done in Args, so there is no need to duplicate logic
		// in Run
		i, _ := strconv.Atoi(args[0])
		err = admin.ShowBigShards(client, outputFmt, help, i)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var showSmallShards = &cobra.Command{
	// esctl admin list-nodes
	Use:   "show-small-shards",
	Short: "No description",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires an integer argument")
		}
		i, err := strconv.Atoi(args[0])
		if err != nil {
			return errors.New("provided argument was not a an integer")
		}
		if i < 1 {
			return errors.New("provided integer needs to be greater than 0")
		}
		return nil

	},
	Run: func(cmd *cobra.Command, args []string) {
		// Boilerplate
		client, err := genClient(context)
		if err != nil {
			log.Fatal(err)
		}

		// Error checking is done in Args, so there is no need to duplicate logic
		// in Run
		i, _ := strconv.Atoi(args[0])
		err = admin.ShowSmallShards(client, outputFmt, help, i)
		if err != nil {
			log.Fatal(err)
		}
	},
}
