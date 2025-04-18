package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/alex-popov-tech/advent_of_code_2024_go/day_1"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_2"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_3"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_4"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_5"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_6"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_7"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_8"
	"github.com/alex-popov-tech/advent_of_code_2024_go/day_9"
)

var (
	day  *int
	part *int
)

type task func()

var tasks = [][]task{
	{},
	{
		func() {},
		func() { day_1.Part1("./inputs/day_1.txt") },
		func() { day_1.Part2("./inputs/day_1.txt") },
	},
	{
		func() {},
		func() { day_2.Part1("./inputs/day_2.txt") },
		func() { day_2.Part2("./inputs/day_2.txt") },
	},
	{
		func() {},
		func() { day_3.Part1("./inputs/day_3.txt") },
		func() { day_3.Part2("./inputs/day_3.txt") },
	},
	{
		func() {},
		func() { day_4.Part1("./inputs/day_4.txt") },
		func() { day_4.Part2("./inputs/day_4.txt") },
	},
	{
		func() {},
		func() { day_5.Part1("./inputs/day_5.txt") },
		func() { day_5.Part2("./inputs/day_5.txt") },
	},
	{
		func() {},
		func() { day_6.Part1("./inputs/day_6.txt") },
		func() { day_6.Part2("./inputs/day_6.txt") },
	},
	{
		func() {},
		func() { day_7.Part1("./inputs/day_7.txt") },
		func() { day_7.Part2("./inputs/day_7.txt") },
	},
	{
		func() {},
		func() { day_8.Part1("./inputs/day_8.txt") },
		func() { day_8.Part2("./inputs/day_8.txt") },
	},
	{
		func() {},
		func() { day_9.Part1("./inputs/day_9.txt") },
		func() { day_9.Part2("./inputs/day_9.txt") },
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
		if (*day) >= len(tasks) {
			log.Fatalf("There is no day '%d' assignment", *day)
		}
		if (*part) >= len(tasks[*day]) {
			log.Fatalf("There is no day '%d' part '%d' assignment", *day, *part)
		}
		tasks[*day][*part]()
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
