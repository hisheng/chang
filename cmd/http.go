package cmd

import (
	"github.com/gin-gonic/gin"
	"github.com/hisheng/chang/handler"
	"github.com/spf13/cobra"
	"log"
)

var (
	HttpServerAddress string = ":8080"
)

var (
	httpCmd = &cobra.Command{
		Use:   "http [text]",
		Short: "开启http服务",
		Long:  `开启http服务`,

		Run: func(cmd *cobra.Command, args []string) {
			router := gin.Default()
			router.GET("/test/detail", handler.TestDetail)
			router.GET("/guxi-record", handler.GuxiRecordList)
			router.GET("/guxi-record/init", handler.GuxiRecordInit)
			router.GET("/area", handler.Area)
			router.GET("/area/init", handler.AreaInit)
			if err := router.Run(HttpServerAddress); err != nil {
				log.Fatal(err)
			}
		},
	}
)
