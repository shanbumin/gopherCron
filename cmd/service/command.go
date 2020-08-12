package service

import (
	"github.com/spf13/cobra"
)

// NewCommand of service
func NewCommand() *cobra.Command {
	opt := newSetupOptions()
	cmd := &cobra.Command{
		Use:   "service",
		Short: "A api service",
		RunE: func(c *cobra.Command, args []string) error {
			//fmt.Printf("%+v\r\n",args) [param1  param2  param3 ...]
			//todo 暂时不需要参数
			//传递选项opt进行处理好了,启动service
			//fmt.Printf("%+v\r\n",opt)
			if err := Run(opt); err != nil {
				return err
			}
			return nil
		},
	}
	//添加命令行选项
	opt.AddFlags(cmd.Flags())
	//fmt.Printf("%+v\r\n",opt)  //todo 这个时候输出是默认值，因为这个时候都还没有解析命令行参数呢
	//这个时候才返回cmd,这样就巧妙的避免使用init来添加Flags了
	return cmd
}
