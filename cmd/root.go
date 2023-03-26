package cmd

import (
	"github.com/spf13/cobra"
)

var database string

var rootCmd = &cobra.Command{
	Use:   "ezgem",
	Short: "ezgem is a quick and simple gemini server to quickly get you started on geminispace",
	Long:  `A simple and sometimes opiniated gemini server to get you onto geminispace in the shortest time possible`,
}

func init() {
	//	rootCmd.PersistentFlags().StringVarP(&database, "database", "d", "g2g.db", "filename of sqlite3 database to use")

}
func Execute() error {
	return rootCmd.Execute()
}
