//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package cmd

import (
	"fmt"
	"github.com/lastbackend/cli/pkg/cli/envs"
	"github.com/lastbackend/cli/pkg/cli/storage"
	"github.com/lastbackend/cli/pkg/cli/view"
	"github.com/spf13/cobra"
)

func init() {
	clusterCmd.AddCommand(ClusterListCmd)
}

const clusterListExample = `
  # Get information about cluster 
  lb cluster ls
`

var ClusterListCmd = &cobra.Command{
	Use:     "ls",
	Short:   "Get available cluster list",
	Example: clusterListExample,
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Remotely clusters list:")

		cli := envs.Get().GetClient()

		ritems, err := cli.Genesis.V1().Cluster().List(envs.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		vg := view.FromGenesisApiClusterListView(ritems)
		vg.Print()

		fmt.Print("\n")
		fmt.Println("Locally clusters list:")

		litems, err := storage.ListLocalCluster()
		if err != nil {
			panic(err)
		}

		vs := view.FromStorageClusterList(litems)
		vs.Print()
	},
}
