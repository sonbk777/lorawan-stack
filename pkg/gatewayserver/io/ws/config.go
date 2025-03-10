// Copyright © 2020 The Things Network Foundation, The Things Industries B.V.
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

package ws

import "time"

// Config defines the LoRa Basics Station configuration of the Gateway Server.
type Config struct {
	UseTrafficTLSAddress bool          `name:"use-traffic-tls-address" description:"Use WSS for the traffic address regardless of the TLS setting"`
	WSPingInterval       time.Duration `name:"ws-ping-interval" description:"Interval to send WS ping messages"`
	MissedPongThreshold  int           `name:"missed-pong-threshold" description:"Number of consecutive missed pongs before disconnection. This value is used only if the gateway sends at least one pong."`
	TimeSyncInterval     time.Duration `name:"time-sync-interval" description:"Interval to send time transfer messages"`
	AllowUnauthenticated bool          `name:"allow-unauthenticated" description:"Allow unauthenticated connections"`
}

// DefaultConfig contains the default configuration.
var DefaultConfig = Config{
	UseTrafficTLSAddress: false,
	WSPingInterval:       30 * time.Second,
	// Assuming 5ppm of drift, this means a drift of 5 microseconds in one second.
	// A drift of 1 millisecond would occur every 200 seconds in such a situation.
	MissedPongThreshold:  2,
	TimeSyncInterval:     200 * time.Second,
	AllowUnauthenticated: false,
}
