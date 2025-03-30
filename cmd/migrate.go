package cmd

import (
	"fmt"
	"os"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/spf13/cobra"
)

var force bool

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Makes the tables",
	Long:  `Migrates the database based on models`,
	Run: func(cmd *cobra.Command, args []string) {
		tx := Config.DB
		if force {

			fmt.Printf("Dropping all tables\n")
			if err := tx.Migrator().DropTable(&Models.Session{}, &Models.User{}); err != nil {
				fmt.Printf("%s\n", err.Error())
				os.Exit(2)
			}
		}

		if err := Config.DB.AutoMigrate(
			&Models.User{},
			&Models.Session{},
		); err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(2)
		} else {
			fmt.Printf("Migration completed successfully\n")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().BoolVarP(&force, "force", "f", false, "Clears all tables and then migrates")
}
