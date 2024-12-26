package cmd

import (
	"log"

	"ebookmod/pkg/api"

	"ebookmod/app"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(apiCmd)
}

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root short description",
	Long:  "Root long description",
}

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Root short description",
	Long:  "Root long description",
	Run:   startAPI,
}

// var routerCmd = &cobra.Command{
//  Use:   "router",
//  Short: "Root short description",
//  Long:  "Root long description",
//  Run:   startRoute,
// }

// func startRoute(*cobra.Command, []string) {

// }

func startAPI(*cobra.Command, []string) {
	// fmt.Println("hellooo...")
	r := app.APIRouter()
	api.Start(r)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}
