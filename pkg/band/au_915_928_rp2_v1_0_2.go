// Copyright © 2021 The Things Network Foundation, The Things Industries B.V.
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

package band

import "go.thethings.network/lorawan-stack/v3/pkg/ttnpb"

// AU_915_928_RP2_v1_0_2 is the band definition for AU915-928 in the RP002-1.0.2 specification.
var AU_915_928_RP2_v1_0_2 = Band{
	ID: AU_915_928,

	SupportsDynamicADR: true,

	MaxUplinkChannels: 72,
	UplinkChannels:    au915928UplinkChannels(0),

	MaxDownlinkChannels: 8,
	DownlinkChannels:    au915928DownlinkChannels,

	// See Radiocommunications (Low Interference Potential Devices) Class Licence 2015
	SubBands: []SubBandParameters{
		{
			MinFrequency: 915000000,
			MaxFrequency: 928000000,
			DutyCycle:    1,
			MaxEIRP:      30,
		},
	},

	DataRates: map[ttnpb.DataRateIndex]DataRate{
		ttnpb.DataRateIndex_DATA_RATE_0:  makeLoRaDataRate(12, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(59, 0)),
		ttnpb.DataRateIndex_DATA_RATE_1:  makeLoRaDataRate(11, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(59, 0)),
		ttnpb.DataRateIndex_DATA_RATE_2:  makeLoRaDataRate(10, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(59, 19)),
		ttnpb.DataRateIndex_DATA_RATE_3:  makeLoRaDataRate(9, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(123, 61)),
		ttnpb.DataRateIndex_DATA_RATE_4:  makeLoRaDataRate(8, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 133)),
		ttnpb.DataRateIndex_DATA_RATE_5:  makeLoRaDataRate(7, 125000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250)),
		ttnpb.DataRateIndex_DATA_RATE_6:  makeLoRaDataRate(8, 500000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250)),
		ttnpb.DataRateIndex_DATA_RATE_7:  makeLRFHSSDataRate(0, 1523000, "1/3", makeDwellTimeMaxMACPayloadSizeFunc(58, 58)),
		ttnpb.DataRateIndex_DATA_RATE_8:  makeLoRaDataRate(12, 500000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(61, 61)),
		ttnpb.DataRateIndex_DATA_RATE_9:  makeLoRaDataRate(11, 500000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(137, 137)),
		ttnpb.DataRateIndex_DATA_RATE_10: makeLoRaDataRate(10, 500000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250)),
		ttnpb.DataRateIndex_DATA_RATE_11: makeLoRaDataRate(9, 500000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250)),
		ttnpb.DataRateIndex_DATA_RATE_12: makeLoRaDataRate(8, 500000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250)),
		ttnpb.DataRateIndex_DATA_RATE_13: makeLoRaDataRate(7, 500000, Cr4_5, makeDwellTimeMaxMACPayloadSizeFunc(250, 250)),
	},
	MaxADRDataRateIndex: ttnpb.DataRateIndex_DATA_RATE_5,

	ReceiveDelay1:        defaultReceiveDelay1,
	ReceiveDelay2:        defaultReceiveDelay2,
	JoinAcceptDelay1:     defaultJoinAcceptDelay1,
	JoinAcceptDelay2:     defaultJoinAcceptDelay2,
	MaxFCntGap:           defaultMaxFCntGap,
	ADRAckLimit:          defaultADRAckLimit,
	ADRAckDelay:          defaultADRAckDelay,
	MinRetransmitTimeout: defaultRetransmitTimeout - defaultRetransmitTimeoutMargin,
	MaxRetransmitTimeout: defaultRetransmitTimeout + defaultRetransmitTimeoutMargin,

	DefaultMaxEIRP: 30,
	TxOffset: []float32{
		0,
		-2,
		-4,
		-6,
		-8,
		-10,
		-12,
		-14,
		-16,
		-18,
		-20,
		-22,
		-24,
		-26,
		-28,
	},

	FreqMultiplier:   100,
	ImplementsCFList: true,
	CFListType:       ttnpb.CFListType_CHANNEL_MASKS,

	Rx1Channel: channelIndexModulo(8),
	Rx1DataRate: func(idx ttnpb.DataRateIndex, offset ttnpb.DataRateOffset, _ bool) (ttnpb.DataRateIndex, error) {
		if idx > ttnpb.DataRateIndex_DATA_RATE_7 {
			return 0, errDataRateIndexTooHigh.WithAttributes("max", 7)
		}
		if offset > 5 {
			return 0, errDataRateOffsetTooHigh.WithAttributes("max", 5)
		}
		return au915928DownlinkDRTable[idx][offset], nil
	},

	GenerateChMasks: makeGenerateChMask72(true),
	ParseChMask:     parseChMask72,

	DefaultRx2Parameters: Rx2Parameters{ttnpb.DataRateIndex_DATA_RATE_8, 923300000},

	Beacon: Beacon{
		DataRateIndex: ttnpb.DataRateIndex_DATA_RATE_8,
		CodingRate:    Cr4_5,
		Frequencies:   usAuBeaconFrequencies,
	},
	PingSlotFrequencies: usAuBeaconFrequencies,

	TxParamSetupReqSupport: true,

	BootDwellTime: DwellTime{
		Uplinks:   boolPtr(true),
		Downlinks: boolPtr(false),
	},
}
