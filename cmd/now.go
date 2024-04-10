/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"sapid/internal"
	"time"

	"github.com/spf13/cobra"
)

// nowCmd represents the now command
var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "오늘의 급식 정보를 불러올 수 있습니다.",
	Long:  `empty`,
	Run: func(cmd *cobra.Command, args []string) {
		meals, err := internal.GetMeals(time.Now())
		if err != nil {
			fmt.Println("급식 정보를 불러올 수 없습니다 : ", err)
			return

		}

		fmt.Printf("아침: %s\n점심: %s\n저녁: %s", meals[0], meals[1], meals[2])
		// fmt.Println(meals)
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nowCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nowCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
