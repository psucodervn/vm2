package cmd

import (
  "github.com/spf13/cobra"
  "fmt"
  "os"
)

var rootCmd = cobra.Command{
  Use:   "vm2",
  Short: "pm2 utilities",
  Run: func(cmd *cobra.Command, args []string) {

  },
}

func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
