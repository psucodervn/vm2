package cmd

import (
  "github.com/spf13/cobra"
  "github.com/manifoldco/promptui"
  "fmt"
  "github.com/psucodervn/vm2/pkg/pm2"
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
    Selected: "> {{ .Name | red | green }}",
    Details: `--------- Process ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Process id:" | faint }}	{{ .PMID }}
{{ "Exec path:" | faint }}	{{ .Env.ExecPath }}
{{ "Status:" | faint }}	{{ .Env.Status }}`,
  }

  searcher := func(input string, index int) bool {
    pepper := procs[index]
    name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
    id := strings.Replace(strings.ToLower(fmt.Sprintf("%d", pepper.PMID)), " ", "", -1)
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
  _, _, err := prompt.Run()

  if err != nil {
    fmt.Printf("Prompt failed %v\n", err)
    return
  }
}