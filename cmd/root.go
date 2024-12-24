package cmd

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "genpwd",
	Short: "一个简单的密码生成器",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		length, _ := cmd.Flags().GetInt("length")

		if length < 1 {
			fmt.Println("错误：密码长度必须大于0")
			os.Exit(1)
		}

		upper, _ := cmd.Flags().GetBool("upper")
		lower, _ := cmd.Flags().GetBool("lower")
		number, _ := cmd.Flags().GetBool("number")
		special, _ := cmd.Flags().GetBool("special")

		var chars string
		if upper {
			chars += upperCase
		}
		if lower {
			chars += lowerCase
		}
		if number {
			chars += digits
		}
		if special {
			chars += specialChar
		}

		if chars == "" {
			fmt.Println("错误：未选择任何字符类型")
			os.Exit(1)
		}

		password := make([]byte, length)
		for i := range password {
			password[i] = chars[rand.Intn(len(chars))]
		}

		fmt.Println(string(password))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

const (
	upperCase   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lowerCase   = "abcdefghijklmnopqrstuvwxyz"
	digits      = "0123456789"
	specialChar = "!@#$%^&*()_-+=<>?/{}~"
)

func init() {
	rootCmd.Flags().IntP("length", "n", 16, "密码长度")
	rootCmd.Flags().BoolP("upper", "A", true, "包含大写字母"+upperCase)
	rootCmd.Flags().BoolP("lower", "a", true, "包含小写字母"+lowerCase)
	rootCmd.Flags().BoolP("number", "0", true, "包含数字"+digits)
	rootCmd.Flags().BoolP("special", "!", true, "包含特殊字符"+specialChar)
}
