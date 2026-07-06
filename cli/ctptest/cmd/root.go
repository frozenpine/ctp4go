/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/frozenpine/ctp4go"
	"github.com/spf13/cobra"
)

var (
	version, goVersion, gitVersion, buildTime string

	rootCtx, rootCancel = context.WithCancel(context.Background())

	ctpArgs = [][2]string{
		{"tdfront", "Ctp trade front"},
		{"mdfront", "Ctp marketdata front"},
		{"user", "Ctp login user"},
		{"pass", "Ctp login pass"},
		{"appid", "Ctp app id"},
		{"authcode", "Ctp auth code"},
	}
)

func getEnvArgs(cmd *cobra.Command) map[string]string {
	prefix, _ := cmd.Flags().GetString("env")
	if !strings.HasSuffix(prefix, "_") {
		prefix = prefix + "_"
	}

	args := make(map[string]string)

	for _, v := range ctpArgs {
		envArg := os.Getenv(fmt.Sprint(prefix, strings.ToUpper(v[0])))
		cmdArg, _ := cmd.Flags().GetString(v[0])
		if cmdArg != "" {
			args[v[0]] = cmdArg
		} else {
			args[v[0]] = envArg
		}
	}

	return args
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "ctptest",
	Short: "CTP look-through authentication",
	Long:  ``,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
		verbose, _ := cmd.Flags().GetCount("verbose")
		logDir, _ := cmd.Flags().GetString("log")
		logKeep, _ := cmd.Flags().GetInt("keep")
		logSize, _ := cmd.Flags().GetInt("size")
		isJson, _ := cmd.Flags().GetBool("json")
		console, _ := cmd.Flags().GetBool("console")

		var logName = cmd.Root().Name()
		if name := cmd.Name(); name != logName {
			logName = strings.Join([]string{logName, name}, ".")
		}

		logOptions := ctp4go.LogOptions{ctp4go.WithLogVerbose(verbose)}
		if logDir != "" {
			logOptions = append(logOptions, ctp4go.WithLogFile(
				filepath.Join(logDir, logName),
				ctp4go.WithLogFileSize(logSize),
				ctp4go.WithLogFileAge(logKeep),
			))
		}
		if isJson {
			logOptions = append(logOptions, ctp4go.WithLogJson())
		}
		if console {
			logOptions = append(logOptions, ctp4go.WithLogConsole())
		}

		logger, err := ctp4go.NewLogger(logName, logOptions...)
		if err != nil {
			return err
		}

		slog.SetDefault(logger)

		initArgs := getEnvArgs(cmd)
		slog.Info("ctp args", slog.Any("args", initArgs))

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	defer rootCancel()

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCtx, rootCancel = signal.NotifyContext(
		context.Background(),
		syscall.SIGKILL, syscall.SIGABRT, syscall.SIGINT,
		syscall.SIGQUIT, syscall.SIGTERM, os.Interrupt,
	)

	rootCmd.Version = fmt.Sprintf(
		"%s, Commit: %s, Build: %s@%s",
		version, gitVersion, buildTime, goVersion,
	)

	for _, cmd := range rootCmd.Commands() {
		cmd.Version = rootCmd.Version
	}

	rootCmd.PersistentFlags().String(
		"log", "./logs", "Log dir",
	)
	rootCmd.PersistentFlags().Int(
		"keep", 7, "Log rotated keep days",
	)
	rootCmd.PersistentFlags().Int(
		"size", 500, "Log output rotate size",
	)
	rootCmd.PersistentFlags().Bool(
		"json", false, "Log in json format",
	)
	rootCmd.PersistentFlags().CountP(
		"verbose", "v", "Verbose level",
	)
	rootCmd.PersistentFlags().Bool(
		"console", false, "Log to console",
	)

	for _, args := range ctpArgs {
		rootCmd.PersistentFlags().String(args[0], "", args[1])
	}
	rootCmd.PersistentFlags().String(
		"env", "CTP_", "Ctp args environment prefix",
	)
}
