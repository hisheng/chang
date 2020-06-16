package cmd

import (
	"github.com/hisheng/chang/conf"
	"github.com/hisheng/chang/db"
	"github.com/hisheng/chang/xueqiu"
	"github.com/robfig/cron"
	"github.com/spf13/cobra"
)

var (
	//CRON Expression Format
	//A cron expression represents a set of times, using 6 space-separated fields.
	//
	//Field name   | Mandatory? | Allowed values  | Allowed special characters
	//----------   | ---------- | --------------  | --------------------------
	//Seconds      | Yes        | 0-59            | * / , -
	//Minutes      | Yes        | 0-59            | * / , -
	//Hours        | Yes        | 0-23            | * / , -
	//Day of month | Yes        | 1-31            | * / , - ?
	//Month        | Yes        | 1-12 or JAN-DEC | * / , -
	//Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
	CRON                     *cron.Cron
	everyMinuteFifteenSecond = "15 * * * * *"
	everyFourSecond          = "*/4 * * * * *"
	everyDay                 = "4 21 16 * * *"
)

func init() {
	conf.InitConf()
	db.GetDb()
}

var (
	crontabCmd = &cobra.Command{
		Use:   "crontab",
		Short: "开启crontab服务",
		Long:  `开启crontab服务`,

		Run: func(cmd *cobra.Command, args []string) {
			go xueqiu.UpdateData()
			//fmt.Println(b.val(a,"name"))

			CRON = cron.New()

			// debug 开发环境 配置测试
			if conf.Conf.Debug {
				// 每5秒 打印一下时间
				//CRON.AddFunc(everyMinuteFifteenSecond, func() {fmt.Println(time.Now().Date())})
			}

			CRON.AddFunc(everyDay, func() {
				go xueqiu.UpdateData()
			})

			CRON.Start()
			defer CRON.Stop()

			select {}
		},
	}
)
