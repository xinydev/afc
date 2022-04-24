package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var dangerCmd string
var noticeCmd string
var baseCmd string

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run user command with confirmation",
	//DisableFlagParsing: true,
	Aliases: []string{"r"},
	RunE: func(cmd *cobra.Command, args []string) error {
		if isNeedConfirm(args) {
			fmt.Printf("%sConfirm???", notice())
			buf := bufio.NewReader(os.Stdin)
			b, err := buf.ReadBytes('\n')
			if err != nil {
				return err
			}
			if strings.ToLower(strings.Trim(string(b), "\n")) != "y" {
				fmt.Println("exit...")
				return nil
			}
		}

		kCmd := exec.Command(baseCmd, args...)
		kCmd.Stdout = os.Stdout
		kCmd.Stderr = os.Stderr
		kCmd.Stdin = os.Stdin
		return kCmd.Run()
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		newArgs := []string{"__complete"}
		newArgs = append(newArgs, args...)
		newArgs = append(newArgs, toComplete)
		kCmd := exec.Command(baseCmd, newArgs...)
		outBytes, _ := kCmd.CombinedOutput()
		validArgs := strings.Split(string(outBytes), "\n")
		if len(validArgs) >= 3 {
			validArgs = validArgs[:len(validArgs)-3]
		}
		return validArgs, cobra.ShellCompDirectiveNoFileComp
	},
}

func isNeedConfirm(args []string) bool {
	arg := strings.Join(args, " ")
	for _, danger := range strings.Split(dangerCmd, ",") {
		if strings.Contains(arg, danger) {
			return true
		}
	}
	return false
}
func notice() string {
	args := strings.Split(noticeCmd, " ")
	if len(args) > 0 {
		cmd := exec.Command(args[0], args[1:len(args)]...)
		out, _ := cmd.CombinedOutput()
		return string(out)
	}
	return ""
}

func init() {
	rootCmd.PersistentFlags().StringVar(&baseCmd, "cmd", "kubectl", "base command")
	rootCmd.PersistentFlags().StringVar(&dangerCmd, "danger", "apply,delete", "dangers args,split by ,")
	rootCmd.PersistentFlags().StringVar(&noticeCmd, "notice", "kubectl config get-contexts", "notice cmd")
	rootCmd.AddCommand(runCmd)
}
