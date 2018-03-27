// Copyright 2016 Palantir Technologies, Inc.
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

package cmd

import (
	"github.com/palantir/godel/framework/pluginapi/v2/pluginapi"
	"github.com/palantir/pkg/cobracli"
)

var PluginInfo = pluginapi.MustNewPluginInfo(
	"com.palantir.godel-goland-plugin",
	"goland-plugin",
	cobracli.Version,
	pluginapi.PluginInfoGlobalFlagOptions(
		pluginapi.GlobalFlagOptionsParamProjectDirFlag("--"+pluginapi.ProjectDirFlagName),
	),
	pluginapi.PluginInfoTaskInfo(
		"goland",
		"GoLand project commands",
	),
)