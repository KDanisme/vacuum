// Copyright 2022 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package cmd

import (
	languageserver "github.com/daveshanley/vacuum/language-server"
	"github.com/spf13/cobra"
)

func GetLanguageServerCommand() *cobra.Command {
	return &cobra.Command{
		SilenceErrors: true,
		SilenceUsage:  true,
		Use:           "language-server",
		Short:         "Language server",
		Long:          `Language server for linting your spec in real time`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return languageserver.NewServer(Version).Run()
		},
	}
}
