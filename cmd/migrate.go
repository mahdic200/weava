package cmd

import (
	"fmt"
	"os"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Utils/ProgressBars/ProgressBar"
	"github.com/spf13/cobra"
)

var force bool

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Makes the tables",
	Long:  `Migrates the database based on models`,
	Run: func(cmd *cobra.Command, args []string) {
		models := []any {
			Models.Admin{},
			Models.User{},
			Models.AdminSession{},
			Models.Session{},
		}
		total := len(models)
		tx := Config.DB

		if force {
			bar := ProgressBar.Default("Dropping All [green]Tables[reset] :", total)
			for _, model := range models {
				if err := tx.Migrator().DropTable(&model); err != nil {
					fmt.Printf("%s\n", err)
					bar.Exit()
					os.Exit(2)
				}
				bar.Add(1)
			}
		}

		bar := ProgressBar.Default("Migrating [green]Tables[reset] :", total)

		for _, model := range models {
			if err := tx.AutoMigrate(&model); err != nil {
				fmt.Printf("%s\n", err)
				bar.Exit()
				os.Exit(2)
			}
			bar.Add(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().BoolVarP(&force, "force", "f", false, "Clears all tables and then migrates")
}


