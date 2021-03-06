/*
 * Copyright 2018-2020 Zedge, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package provider

import "github.com/kubecd/kubecd/pkg/model"

type GitlabClusterProvider struct{ baseClusterProvider }

func (p *GitlabClusterProvider) GetClusterName() string {
	return "gitlab-deploy"
}

func (p *GitlabClusterProvider) GetUserName() string {
	return "gitlab-deploy"
}

func (p *GitlabClusterProvider) GetNamespace(env *model.Environment) string {
	return env.KubeNamespace
}

func (p *GitlabClusterProvider) GetClusterInitCommands() ([][]string, error) {
	return [][]string{}, nil
}
