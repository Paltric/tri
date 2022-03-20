package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/tri/todo"
	"github.com/spf13/viper"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search item from todo list",
	Run:   searchRun,
}

func searchRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Fatalln("open file with a error:", err)
	}

	for _, i := range items {
		if strings.Contains(i.Text, args[0]) {
			w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
			fmt.Fprintln(w, i.Lable()+"\t"+i.PrettyDone()+"\t"+i.Ipretty()+"\t"+i.Text+"\t")
			w.Flush()
		}
	}
}

func init() {
	rootCmd.AddCommand(searchCmd)
}
