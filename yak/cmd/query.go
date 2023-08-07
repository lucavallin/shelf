// Copyright © 2018 Luca Cavallin <me@lucavall.in>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"github.com/spf13/cobra"
)

// queryCmd represents the query command
var (
	xmlFilePath string
	elapsedDays int
	queryCmd = &cobra.Command{
		Use:   "query",
		Short: "Import and query a herd XML file given a number of days",
		Long: "Import and query a herd XML file given a number of days",
		Run: func(cmd *cobra.Command, args []string) {
			Import()
		},
	}
)

func init() {
	queryCmd.Flags().StringVarP(&xmlFilePath, "file", "f", "","Path to Yak Webshop XML file")
	queryCmd.Flags().IntVarP(&elapsedDays, "elapsedDays", "d", 0,"Elapsed days to query stock for")
	rootCmd.AddCommand(queryCmd)
}

// Import queries a herd from the XML file
func Import() {
	herdProvider := herd.NewXMLFileRepository(xmlFilePath)
	herd := herdProvider.Get()

	stock := herd.GetStock(elapsedDays)

	fmt.Println("Stock:")
	fmt.Printf("\t%.3f liters of milk\n", stock.Milk)
	fmt.Printf("\t%d skins of wool\n", stock.Wool)
	fmt.Println("Herd:")
	for _, yak := range herd.Yaks {
		fmt.Printf("\t%s %.2f years old\n", yak.Name, yak.Age)
	}
}