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
	"github.com/lucavallin/yak-webshop/pkg/api"
	"github.com/lucavallin/yak-webshop/pkg/herd"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server for the Yak Webshop",
	Long: "Start API server for the Yak Webshop",
	Run: func(cmd *cobra.Command, args []string) {
		Serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

// Serve stars the APIs for the yak-webshop
func Serve() {
	// Not beautiful but at least the IO logic is isolated
	herdPath, _ := filepath.Abs("data/herd.xml")
	herdRepo := herd.NewXMLFileRepository(herdPath)
	app := api.NewApp(herdRepo)

	app.Run(os.Getenv("PORT"))
}