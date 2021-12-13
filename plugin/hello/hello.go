package hello

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"murphysec-cli-simple/plugin/plugin_base"
)

type Plugin struct {
	arg string
}

var Instance plugin_base.Plugin = &Plugin{}

func (_ *Plugin) Info() plugin_base.PluginInfo {
	return plugin_base.PluginInfo{Name: "hello", ShortDescription: "just a hello world"}
}

func (p *Plugin) MatchPath(ctx context.Context, dir string) bool {
	fmt.Println("hello world MatchPath", p.arg)
	return false
}

func (p *Plugin) DoScan(ctx context.Context, dir string) interface{} {
	fmt.Println("hello world DoScan", p.arg)
	return nil
}

func (p *Plugin) SetupScanCmd(c *cobra.Command) {
	c.PersistentFlags().StringVarP(&p.arg, "foo", "", "", "--foo bar")
}