package main

import (
	"sort"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/lgnixai/ollamax"
	"github.com/ollama/ollama/format"
)

func listCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "list all models",
		RunE: func(cmd *cobra.Command, args []string) error {
			models, err := ollamax.ListModel()

			if err != nil {
				return err
			}
			sort.Slice(models, func(i, j int) bool {
				return models[i].Size > models[j].Size
			})

			tab := tablewriter.NewWriter(cmd.OutOrStdout())
			tab.SetHeader([]string{"Model", "Size"})
			for _, m := range models {
				tab.Append([]string{m.ShortName, format.HumanBytes(m.Size)})
				//cmd.Println(m.Name, format.HumanBytes(m.Size))
			}
			tab.Render()
			return nil
		},
	}
}
