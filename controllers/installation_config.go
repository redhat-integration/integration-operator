/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

// InstallationConfig stores configuration required for the installation of the product operators
type InstallationConfig struct {
	Channel3scale          string
	Channel3scaleAPIcast   string
	ChannelAMQBroker       string
	ChannelAMQInterconnect string
	ChannelAMQStreams      string
	ChannelAPIDesigner     string
	ChannelCamelK          string
	ChannelFuseConsole     string
	ChannelFuseOnline      string
	ChannelServiceRegistry string
}
