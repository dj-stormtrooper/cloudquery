package fetch

import (
	"github.com/spf13/cobra"
)

const (
	fetchShort = "Fetch resources from configured providers"
	fetchLong  = `Fetch resources from configured providers
	
	This requires a cloudquery.yml file which can be generated by "cloudquery init"
	`
	fetchExample = `  # Fetch configured providers to PostgreSQL as configured in cloudquery.yml
	cloudquery fetch`
)

func NewCmdFetch() *cobra.Command {
	fetchCmd := &cobra.Command{
		Use:     "fetch",
		Short:   fetchShort,
		Long:    fetchLong,
		Example: fetchExample,
		Args:    cobra.ExactArgs(1),
		RunE:    fetch,
	}

	return fetchCmd
}

func fetch(cmd *cobra.Command, args []string) error {
	return nil
}
