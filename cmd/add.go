package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Phantas0s/gocket/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add URL...",
	Short: "Add URLs to Pocket",
	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func add(URLs []string) {
	if len(URLs) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			URLs = strings.Split(scanner.Text(), " ")
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}
	}

	pocket := internal.CreatePocket(consumerKey)
	for _, v := range URLs {
		pocket.Add(v)
	}
}
