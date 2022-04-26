package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
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
	Run: func(cmd *cobra.Command, args []string) {
		if isNeedConfirm(args) {
			fmt.Printf("%sConfirm??? y or n", notice())
			buf := bufio.NewReader(os.Stdin)
			b, err := buf.ReadBytes('\n')
			if err != nil {
				return
			}
			if strings.ToLower(strings.Trim(string(b), "\n")) != "y" {
				fmt.Println("exit...")
				return
			}
		}

		kCmd := exec.Command(baseCmd, args...)
		kCmd.Stdout = os.Stdout
		kCmd.Stderr = os.Stderr
		kCmd.Stdin = os.Stdin
		_ = kCmd.Run()
		return
	},
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		newArgs := []string{"__complete"}
		newArgs = append(newArgs, args...)
		newArgs = append(newArgs, toComplete)
		kCmd := exec.Command(baseCmd, newArgs...)
		outBytes, _ := kCmd.CombinedOutput()

		validArgs := strings.Split(string(outBytes), "\n")

		// check directive
		compDirective := cobra.ShellCompDirectiveDefault
		if len(validArgs) >= 2 {
			r, _ := regexp.Compile("Shell\\w+")
			if comp, ok := map[string]cobra.ShellCompDirective{
				"ShellCompDirectiveError":         cobra.ShellCompDirectiveError,
				"ShellCompDirectiveNoSpace":       cobra.ShellCompDirectiveNoSpace,
				"ShellCompDirectiveNoFileComp":    cobra.ShellCompDirectiveNoFileComp,
				"ShellCompDirectiveFilterFileExt": cobra.ShellCompDirectiveFilterFileExt,
				"ShellCompDirectiveFilterDirs":    cobra.ShellCompDirectiveFilterDirs,
			}[r.FindString(validArgs[len(validArgs)-2])]; ok {
				compDirective = comp
			}
		}

		if len(validArgs) >= 3 {
			validArgs = validArgs[:len(validArgs)-3]
		}
		return validArgs, compDirective
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
