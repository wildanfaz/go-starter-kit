package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wildanfaz/go-starter-kit/internal/routers"
	"github.com/wildanfaz/go-starter-kit/internal/utils"
)

func InitCmd(ctx context.Context) {
	var rootCmd = &cobra.Command{
		Short: "Project Title",
	}

	rootCmd.AddCommand(startEchoApp)

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		panic(err)
	}
}

var startEchoApp = &cobra.Command{
	Use:   "start",
	Short: "Start the application",
	Run: func(cmd *cobra.Command, args []string) {
		// logger
		utils.InitLogger()

		// routers
		routers.InitEchoRouter()
	},
}
