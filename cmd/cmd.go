package cmd

import (
	_ "blog/apps"
	"blog/apps/user"
	"blog/conf"
	"blog/ioc"
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	configPath string
)

var (
	createUsername string
	createPassword string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a now user",
	RunE: func(cmd *cobra.Command, args []string) error {
		userService, ok := ioc.Contorller.Get(user.AppName).(user.Service)
		if !ok{
			return fmt.Errorf("user service not found in ioc container")
		}
		req:=user.NewCreateUserReuqest()
		req.Username=createUsername
		req.Password=createPassword
		u,err:=userService.CreateUser(context.Background(),req)
		if err!=nil{
			return err
		}
		fmt.Printf("create user success, id=%d username=%s\n",u.Id,u.Username)
		return nil
	},
}

var RootCmd = &cobra.Command{
	Use:   "vblog",
	Short: "vblog service",
	Run: func(cmd *cobra.Command, args []string) {
		envConfigPath := os.Getenv("CONFIG_PATH")
		if envConfigPath == "" {
			envConfigPath = "etc/application.yaml"
		}
		cobra.CheckErr(conf.C().App.Start())
	},
}

func Execute() error {
	cobra.OnInitialize(func() {
		cobra.CheckErr(conf.FromYaml(configPath))
		cobra.CheckErr(ioc.Contorller.Init())
		cobra.CheckErr(ioc.Api.Init())
	})
	return RootCmd.Execute()
}

func init() {
	createCmd.Flags().StringVarP(&createUsername, "username", "u", "", "username of the new user")
	createCmd.Flags().StringVarP(&createPassword, "password", "p", "", "password of the new user")
	cobra.CheckErr(createCmd.MarkFlagRequired("username"))
	cobra.CheckErr(createCmd.MarkFlagRequired("password"))
	RootCmd.AddCommand(createCmd)
	RootCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "etc/application.yaml", "the service config file")
}
