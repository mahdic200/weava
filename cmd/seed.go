package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models/User"
	"github.com/mahdic200/weava/Utils"
	"github.com/mahdic200/weava/Utils/ProgressBars/ProgressBar"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeds the database",
	Long:  `Seeds the database`,
	Run: func(cmd *cobra.Command, args []string) {
		bar := ProgressBar.Default("Seeding [green]Database[reset] :", 100_000)

		tx := Config.DB
		pass, _ := Utils.GenerateHashPassword("password")

		for i := 1; i <= 100_000; i++ {
			data := map[string]string{
				"first_name": "admin",
				"email":      fmt.Sprintf("user_%d@gmail.com", i),
				"phone":      fmt.Sprintf("0911%07d", i),
				"password":   pass,
				"created_at": time.Now().String(),
			}
			if err := User.Create(tx, data).Error; err != nil {
				fmt.Printf("\nCould not seed the database : %s\n", err.Error())
				bar.Exit()
				os.Exit(2)
			}
			bar.Add(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)
	// seedCmd.Flags().IntVarP()
}
