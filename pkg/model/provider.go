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

package model

type GkeProvider struct {
	Project     string  `json:"project"`
	ClusterName string  `json:"clusterName"`
	Zone        *string `json:"zone"`
	Region      *string `json:"region"`
}

type AksProvider struct {
	ResourceGroup string `json:"resourceGroup"`
	ClusterName   string `json:"clusterName"`
}

type MinikubeProvider struct {
}

type DockerForDesktopProvider struct {
}

type ExistingContextProvider struct {
	ContextName string `json:"contextName"`
}

type Provider struct {
	GKE              *GkeProvider              `json:"gke,omitempty"`
	Minikube         *MinikubeProvider         `json:"minikube,omitempty"`
	AKS              *AksProvider              `json:"aks,omitempty"`
	DockerForDesktop *DockerForDesktopProvider `json:"dockerForDesktop,omitempty"`
	ExistingContext  *ExistingContextProvider  `json:"existingContext,omitempty"`
}
