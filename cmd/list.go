package cmd

import (
	"fmt"
	"log"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/tri/todo"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List will show current todo list",
	Long:  `list command will list all todo items of the list`,
	Run:   listRun,
}

var (
	doneOpt bool
	allOpt  bool
)

func listRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}

	// fmt.Println(items)
	// sort by priority of the todo list.
	sort.Sort(todo.ByPri(items)) //here:type conversion

	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, i := range items {
		if i.Done == doneOpt || allOpt {
			fmt.Fprintln(w, i.Lable()+"\t"+i.PrettyDone()+"\t"+i.Ipretty()+"\t"+i.Text+"\t")
		}
		//
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().BoolVar(&doneOpt, "done", false, "Show 'Done' Todos")
	listCmd.Flags().BoolVarP(&allOpt, "all", "a", false, "Show 'All' Todos")
}
