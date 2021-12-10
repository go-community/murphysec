package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"murphysec-cli-simple/plugins"
)

var scanDir string

func scanCmd() *cobra.Command {
	c := &cobra.Command{
		Use:              "scan",
		Run:              scanHandler,
		TraverseChildren: true,
	}
	c.PersistentFlags().StringVarP(&scanDir, "dir", "d", ".", "project root dir")
	for _, it := range plugins.Plugins {
		pc := &cobra.Command{
			Use:              it.Info().Name,
			Short:            it.Info().ShortDescription,
			TraverseChildren: true,
			Run: func(cmd *cobra.Command, args []string) {
				it.MatchPath(context.Background(), scanDir)
			},
		}
		it.SetupScanCmd(pc)
		c.AddCommand(pc)
	}
	return c
}

func scanHandler(cmd *cobra.Command, args []string) {
	panic("todo: auto scan") // todo
}