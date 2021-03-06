// Copyright 2017 The OpenSDS Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"os"
)

type Client struct {
	*ProfileMgr
	*DockMgr
	*PoolMgr
	*VolumeMgr

	cfg *Config
}

type Config struct {
	Endpoint string
}

func NewClient(c *Config) *Client {
	// If endpoint field not specified, the info will be fetched from
	// environment variable.
	if c.Endpoint == "" {
		c.Endpoint = os.Getenv("OPENSDS_ENDPOINT")
	}

	return &Client{
		cfg:        c,
		ProfileMgr: NewProfileMgr(c.Endpoint),
		DockMgr:    NewDockMgr(c.Endpoint),
		PoolMgr:    NewPoolMgr(c.Endpoint),
		VolumeMgr:  NewVolumeMgr(c.Endpoint),
	}
}

func (c *Client) Reset() {
	c = &Client{}
}
