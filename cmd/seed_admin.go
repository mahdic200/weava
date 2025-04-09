package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models"
	"github.com/mahdic200/weava/Utils"
	"github.com/mahdic200/weava/Utils/ProgressBars/ProgressBar"
	"github.com/spf13/cobra"
)

var number uint64

// seedAdminCmd represents the seed command
var seedAdminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Seeds the admins table",
	Long:  `Seeds the admins table`,
	Run: func(cmd *cobra.Command, args []string) {
		bar := ProgressBar.Default("Seeding [green]Database[reset] :", int(number))

		tx := Config.DB
		pass, _ := Utils.GenerateHashPassword("password")
		if number == 0 {
			return
		}
		for i := 1; i <= int(number); i++ {
			now := time.Now()
			data := Models.Admin{
				First_name: "admin",
				Email:      fmt.Sprintf("admin_%d@gmail.com", i),
				Phone:      fmt.Sprintf("0911%07d", i),
				Password:   pass,
				Created_at: &now,
			}
			if err := tx.Select("First_name", "Email", "Phone", "Password", "Created_at").Create(&data).Error; err != nil {
				fmt.Printf("\nCould not seed the database : %s\n", err.Error())
				bar.Exit()
				os.Exit(2)
			}
			bar.Add(1)
		}
	},
}

func init() {
	seedCmd.AddCommand(seedAdminCmd)
	seedAdminCmd.Flags().Uint64VarP(&number, "number", "n", 0, "When it's set, database will be filled by fake users")
}
