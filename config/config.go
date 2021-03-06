// Copyright 2016 The kingshard Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

//整个config文件对应的结构
type Config struct {
	Addr        string       `yaml:"addr"`
	Ups         string       `yaml:"ups"`
	LogPath     string       `yaml:"log_path"`
	LogLevel    string       `yaml:"log_level"`
	LogSql      string       `yaml:"log_sql"`
	SlowLogTime int          `yaml:"slow_log_time"`
	AllowIps    string       `yaml:"allow_ips"`
	BlsFile     string       `yaml:"blacklist_sql_file"`
	Charset     string       `yaml:"proxy_charset"`
	Nodes       []NodeConfig `yaml:"nodes"`

	Schema SchemaConfig `yaml:"schema"`
}

//node节点对应的配置
type NodeConfig struct {
	Name             string `yaml:"name"`
	DownAfterNoAlive int    `yaml:"down_after_noalive"`
	MaxConnNum       int    `yaml:"max_conns_limit"`

	Master string `yaml:"master"`
	Slave  string `yaml:"slave"`
}

//schema对应的结构体
type SchemaConfig struct {
	DB        string        `yaml:"db"`
	Nodes     []string      `yaml:"nodes"`
	Default   string        `yaml:"default"` //default route rule
	ShardRule []ShardConfig `yaml:"shard"`   //route rule
}

//range,hash or date
type ShardConfig struct {
	Table         string   `yaml:"table"`
	Key           string   `yaml:"key"`
	Nodes         []string `yaml:"nodes"`
	Locations     []int    `yaml:"locations"`
	Type          string   `yaml:"type"`
	TableRowLimit int      `yaml:"table_row_limit"`
	DateRange     []string `yaml:"date_range"`
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func ParseConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}

func WriteConfigFile(cfg *Config) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	execPath, err := os.Getwd()
	if err != nil {
		return err
	}

	configPath := execPath + "/etc/ks.yaml"
	err = ioutil.WriteFile(configPath, data, 0755)
	if err != nil {
		return err
	}

	return nil
}
