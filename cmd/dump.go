package cmd

import (
	"context"
	"encoding/json"
	"os"

	"github.com/skuid/spec"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type dumpCmdArgs struct {
	table   string
	label   spec.SelectorSet
	verbose bool
}

var dumpArgs = &dumpCmdArgs{}

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump the JSON representation of releases",
	Run:   dump,
}

func init() {
	RootCmd.AddCommand(dumpCmd)
	f := dumpCmd.Flags()
	f.VarP(&dumpArgs.label, "label", "l", `The labels to filter by. Each label should have the format "k=v".
    	Can be specified multiple times, or a comma-separated list.`)

	f.BoolVar(&dumpArgs.verbose, "v", false, "Pretty-print the JSON")
}

func dump(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), viper.GetDuration("timeout"))
	defer cancel()
	releases, err := releaseStore.List(ctx, dumpArgs.label.ToMap())
	exitOnErr(err)

	encoder := json.NewEncoder(os.Stdout)
	if dumpArgs.verbose {
		encoder.SetIndent("", "    ")
	}
	err = encoder.Encode(releases)
	exitOnErr(err)
}
