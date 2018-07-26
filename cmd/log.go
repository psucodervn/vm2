package cmd

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/psucodervn/vm2/pkg/pm2"
	"github.com/spf13/cobra"
	"log"
	"strings"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "View pm2 log",
	Run:   viewLog,
}

func init() {
	rootCmd.AddCommand(logCmd)
}

type pepper struct {
	Name     string
	HeatUnit int
	Peppers  int
}

func viewLog(cmd *cobra.Command, args []string) {
	procs, _ := pm2.List()

	templates := &promptui.SelectTemplates{
		Label:    "{{ .Name }}?",
		Active:   "> {{ .Name | green }} (id = {{ .PMID | red }})",
		Inactive: "  {{ .Name | green }} (id = {{ .PMID | red }})",
		Selected: "> pm2 log \"{{ .Name | red | green }}\"",
		Details: `--------- Process ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Process id:" | faint }}	{{ .PMID }}
{{ "Exec path:" | faint }}	{{ .Env.ExecPath }}
{{ "Status:" | faint }}	{{ .Env.Status }}`,
	}

	searcher := func(input string, index int) bool {
		proc := procs[index]
		name := strings.Replace(strings.ToLower(proc.Name), " ", "", -1)
		id := strings.Replace(strings.ToLower(fmt.Sprintf("%d", proc.PMID)), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input) || strings.Contains(id, input)
	}

	prompt := promptui.Select{
		Label:     "Choose process to view log:",
		Items:     procs,
		Templates: templates,
		Size:      10,
		Searcher:  searcher,
	}
	prompt.StartInSearchMode = true
	idx, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed: %v\n", err)
		return
	}

	if err := pm2.ViewLog(procs[idx]); err != nil {
		log.Fatalln(err)
	}
}
