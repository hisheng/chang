package croncmd

import (
	"github.com/hisheng/chang/repository"
	"github.com/hisheng/chang/service"
	"github.com/hisheng/chang/xueqiu"
	"github.com/spf13/cobra"
)

/*同步 bonus 分红 详情记录**/

var (
	BonusCmd = &cobra.Command{
		Use:   "bonus [text]",
		Short: "同步 bonus 分红 详情记录",
		Long:  `同步 bonus 分红 详情记录`,

		Run: func(cmd *cobra.Command, args []string) {
			bonusReposirory := repository.NewBonusRepository()
			bonusReposirory.CreateTable()

			bonusService := service.NewBonusService()

			symbol := xueqiu.Symbol_{}
			symbols := symbol.Gets()
			for _, symb := range symbols {
				bonusService.CurlGet(symb.Symbol)
			}
		},
	}
)
