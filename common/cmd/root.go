package cmd

import (
	"fmt"

	"project/common/config"
	constant "project/proto"
	"project/tools/log"

	"github.com/spf13/cobra"
)

type RootCmd struct {
	Command        cobra.Command
	Name           string
	port           int
	prometheusPort int
}

type CmdOpts struct {
	loggerPrefixName string
}

func WithCronTaskLogName() func(*CmdOpts) {
	return func(opts *CmdOpts) {
		opts.loggerPrefixName = "MicroClean.CronTask.log.all"
	}
}

func WithLogName(logName string) func(*CmdOpts) {
	return func(opts *CmdOpts) {
		opts.loggerPrefixName = logName
	}
}

func NewRootCmd(name string, opts ...func(*CmdOpts)) (rootCmd *RootCmd) {
	rootCmd = &RootCmd{Name: name}
	c := cobra.Command{
		Use:   fmt.Sprintf("start micro-clean %s", name),
		Short: fmt.Sprintf(`Start %s `, name),
		Long:  fmt.Sprintf(`Start %s `, name),
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if err := rootCmd.getConfFromCmdAndInit(cmd); err != nil {
				panic(err)
			}
			cmdOpts := &CmdOpts{}
			for _, opt := range opts {
				opt(cmdOpts)
			}
			if cmdOpts.loggerPrefixName == "" {
				cmdOpts.loggerPrefixName = "MicroClean.log.all"
			}

			log.NewLogger(cmdOpts.loggerPrefixName, config.Config.Log.RotationTime, config.Config.Log.RemainRotationCount)

			return nil
		},
	}
	rootCmd.Command = c
	rootCmd.addConfFlag()
	return rootCmd
}

func (r *RootCmd) PreLoadConfig() {
	// if err := config.InitConfig(""); err != nil {
	// 	panic(err)
	// }
}

func (r *RootCmd) SetSvcName(name string) {
	r.Name = name
}

func (r *RootCmd) addConfFlag() {
	r.Command.Flags().StringP(constant.FlagConf, "c", "", "Path to config file folder")
}

func (r *RootCmd) AddPortFlag() {
	r.Command.Flags().IntP(constant.FlagPort, "p", 0, "server listen port")
}

func (r *RootCmd) getPortFlag(cmd *cobra.Command) int {
	port, _ := cmd.Flags().GetInt(constant.FlagPort)
	return port
}

func (r *RootCmd) GetPortFlag() int {
	return r.port
}

func (r *RootCmd) AddPrometheusPortFlag() {
	r.Command.Flags().IntP(constant.FlagPrometheusPort, "", 0, "server prometheus listen port")
}

func (r *RootCmd) getPrometheusPortFlag(cmd *cobra.Command) int {
	port, _ := cmd.Flags().GetInt(constant.FlagPrometheusPort)
	return port
}

func (r *RootCmd) GetPrometheusPortFlag() int {
	if r.prometheusPort == 0 {
		switch r.Name {
		case config.Config.RPCRegisterName.MicroCleanAuthName:
			return config.Config.Prometheus.AuthPrometheusPort[0]
		case config.Config.RPCRegisterName.MicroCleanUserName:
			return config.Config.Prometheus.UserPrometheusPort[0]
		case config.Config.RPCRegisterName.MicroCleanItemName:
			return config.Config.Prometheus.ItemPrometheusPort[0]
		case config.Config.RPCRegisterName.MicroCleanPushName:
			return config.Config.Prometheus.PushPrometheusPort[0]
		default:
			return 0

		}
	}

	return r.prometheusPort
}

func (r *RootCmd) getConfFromCmdAndInit(cmdLines *cobra.Command) error {
	configFolderPath, _ := cmdLines.Flags().GetString(constant.FlagConf)
	fmt.Println("configFolderPath:", configFolderPath)
	return config.LoadConfig(config.FileName, configFolderPath)
}

func (r *RootCmd) Execute() error {
	return r.Command.Execute()
}

func (r *RootCmd) AddCommand(cmds ...*cobra.Command) {
	r.Command.AddCommand(cmds...)
}
