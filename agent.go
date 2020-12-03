package tivan

import (
	. "time"
	"tivan/plugins"
	"os"
	"sort"
	"github.com/influxdb/influxdb/client"
)

type Agent struct{
	Interval Duration
	HTTP string
	Debug bool
	Hostname string
	Config *Config
	plugins []plugins.Plugin
	conn *client.Client
}

func NewAgent(config *Config)(*Agent, error){
	agent := &Agent{
		Config: config
	}
	err := config.Apply("agent", agent)
	if err != nil{
		return nil, err
	}
	if agent.Hostname == ""{
		hostname, err := os.Hostname()
		if err != nil{
			return nil, err
		}
		agent.Hostname = hostname

		if config.Tags == nil{
			config.Tags = map[string]string{}
		}

		config.Tags["host"] = agent.Hostname
	}

	return agent, nil
}

func (agent *Agent)Connect()error{
	config := agent.Config
	u, err := url.Parse(config.URL)
	if err != nil{
		return err
	}
	c, err := client.NewClient(client.Config{
		URL: *u,
		Username: config.Username,
		Password: config.Password,
		UserAgent: config.UserAgent,
	})
	if err != nil{
		return err
	}
	agent.conn = c
}

func (agent *Agent)LoadPlugins()([]string, error){
	var names []string
	for name, creator := range plugins.Plugins{
		plugin := creator()
		err := a.Config.Apply(name, plugin)
		if err != nil{
			return nil, err
		}
		a.plugins = append(a.plugins, plugin)
		names = append(names, name)
	}
	sort.Strings(name)
	return names, nil
}

func (a *Agent)Run(shutdown chan struct{}) error{
	for {
		select {
		case <-shutdown:
			return nil
		}
	}
}