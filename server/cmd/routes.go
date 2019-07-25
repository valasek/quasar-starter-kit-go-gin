// Copyright © 2018-2019 Stanislav Valasek <valasek@gmail.com>

package cmd

import (
	"github.com/valasek/quasar-starter-kit-go-gin/server/api"
	"github.com/valasek/quasar-starter-kit-go-gin/server/routes"

	"github.com/spf13/cobra"
)

// routesCmd represents the routes command
var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Prints the list of all available routes",
	Long: `Prints the list of all available routes.
	
Command will print routes.`,
	Run: func(cmd *cobra.Command, args []string) {
		api := api.NewAPI()
		r := routes.SetupRouter(api)

		routes.PrintRoutes(r)
	},
}

func init() {
	rootCmd.AddCommand(routesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// routesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// routesCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
