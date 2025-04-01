package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/mahdic200/weava/Config"
	"github.com/mahdic200/weava/Models/User"
	"github.com/mahdic200/weava/Utils"
	"github.com/spf13/cobra"
)

// seedCmd represents the seed command
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seeds the database",
	Long:  `Seeds the database`,
	Run: func(cmd *cobra.Command, args []string) {
		tx := Config.DB
		pass, _ := Utils.GenerateHashPassword("password")
		if err := User.Create(tx, map[string]string{"first_name": "admin", "email": "admin@gmail.com", "phone": "09531532475", "password": pass, "created_at": time.Now().String()}).Error; err != nil {
			fmt.Printf("Could not seed the database : %s\n", err.Error())
			os.Exit(2)
		}
		fmt.Printf("Seeded the database successfully\n")
	},
}

func init() {
	rootCmd.AddCommand(seedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// seedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// seedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
