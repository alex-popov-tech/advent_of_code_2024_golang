package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/alex-popov-tech/advent_of_code_2024_go/day_1"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_2"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_3"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_4"
)

var (
	day  *int
	part *int
)

type task func()

var tasks = [][]task{
	{
		func() { day_1.Part1("./inputs/day_1.txt") },
		func() { day_1.Part2("./inputs/day_1.txt") },
	},
	{
		func() { day_2.Part1("./inputs/day_2.txt") },
		func() { day_2.Part2("./inputs/day_2.txt") },
	},
	{
		func() { day_3.Part1("./inputs/day_3.txt") },
		func() { day_3.Part2("./inputs/day_3.txt") },
	},
	{
		func() { day_4.Part1("./inputs/day_4.txt") },
		// func() { day_4.Part2("./inputs/day_4.txt") },
	},
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "Advent of Code 2024 in golang",
	Long:  "Advent of Code 2024 in golang",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if *day < 1 || *day > 25 {
			return fmt.Errorf("day must be between 1 and 25, got %d", day)
		}

		if *part < 1 || *part > 2 {
			return fmt.Errorf("part must be either 1 or 2, got %d", part)
		}

		return nil
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		tasks[*day-1][*part-1]()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.advent_of_code_2024_go.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	day = rootCmd.Flags().IntP("day", "d", 0, "# of day to run")
	rootCmd.MarkFlagRequired("day")
	part = rootCmd.Flags().IntP("part", "p", 0, "# of part to run")
	rootCmd.MarkFlagRequired("part")
}
