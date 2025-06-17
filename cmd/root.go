package cmd

import (
    "github.com/spf13/cobra"
    "fmt"
)

var rootCmd = &cobra.Command{
    Use:   "inventaris",
    Short: "Aplikasi Inventaris CLI",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}