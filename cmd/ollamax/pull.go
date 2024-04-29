package main

import (
	"github.com/lgnixai/ollamax"
	"github.com/ollama/ollama/api"
	"github.com/spf13/cobra"
)

func pullCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "pull",
		Short: "pull model",
		RunE: func(cmd *cobra.Command, args []string) error {
			bar := progressbar.Default(100)
			err := ollamax.PullModel(cmd.Context(), args[0], func(r api.ProgressResponse) {
				if r.Total == 0 {
					return
				}
				bar.Set(int(r.Completed * 100 / r.Total))
			})
			bar.Finish()
			return err
		},
	}
}
