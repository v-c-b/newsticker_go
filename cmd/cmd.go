package cmd
import (
  "fmt"
  "github.com/spf13/cobra"
  "os"
)

var rootCmd = &cobra.Command{
  Use: "newsticker_go",
  Short: "Read Atom feeds",
  Long: `Just a small CLI application.
  Read Atom feeds`,
}

func Exec() {
  rootCmd.AddCommand(cmdLS)
  rootCmd.AddCommand(cmdDescribe)
  err := rootCmd.Execute()
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

type News struct {
  Id int
  Title string
  Text string
  Author string
}

var cmdLS = &cobra.Command{
  Use: "ls",
  Short: "List news",
  Long: `List first 5 news`,
  Run: func(cmd *cobra.Command, args []string) {
    list()
  },
}

var cmdDescribe = &cobra.Command{
  Use: "describe [id]",
  Short: "Show details for an article",
  Long: `Details for an artice`,
  Args: cobra.MinimumNArgs(1),
  Run: func(cmd *cobra.Command, args []string) { describe(args[0])
  },
}
