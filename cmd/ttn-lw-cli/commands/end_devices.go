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

package commands

import (
	"bufio"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/hex"
	stdio "io"
	"mime"
	"os"
	"path"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/util"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/random"
	"go.thethings.network/lorawan-stack/v3/pkg/specification/macspec"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/types"
	"google.golang.org/grpc"
)

var (
	selectEndDeviceListFlags   = util.NormalizedFlagSet()
	selectEndDeviceFlags       = util.NormalizedFlagSet()
	setEndDeviceFlags          = util.NormalizedFlagSet()
	endDevicePictureFlags      = util.NormalizedFlagSet()
	endDeviceLocationFlags     = util.NormalizedFlagSet()
	getDefaultMACSettingsFlags = util.NormalizedFlagSet()
	allEndDeviceSetFlags       = util.NormalizedFlagSet()
	allEndDeviceSelectFlags    = util.NormalizedFlagSet()
	listBandsFlags             = util.NormalizedFlagSet()
	listPhyVersionFlags        = util.NormalizedFlagSet()
	getNetIDFlags              = util.NormalizedFlagSet()
	getDevAddrPrefixesFlags    = util.NormalizedFlagSet()

	selectAllEndDeviceFlags = util.SelectAllFlagSet("end devices")
	toUnderscore            = strings.NewReplacer("-", "_")

	claimAuthenticationCodePaths = []string{
		"claim_authentication_code",
		"claim_authentication_code.value",
		"claim_authentication_code.valid_from",
		"claim_authentication_code.valid_to",
	}
)

func selectEndDeviceIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.Bool("application-id", false, "")
	flagSet.Bool("device-id", false, "")
	flagSet.Bool("join-eui", false, "")
	flagSet.Bool("dev-eui", false, "")
	addDeprecatedDeviceFlags(flagSet)
	return flagSet
}

func endDeviceIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("application-id", "", "")
	flagSet.String("device-id", "", "")
	flagSet.String("join-eui", "", "(hex)")
	flagSet.String("dev-eui", "", "(hex)")
	addDeprecatedDeviceFlags(flagSet)
	return flagSet
}

func addDeprecatedDeviceFlags(flagSet *pflag.FlagSet) {
	util.DeprecateFlag(flagSet, "app-eui", "join-eui")
	util.DeprecateFlag(flagSet, "session.keys.nwk_s_key", "session.keys.f_nwk_s_int_key")
	util.DeprecateFlag(flagSet, "pending_session.keys.nwk_s_key", "pending_session.keys.f_nwk_s_int_key")
	util.DeprecateFlag(flagSet, "session.keys.nwk_s_key.key", "session.keys.f_nwk_s_int_key.key")
	util.DeprecateFlag(flagSet, "pending_session.keys.nwk_s_key.key", "pending_session.keys.f_nwk_s_int_key.key")

	util.HideFlag(flagSet, "mac_settings.use_adr")
	util.HideFlag(flagSet, "mac_settings.adr_margin")
}

func forwardDeprecatedDeviceFlags(flagSet *pflag.FlagSet) {
	util.ForwardFlag(flagSet, "app-eui", "join-eui")
	util.ForwardFlag(flagSet, "session.keys.nwk_s_key", "session.keys.f_nwk_s_int_key")
	util.ForwardFlag(flagSet, "pending_session.keys.nwk_s_key", "pending_session.keys.f_nwk_s_int_key")
	util.ForwardFlag(flagSet, "session.keys.nwk_s_key.key", "session.keys.f_nwk_s_int_key.key")
	util.ForwardFlag(flagSet, "pending_session.keys.nwk_s_key.key", "pending_session.keys.f_nwk_s_int_key.key")
}

var (
	errConflictingPaths             = errors.DefineInvalidArgument("conflicting_paths", "conflicting set and unset field mask paths")
	errEndDeviceEUIUpdate           = errors.DefineInvalidArgument("end_device_eui_update", "end device EUIs can not be updated")
	errEndDeviceKeysWithProvisioner = errors.DefineInvalidArgument("end_device_keys_provisioner", "end device ABP or OTAA keys cannot be set when there is a provisioner")
	errInconsistentEndDeviceEUI     = errors.DefineInvalidArgument("inconsistent_end_device_eui", "given end device EUIs do not match registered EUIs")
	errInvalidMACVersion            = errors.DefineInvalidArgument("mac_version", "LoRaWAN MAC version is invalid")
	errInvalidPHYVersion            = errors.DefineInvalidArgument("phy_version", "LoRaWAN PHY version is invalid")
	errNoEndDeviceEUI               = errors.DefineInvalidArgument("no_end_device_eui", "no end device EUIs set")
	errInvalidJoinEUI               = errors.DefineInvalidArgument("invalid_join_eui", "invalid JoinEUI")
	errInvalidDevEUI                = errors.DefineInvalidArgument("invalid_dev_eui", "invalid DevEUI")
	errInvalidNetID                 = errors.DefineInvalidArgument("invalid_net_id", "invalid NetID")
	errNoEndDeviceID                = errors.DefineInvalidArgument("no_end_device_id", "no end device ID set")
)

func getEndDeviceID(flagSet *pflag.FlagSet, args []string, requireID bool) (*ttnpb.EndDeviceIdentifiers, error) {
	forwardDeprecatedDeviceFlags(flagSet)
	applicationID, _ := flagSet.GetString("application-id")
	deviceID, _ := flagSet.GetString("device-id")
	switch len(args) {
	case 0:
	case 1:
		logger.Warn("Only single ID found in arguments, not considering arguments")
	case 2:
		applicationID = args[0]
		deviceID = args[1]
	default:
		logger.Warn("Multiple IDs found in arguments, considering the first")
		applicationID = args[0]
		deviceID = args[1]
	}
	if applicationID == "" && requireID {
		return nil, errNoApplicationID.New()
	}
	if deviceID == "" && requireID {
		return nil, errNoEndDeviceID.New()
	}
	ids := &ttnpb.EndDeviceIdentifiers{
		ApplicationIds: &ttnpb.ApplicationIdentifiers{ApplicationId: applicationID},
		DeviceId:       deviceID,
	}
	if joinEUIHex, _ := flagSet.GetString("join-eui"); joinEUIHex != "" {
		var joinEUI types.EUI64
		if err := joinEUI.UnmarshalText([]byte(joinEUIHex)); err != nil {
			return nil, errInvalidJoinEUI.WithCause(err)
		}
		ids.JoinEui = joinEUI.Bytes()
	}
	if devEUIHex, _ := flagSet.GetString("dev-eui"); devEUIHex != "" {
		var devEUI types.EUI64
		if err := devEUI.UnmarshalText([]byte(devEUIHex)); err != nil {
			return nil, errInvalidDevEUI.WithCause(err)
		}
		ids.DevEui = devEUI.Bytes()
	}
	return ids, nil
}

func generateKey() *types.AES128Key {
	var key types.AES128Key
	rand.Read(key[:])
	return &key
}

var (
	errJoinServerDisabled    = errors.DefineFailedPrecondition("join_server_disabled", "Join Server is disabled")
	errNetworkServerDisabled = errors.DefineFailedPrecondition("network_server_disabled", "Network Server is disabled")
	errEndDeviceClaimInfo    = errors.DefineFailedPrecondition("end_device_claim_info", "could not get end device claim info from DCS")
	errEndDeviceClaim        = errors.DefineFailedPrecondition("end_device_claim", "could not claim end device")
	errEndDeviceUnclaim      = errors.DefineFailedPrecondition("end_device_unclaim", "could not unclaim end device")
)

var (
	endDevicesCommand = &cobra.Command{
		Use:     "end-devices",
		Aliases: []string{"end-device", "devices", "device", "dev", "ed", "d"},
		Short:   "End Device commands",
	}
	endDevicesListFrequencyPlans = &cobra.Command{
		Use:               "list-frequency-plans",
		Aliases:           []string{"get-frequency-plans", "frequency-plans", "fps"},
		Short:             "List available frequency plans for end devices",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !config.NetworkServerEnabled {
				return errNetworkServerDisabled.New()
			}

			baseFrequency, _ := cmd.Flags().GetUint32("base-frequency")
			ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewConfigurationClient(ns).ListFrequencyPlans(ctx, &ttnpb.ListFrequencyPlansRequest{
				BaseFrequency: baseFrequency,
			})
			if err != nil {
				return err
			}
			return io.Write(os.Stdout, config.OutputFormat, res.FrequencyPlans)
		},
	}
	endDevicesListCommand = &cobra.Command{
		Use:     "list [application-id]",
		Aliases: []string{"ls"},
		Short:   "List end devices",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedDeviceFlags(cmd.Flags())

			appID := getApplicationID(cmd.Flags(), args)
			if appID == nil {
				return errNoApplicationID.New()
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectEndDeviceListFlags)
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.EndDeviceRegistry/List"].Allowed)

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			limit, page, opt, getTotal := withPagination(cmd.Flags())
			res, err := ttnpb.NewEndDeviceRegistryClient(is).List(ctx, &ttnpb.ListEndDevicesRequest{
				ApplicationIds: appID,
				FieldMask:      ttnpb.FieldMask(paths...),
				Limit:          limit,
				Page:           page,
				Order:          getOrder(cmd.Flags()),
			}, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res.EndDevices)
		},
	}
	endDevicesSearchCommand = &cobra.Command{
		Use:   "search [application-id]",
		Short: "Search for end devices",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedDeviceFlags(cmd.Flags())

			appID := getApplicationID(cmd.Flags(), args)
			if appID == nil {
				return errNoApplicationID.New()
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectEndDeviceListFlags)

			req := &ttnpb.SearchEndDevicesRequest{}
			_, err := req.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			var (
				opt      grpc.CallOption
				getTotal func() uint64
			)
			req.Limit, req.Page, opt, getTotal = withPagination(cmd.Flags())
			req.ApplicationIds = appID
			req.FieldMask = ttnpb.FieldMask(paths...)

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewEndDeviceRegistrySearchClient(is).SearchEndDevices(ctx, req, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res.EndDevices)
		},
	}
	endDevicesGetCommand = &cobra.Command{
		Use:     "get [application-id] [device-id]",
		Aliases: []string{"info"},
		Short:   "Get an end device",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedDeviceFlags(cmd.Flags())

			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectEndDeviceFlags)

			isPaths, nsPaths, asPaths, jsPaths := splitEndDeviceGetPaths(paths...)

			if len(nsPaths) > 0 {
				isPaths = append(isPaths, "network_server_address")
			}
			if len(asPaths) > 0 {
				isPaths = append(isPaths, "application_server_address")
			}
			if len(jsPaths) > 0 {
				isPaths = append(isPaths, "join_server_address")
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			logger.WithField("paths", isPaths).Debug("Get end device from Identity Server")
			device, err := ttnpb.NewEndDeviceRegistryClient(is).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIds: devID,
				FieldMask:    ttnpb.FieldMask(isPaths...),
			})
			if err != nil {
				return err
			}

			if len(jsPaths) > 0 && device.JoinServerAddress == "" {
				logger.WithField("paths", jsPaths).Debug("End Device uses external Join Server, deselecting Join Server paths")
				jsPaths = nil
			}

			nsMismatch, asMismatch, jsMismatch := compareServerAddressesEndDevice(device, config)
			if len(nsPaths) > 0 && nsMismatch {
				logger.WithField("paths", nsPaths).Warn("Deselecting Network Server paths")
				nsPaths = nil
			}
			if len(asPaths) > 0 && asMismatch {
				logger.WithField("paths", asPaths).Warn("Deselecting Application Server paths")
				asPaths = nil
			}
			if len(jsPaths) > 0 && jsMismatch {
				logger.WithField("paths", jsPaths).Warn("Deselecting Join Server paths")
				jsPaths = nil
			}

			if len(jsPaths) > 0 && device.ClaimAuthenticationCode.GetValue() != "" {
				// ClaimAuthenticationCode is already retrieved from the IS. We can unset the related JS paths.
				jsPaths = ttnpb.ExcludeFields(jsPaths, claimAuthenticationCodePaths...)
			}

			res, err := getEndDevice(device.Ids, nsPaths, asPaths, jsPaths, true)
			if err != nil {
				return err
			}

			if err := device.SetFields(res, "ids.dev_addr"); err != nil {
				return err
			}
			if err := device.SetFields(res, append(append(nsPaths, asPaths...), jsPaths...)...); err != nil {
				return err
			}
			if device.CreatedAt == nil || (res.CreatedAt != nil && ttnpb.StdTime(res.CreatedAt).Before(*ttnpb.StdTime(device.CreatedAt))) {
				device.CreatedAt = res.CreatedAt
			}
			if res.UpdatedAt != nil && ttnpb.StdTime(res.UpdatedAt).After(*ttnpb.StdTime(device.UpdatedAt)) {
				device.UpdatedAt = res.UpdatedAt
			}
			return io.Write(os.Stdout, config.OutputFormat, device)
		},
	}
	endDevicesCreateCommand = &cobra.Command{
		Use:     "create [application-id] [device-id]",
		Aliases: []string{"add", "register"},
		Short:   "Create an end device",
		RunE: asBulk(func(cmd *cobra.Command, args []string) (err error) {
			forwardDeprecatedDeviceFlags(cmd.Flags())

			devID, err := getEndDeviceID(cmd.Flags(), args, false)
			if err != nil {
				return err
			}
			paths := util.UpdateFieldMask(cmd.Flags(), setEndDeviceFlags)

			abp, _ := cmd.Flags().GetBool("abp")
			multicast, _ := cmd.Flags().GetBool("multicast")
			abp = abp || multicast
			device := &ttnpb.EndDevice{
				Ids: &ttnpb.EndDeviceIdentifiers{},
			}
			if inputDecoder != nil {
				err := inputDecoder.Decode(device)
				if err != nil {
					return err
				}
				decodedPaths := ttnpb.NonZeroFields(device, ttnpb.EndDeviceFieldPathsNestedWithoutWrappers...)
				decodedPaths = ttnpb.BottomLevelFields(decodedPaths)
				paths = ttnpb.AddFields(paths, decodedPaths...)

				if abp && device.SupportsJoin {
					logger.Warn("Reading from standard input, ignoring --abp and --multicast flags")
				}
				abp = !device.SupportsJoin
			}

			setDefaults, _ := cmd.Flags().GetBool("defaults")
			if setDefaults {
				if config.NetworkServerEnabled {
					device.NetworkServerAddress = getHost(config.NetworkServerGRPCAddress)
					paths = append(paths, "network_server_address")
				}
				if config.ApplicationServerEnabled {
					device.ApplicationServerAddress = getHost(config.ApplicationServerGRPCAddress)
					paths = append(paths, "application_server_address")
				}
			}

			if picture, err := cmd.Flags().GetString("picture"); err == nil && picture != "" {
				device.Picture, err = readPicture(picture)
				if err != nil {
					return err
				}
			}

			if abp {
				device.SupportsJoin = false
				if config.NetworkServerEnabled {
					paths = append(paths, "supports_join")
				}
				if withSession, _ := cmd.Flags().GetBool("with-session"); withSession {
					if device.ProvisionerId != "" {
						return errEndDeviceKeysWithProvisioner.New()
					}
					ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
					if err != nil {
						return err
					}
					devAddrRes, err := ttnpb.NewNsClient(ns).GenerateDevAddr(ctx, ttnpb.Empty)
					if err != nil {
						return err
					}
					device.Ids.DevAddr = devAddrRes.DevAddr
					device.Session = &ttnpb.Session{
						DevAddr: devAddrRes.DevAddr,
						Keys: &ttnpb.SessionKeys{
							FNwkSIntKey: &ttnpb.KeyEnvelope{Key: generateKey().Bytes()},
							AppSKey:     &ttnpb.KeyEnvelope{Key: generateKey().Bytes()},
						},
					}
					paths = append(paths,
						"session.keys.session_key_id",
						"session.keys.f_nwk_s_int_key.key",
						"session.keys.app_s_key.key",
						"session.dev_addr",
					)
					var macVersion ttnpb.MACVersion
					s, err := setEndDeviceFlags.GetString("lorawan_version")
					if err != nil {
						return err
					}
					if err := macVersion.UnmarshalText([]byte(s)); err != nil {
						return errInvalidMACVersion.WithCause(err)
					}
					if err := macVersion.Validate(); err != nil {
						return errInvalidMACVersion.WithCause(err)
					}
					if macspec.UseNwkKey(macVersion) {
						device.Session.Keys.SNwkSIntKey = &ttnpb.KeyEnvelope{Key: generateKey().Bytes()}
						device.Session.Keys.NwkSEncKey = &ttnpb.KeyEnvelope{Key: generateKey().Bytes()}
						paths = append(paths,
							"session.keys.s_nwk_s_int_key.key",
							"session.keys.nwk_s_enc_key.key",
						)
					}
				}
			} else {
				device.SupportsJoin = true
				if config.NetworkServerEnabled {
					paths = append(paths, "supports_join")
				}
				if setDefaults {
					if config.JoinServerEnabled {
						device.JoinServerAddress = getHost(config.JoinServerGRPCAddress)
						paths = append(paths,
							"join_server_address",
						)
						if device.Ids.JoinEui == nil && (devID == nil || devID.JoinEui == nil) {
							// Get the default JoinEUI for Join Server.
							logger.WithField("join_server_address", config.JoinServerGRPCAddress).Info("JoinEUI empty but defaults flag is set, fetch default JoinEUI of the Join Server")
							js, err := api.Dial(ctx, config.JoinServerGRPCAddress)
							if err != nil {
								return err
							}
							res, err := ttnpb.NewJsClient(js).GetDefaultJoinEUI(ctx, ttnpb.Empty)
							if err != nil {
								return err
							}
							joinEUI := types.MustEUI64(res.JoinEui)
							logger.WithField("default_join_eui", joinEUI).Info("Successfully obtained Join Server's default JoinEUI")
							device.Ids.JoinEui = joinEUI.Bytes()
						}
					}
				}
				if withKeys, _ := cmd.Flags().GetBool("with-root-keys"); withKeys {
					if device.ProvisionerId != "" {
						return errEndDeviceKeysWithProvisioner.New()
					}
					device.RootKeys = &ttnpb.RootKeys{
						RootKeyId: "ttn-lw-cli-generated",
						AppKey:    &ttnpb.KeyEnvelope{Key: generateKey().Bytes()},
					}
					if s, err := setEndDeviceFlags.GetString("lorawan_version"); err == nil && s != "" {
						var macVersion ttnpb.MACVersion
						if err := macVersion.UnmarshalText([]byte(s)); err != nil {
							return errInvalidMACVersion.WithCause(err)
						}
						if err := macVersion.Validate(); err != nil {
							return errInvalidMACVersion.WithCause(err)
						}
						if macspec.UseNwkKey(macVersion) {
							device.RootKeys.NwkKey = &ttnpb.KeyEnvelope{Key: generateKey().Bytes()}
						}
					}
					paths = append(paths,
						"root_keys.root_key_id",
						"root_keys.app_key.key",
						"root_keys.nwk_key.key",
					)
				}
			}
			if withClaimAuthenticationCode, _ := cmd.Flags().GetBool("with-claim-authentication-code"); withClaimAuthenticationCode {
				device.ClaimAuthenticationCode = &ttnpb.EndDeviceAuthenticationCode{
					Value: strings.ToUpper(hex.EncodeToString(random.Bytes(4))),
				}
				paths = append(paths, "claim_authentication_code")
			}

			claimOnExternalJS := len(device.ClaimAuthenticationCode.GetValue()) > 0

			if hasUpdateDeviceLocationFlags(cmd.Flags()) {
				updateDeviceLocation(device, cmd.Flags())
				paths = append(paths, "locations")
			}

			_, err = device.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}

			device.Attributes = mergeAttributes(device.Attributes, cmd.Flags())
			if devID != nil {
				if devID.DeviceId != "" {
					device.Ids.DeviceId = devID.DeviceId
				}
				if devID.ApplicationIds != nil {
					device.Ids.ApplicationIds = devID.ApplicationIds
				}
				if device.SupportsJoin && devID.JoinEui != nil {
					device.Ids.JoinEui = devID.JoinEui
				}
				if devID.DevEui != nil {
					device.Ids.DevEui = devID.DevEui
				}
			}
			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}

			application, err := ttnpb.NewApplicationRegistryClient(is).Get(ctx, &ttnpb.GetApplicationRequest{
				ApplicationIds: devID.ApplicationIds,
				FieldMask: ttnpb.FieldMask(
					"network_server_address",
					"application_server_address",
					"join_server_address",
				),
			})
			if err != nil {
				return err
			}

			compareServerAddressesApplication(application, config)

			requestDevEUI, _ := cmd.Flags().GetBool("request-dev-eui")
			if requestDevEUI {
				logger.Debug("request-dev-eui flag set, requesting a DevEUI")
				devEUIResponse, err := ttnpb.NewApplicationRegistryClient(is).IssueDevEUI(ctx, devID.ApplicationIds)
				if err != nil {
					return err
				}
				devEUI := types.MustEUI64(devEUIResponse.DevEui).OrZero()
				logger.WithField("dev_eui", devEUI.String()).
					Info("Successfully obtained DevEUI")
				device.Ids.DevEui = devEUI.Bytes()
			}
			newPaths, err := parsePayloadFormatterParameterFlags("formatters", device.Formatters, cmd.Flags())
			if err != nil {
				return err
			}
			paths = append(paths, newPaths...)

			if device.GetIds().GetApplicationIds().GetApplicationId() == "" {
				return errNoApplicationID.New()
			}
			if device.Ids.DeviceId == "" {
				return errNoEndDeviceID.New()
			}

			isPaths, nsPaths, asPaths, jsPaths := splitEndDeviceSetPaths(device.SupportsJoin, paths...)

			// If CAC is set, attempt to claim the End Device via the DCS instead of registering on the Join Server.
			if claimOnExternalJS {
				dcs, err := api.Dial(ctx, config.DeviceClaimingServerGRPCAddress)
				if err != nil {
					return err
				}
				claimInfoResp, err := ttnpb.NewEndDeviceClaimingServerClient(dcs).GetInfoByJoinEUI(ctx, &ttnpb.GetInfoByJoinEUIRequest{
					JoinEui: device.Ids.JoinEui,
				})
				if err != nil {
					return errEndDeviceClaimInfo.WithCause(err)
				}
				if claimInfoResp.SupportsClaiming {
					_, err = ttnpb.NewEndDeviceClaimingServerClient(dcs).Claim(ctx, &ttnpb.ClaimEndDeviceRequest{
						TargetApplicationIds: device.Ids.ApplicationIds,
						TargetDeviceId:       device.Ids.DeviceId,
						SourceDevice: &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers_{
							AuthenticatedIdentifiers: &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers{
								JoinEui:            device.Ids.JoinEui,
								DevEui:             device.Ids.DevEui,
								AuthenticationCode: device.ClaimAuthenticationCode.Value,
							},
						},
					})
					if err != nil {
						return errEndDeviceClaim.WithCause(err)
					}
					logger.Info("Device successfully claimed, skip registering on the cluster Join Server")
					// Remove Cluster Join Server related paths.
					jsPaths = []string{}
					isPaths = ttnpb.ExcludeFields(isPaths, "join_server_address")
					device.JoinServerAddress = ""
				} else {
					// TODO: Remove this path once the legacy DCS is deprecated (https://github.com/TheThingsIndustries/lorawan-stack/issues/3036).
					logger.Info("Device not configured for claiming, register in the cluster Join Server")
				}
			}
			// Require EUIs for devices that need to be added to the Join Server.
			if len(jsPaths) > 0 && (device.Ids.JoinEui == nil || device.Ids.DevEui == nil) {
				return errNoEndDeviceEUI.New()
			}
			isDevice := &ttnpb.EndDevice{}
			logger.WithField("paths", isPaths).Debug("Create end device on Identity Server")
			if err := isDevice.SetFields(device, append(isPaths, "ids")...); err != nil {
				return err
			}
			isRes, err := ttnpb.NewEndDeviceRegistryClient(is).Create(ctx, &ttnpb.CreateEndDeviceRequest{
				EndDevice: isDevice,
			})
			if err != nil {
				return err
			}

			if err := device.SetFields(isRes, append(isPaths, "created_at", "updated_at")...); err != nil {
				return err
			}

			res, err := setEndDevice(device, nil, nsPaths, asPaths, jsPaths, nil, true, false)
			if err != nil {
				logger.WithError(err).Error("Could not create end device, rolling back...")
				if err := deleteEndDevice(context.Background(), device.Ids, claimOnExternalJS); err != nil {
					logger.WithError(err).Error("Could not roll back end device creation")
				}
				return err
			}

			if err := device.SetFields(res, append(append(nsPaths, asPaths...), jsPaths...)...); err != nil {
				return err
			}
			if device.CreatedAt == nil || (res.CreatedAt != nil && ttnpb.StdTime(res.CreatedAt).Before(*ttnpb.StdTime(device.CreatedAt))) {
				device.CreatedAt = res.CreatedAt
			}
			if res.UpdatedAt != nil && ttnpb.StdTime(res.UpdatedAt).After(*ttnpb.StdTime(device.UpdatedAt)) {
				device.UpdatedAt = res.UpdatedAt
			}

			return io.Write(os.Stdout, config.OutputFormat, device)
		}),
	}
	endDevicesSetCommand = &cobra.Command{
		Use:     "set [application-id] [device-id]",
		Aliases: []string{"update"},
		Short:   "Set properties of an end device",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedDeviceFlags(cmd.Flags())

			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			paths := util.UpdateFieldMask(cmd.Flags(), setEndDeviceFlags, endDevicePictureFlags)
			rawUnsetPaths, _ := cmd.Flags().GetStringSlice("unset")
			unsetPaths := util.NormalizePaths(rawUnsetPaths)

			if hasUpdateDeviceLocationFlags(cmd.Flags()) {
				paths = append(paths, "locations")
			}

			if len(paths)+len(unsetPaths) == 0 {
				logger.Warn("No fields selected, won't update anything")
				return nil
			}
			if remainingPaths := ttnpb.ExcludeFields(paths, unsetPaths...); len(remainingPaths) != len(paths) {
				overlapPaths := ttnpb.ExcludeFields(paths, remainingPaths...)
				return errConflictingPaths.WithAttributes("field_mask_paths", overlapPaths)
			}
			device := &ttnpb.EndDevice{
				Ids: &ttnpb.EndDeviceIdentifiers{},
			}
			if ttnpb.HasAnyField(paths, setEndDeviceToJS...) || ttnpb.HasAnyField(unsetPaths, setEndDeviceToJS...) {
				device.SupportsJoin = true
			}
			_, err = device.SetFromFlags(setEndDeviceFlags, "")
			if err != nil {
				return err
			}
			newPaths, err := parsePayloadFormatterParameterFlags("formatters", device.Formatters, cmd.Flags())
			if err != nil {
				return err
			}
			paths = append(paths, newPaths...)
			device.Attributes = mergeAttributes(device.Attributes, cmd.Flags())
			device.Ids = devID

			paths = append(paths, unsetPaths...)
			isPaths, nsPaths, asPaths, jsPaths := splitEndDeviceSetPaths(device.SupportsJoin, paths...)
			if len(nsPaths) > 0 && config.NetworkServerEnabled {
				if device.NetworkServerAddress == "" {
					device.NetworkServerAddress = getHost(config.NetworkServerGRPCAddress)
				}
				isPaths = append(isPaths, "network_server_address")
			}
			if len(asPaths) > 0 && config.ApplicationServerEnabled {
				if device.ApplicationServerAddress == "" {
					device.ApplicationServerAddress = getHost(config.ApplicationServerGRPCAddress)
				}
				isPaths = append(isPaths, "application_server_address")
			}
			if len(jsPaths) > 0 && config.JoinServerEnabled {
				if device.JoinServerAddress == "" {
					device.JoinServerAddress = getHost(config.JoinServerGRPCAddress)
				}
				isPaths = append(isPaths, "join_server_address")
			}

			if picture, err := cmd.Flags().GetString("picture"); err == nil && picture != "" {
				device.Picture, err = readPicture(picture)
				if err != nil {
					return err
				}
				isPaths = append(paths, "picture")
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			logger.WithField("paths", isPaths).Debug("Get end device from Identity Server")

			// Always get the join server address to determine if the device uses an external Join Server.
			isGetPaths := ttnpb.AddFields(isPaths, "join_server_address")
			existingDevice, err := ttnpb.NewEndDeviceRegistryClient(is).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIds: devID,
				FieldMask:    ttnpb.FieldMask(ttnpb.ExcludeFields(isGetPaths, unsetPaths...)...),
			})
			if err != nil {
				return err
			}

			// EUIs can not be updated, so we only accept EUI flags if they are equal to the existing ones.
			if device.Ids.JoinEui != nil {
				if existingDevice.Ids.JoinEui != nil && !bytes.Equal(device.Ids.JoinEui, existingDevice.Ids.JoinEui) {
					return errEndDeviceEUIUpdate.New()
				}
			} else {
				device.Ids.JoinEui = existingDevice.Ids.JoinEui
			}
			if device.Ids.DevEui != nil {
				if existingDevice.Ids.DevEui != nil && !bytes.Equal(device.Ids.DevEui, existingDevice.Ids.DevEui) {
					return errEndDeviceEUIUpdate.New()
				}
			} else {
				device.Ids.DevEui = existingDevice.Ids.DevEui
			}

			// Require EUIs for devices that need to be updated in the Join Server.
			if len(jsPaths) > 0 && (device.Ids.JoinEui == nil || device.Ids.DevEui == nil) {
				return errNoEndDeviceEUI.New()
			}
			nsMismatch, asMismatch, jsMismatch := compareServerAddressesEndDevice(existingDevice, config)

			if nsMismatch || asMismatch {
				return errAddressMismatchEndDevice.New()
			}

			if len(jsPaths) > 0 && existingDevice.JoinServerAddress == "" {
				// End Device uses external Join Server. Disable dialing cluster Join Server.
				// If End Device claim needs to be updated, add those fields here and dial the DCS.
				logger.WithField("paths", jsPaths).Debug("End Device uses external Join Server, deselecting Join Server paths")
				jsPaths = []string{}
			} else if jsMismatch {
				return errAddressMismatchEndDevice.New()
			}

			if hasUpdateDeviceLocationFlags(cmd.Flags()) {
				if err := device.SetFields(existingDevice, "locations"); err != nil {
					return err
				}
				updateDeviceLocation(device, cmd.Flags())
			}

			touch, _ := cmd.Flags().GetBool("touch")
			res, err := setEndDevice(device, isPaths, nsPaths, asPaths, jsPaths, unsetPaths, false, touch)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	// TODO: Remove (https://github.com/TheThingsNetwork/lorawan-stack/issues/999)
	endDevicesProvisionCommand = &cobra.Command{
		Use:   "provision",
		Short: "Provision end devices using vendor-specific data",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Warn("This command is deprecated. Please use `device template from-data` instead")

			appID := getApplicationID(cmd.Flags(), nil)
			if appID == nil {
				return errNoApplicationID.New()
			}

			provisionerID, _ := cmd.Flags().GetString("provisioner-id")
			data, err := getDataBytes("", cmd.Flags())
			if err != nil {
				return err
			}

			req := &ttnpb.ProvisionEndDevicesRequest{
				ApplicationIds:   appID,
				ProvisionerId:    provisionerID,
				ProvisioningData: data,
			}

			var joinEUI types.EUI64
			if joinEUIHex, _ := cmd.Flags().GetString("join-eui"); joinEUIHex != "" {
				if err := joinEUI.UnmarshalText([]byte(joinEUIHex)); err != nil {
					return errInvalidJoinEUI.WithCause(err)
				}
			}
			if inputDecoder != nil {
				list := &ttnpb.ProvisionEndDevicesRequest_IdentifiersList{
					JoinEui: joinEUI.Bytes(),
				}
				for {
					var ids ttnpb.EndDeviceIdentifiers
					err := inputDecoder.Decode(&ids)
					if errors.Is(err, stdio.EOF) {
						break
					}
					if err != nil {
						return err
					}
					ids.ApplicationIds = appID
					list.EndDeviceIds = append(list.EndDeviceIds, &ids)
				}
				req.EndDevices = &ttnpb.ProvisionEndDevicesRequest_List{
					List: list,
				}
			} else {
				if startDevEUIHex, _ := cmd.Flags().GetString("start-dev-eui"); startDevEUIHex != "" {
					var startDevEUI types.EUI64
					if err := startDevEUI.UnmarshalText([]byte(startDevEUIHex)); err != nil {
						return errInvalidDevEUI.WithCause(err)
					}
					req.EndDevices = &ttnpb.ProvisionEndDevicesRequest_Range{
						Range: &ttnpb.ProvisionEndDevicesRequest_IdentifiersRange{
							StartDevEui: startDevEUI.Bytes(),
							JoinEui:     joinEUI.Bytes(),
						},
					}
				} else {
					req.EndDevices = &ttnpb.ProvisionEndDevicesRequest_FromData{
						FromData: &ttnpb.ProvisionEndDevicesRequest_IdentifiersFromData{
							JoinEui: joinEUI.Bytes(),
						},
					}
				}
			}

			js, err := api.Dial(ctx, config.JoinServerGRPCAddress)
			if err != nil {
				return err
			}
			stream, err := ttnpb.NewJsEndDeviceRegistryClient(js).Provision(ctx, req)
			if err != nil {
				return err
			}
			for {
				dev, err := stream.Recv()
				if errors.Is(err, stdio.EOF) {
					return nil
				}
				if err != nil {
					return err
				}
				if err := io.Write(os.Stdout, config.OutputFormat, dev); err != nil {
					return err
				}
			}
		},
	}
	endDevicesResetCommand = &cobra.Command{
		Use:   "reset [application-id] [device-id]",
		Short: "Reset state of an end device to factory defaults",
		RunE: func(cmd *cobra.Command, args []string) error {
			forwardDeprecatedDeviceFlags(cmd.Flags())

			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectEndDeviceFlags)

			isPaths, nsPaths, asPaths, jsPaths := splitEndDeviceGetPaths(paths...)

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			logger.WithField("paths", isPaths).Debug("Get end device from Identity Server")
			device, err := ttnpb.NewEndDeviceRegistryClient(is).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIds: devID,
				FieldMask:    ttnpb.FieldMask(isPaths...),
			})
			if err != nil {
				return err
			}

			nsMismatch, _, _ := compareServerAddressesEndDevice(device, config)
			if nsMismatch {
				return errors.New("Network Server address does not match")
			}

			ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
			if err != nil {
				return err
			}
			logger.WithField("paths", nsPaths).Debug("Reset end device to factory defaults on Network Server")
			nsDevice, err := ttnpb.NewNsEndDeviceRegistryClient(ns).ResetFactoryDefaults(ctx, &ttnpb.ResetAndGetEndDeviceRequest{
				EndDeviceIds: devID,
				FieldMask:    ttnpb.FieldMask(ttnpb.AddFields(nsPaths, "supports_join")...),
			})
			if err != nil {
				return err
			}
			if err = device.SetFields(nsDevice, "ids.dev_addr"); err != nil {
				return err
			}
			if err = device.SetFields(nsDevice, ttnpb.AllowedBottomLevelFields(nsPaths, getEndDeviceFromNS)...); err != nil {
				return err
			}
			device.UpdateTimestamps(nsDevice)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			logger.WithField("paths", asPaths).Debug("Reset end device to factory defaults on Application Server")
			asDevice, err := ttnpb.NewAsEndDeviceRegistryClient(as).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIds: devID,
				FieldMask:    ttnpb.FieldMask(asPaths...),
			})
			if err != nil {
				return err
			}
			var fieldsToReset []string
			if device.SupportsJoin {
				fieldsToReset = []string{"session", "pending_session"}
			} else {
				fieldsToReset = []string{"session.last_a_f_cnt_down"}
			}
			if err = asDevice.SetFields(nil, fieldsToReset...); err != nil {
				return err
			}
			_, err = ttnpb.NewAsEndDeviceRegistryClient(as).Set(ctx, &ttnpb.SetEndDeviceRequest{
				EndDevice: asDevice,
				FieldMask: ttnpb.FieldMask(asPaths...),
			})
			if err != nil {
				return err
			}
			if err := device.SetFields(asDevice, asPaths...); err != nil {
				return err
			}
			device.UpdateTimestamps(asDevice)

			if device.SupportsJoin {
				js, err := api.Dial(ctx, config.JoinServerGRPCAddress)
				if err != nil {
					return err
				}
				logger.WithField("paths", jsPaths).Debug("Reset end device to factory defaults on Join Server")
				jsDevice, err := ttnpb.NewJsEndDeviceRegistryClient(js).Get(ctx, &ttnpb.GetEndDeviceRequest{
					EndDeviceIds: devID,
					FieldMask:    ttnpb.FieldMask(jsPaths...),
				})
				if err != nil {
					return err
				}
				if err = jsDevice.SetFields(nil, "last_dev_nonce", "used_dev_nonces", "last_join_nonce", "last_rj_count_0", "last_rj_count_1"); err != nil {
					return err
				}
				_, err = ttnpb.NewJsEndDeviceRegistryClient(js).Set(ctx, &ttnpb.SetEndDeviceRequest{
					EndDevice: jsDevice,
					FieldMask: ttnpb.FieldMask(jsPaths...),
				})
				if err != nil {
					return err
				}
				if err := device.SetFields(jsDevice, jsPaths...); err != nil {
					return err
				}
				device.UpdateTimestamps(jsDevice)
			}

			// Remove temporary fields (e.g. "supports_join") that were not selected by user
			joinedPaths := ttnpb.AddFields(isPaths, ttnpb.AddFields(nsPaths, ttnpb.AddFields(asPaths, jsPaths...)...)...)
			if diff := ttnpb.ExcludeFields(joinedPaths, paths...); len(diff) > 0 {
				if err := device.SetFields(nil, diff...); err != nil {
					return err
				}
			}
			return io.Write(os.Stdout, config.OutputFormat, device)
		},
	}
	endDevicesDeleteCommand = &cobra.Command{
		Use:     "delete [application-id] [device-id]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete an end device",
		RunE: func(cmd *cobra.Command, args []string) error {
			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			existingDevice, err := ttnpb.NewEndDeviceRegistryClient(is).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIds: devID,
				FieldMask: ttnpb.FieldMask(
					"network_server_address",
					"application_server_address",
					"join_server_address",
				),
			})
			if err != nil {
				return err
			}

			// EUIs must match registered EUIs if set.
			if devID.JoinEui != nil {
				if existingDevice.Ids.JoinEui != nil && !bytes.Equal(devID.JoinEui, existingDevice.Ids.JoinEui) {
					return errInconsistentEndDeviceEUI.New()
				}
			} else {
				devID.JoinEui = existingDevice.Ids.JoinEui
			}
			if devID.DevEui != nil {
				if existingDevice.Ids.DevEui != nil && !bytes.Equal(devID.DevEui, existingDevice.Ids.DevEui) {
					return errInconsistentEndDeviceEUI.New()
				}
			} else {
				devID.DevEui = existingDevice.Ids.DevEui
			}

			var skipClusterJS bool
			nsMismatch, asMismatch, jsMismatch := compareServerAddressesEndDevice(existingDevice, config)

			if nsMismatch || asMismatch {
				return errAddressMismatchEndDevice.New()
			}

			if existingDevice.JoinServerAddress == "" && devID.GetJoinEui() != nil {
				// Attempt to unclaim device via the DCS.
				dcs, err := api.Dial(ctx, config.DeviceClaimingServerGRPCAddress)
				if err != nil {
					return err
				}
				claimInfoResp, err := ttnpb.NewEndDeviceClaimingServerClient(dcs).GetInfoByJoinEUI(ctx, &ttnpb.GetInfoByJoinEUIRequest{
					JoinEui: devID.JoinEui,
				})
				if err != nil {
					return errEndDeviceClaimInfo.WithCause(err)
				}
				if claimInfoResp.SupportsClaiming {
					_, err = ttnpb.NewEndDeviceClaimingServerClient(dcs).Unclaim(ctx, devID)
					if err != nil {
						logger.WithError(err).Warn("Failed to unclaim end device")
					} else {
						logger.Info("Device successfully unclaimed")
					}
					skipClusterJS = true
				}
			} else if jsMismatch {
				// Check if there's an address mismatch only if using the cluster Join Server.
				return errAddressMismatchEndDevice.New()
			}

			return deleteEndDevice(ctx, devID, skipClusterJS)
		},
	}
	endDevicesClaimCommand = &cobra.Command{
		Use:   "claim [application-id]",
		Short: "Claim an end device (EXPERIMENTAL)",
		Long: `Claim an end device (EXPERIMENTAL)

The claiming procedure transfers devices from the source application to the
target application using the Device Claiming Server, thereby transferring
ownership of the device.

Authentication of device claiming is by the device's JoinEUI, DevEUI and claim
authentication code as stored in the Join Server. This information is typically
encoded in a QR code. This command supports claiming by QR code (via stdin), as
well as providing the claim information through the flags --source-join-eui,
--source-dev-eui, --source-authentication-code.

Claim authentication code validity is controlled by the owner of the device by
setting the value and optionally a time window when the code is valid. As part
of the claiming, the claim authentication code is invalidated by default to
block subsequent claiming attempts. You can keep the claim authentication code
valid by specifying --invalidate-authentication-code=false.

As part of claiming, you can optionally provide the target NetID, Network Server
KEK label and Application Server ID and KEK label. The Network Server and
Application Server addresses will be taken from the CLI configuration. These
values will be stored in the Join Server.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			targetAppID := getApplicationID(cmd.Flags(), args)
			if targetAppID == nil {
				return errNoApplicationID.New()
			}

			req := &ttnpb.ClaimEndDeviceRequest{
				TargetApplicationIds: targetAppID,
			}

			var joinEUI, devEUI *types.EUI64
			if joinEUIHex, _ := cmd.Flags().GetString("source-join-eui"); joinEUIHex != "" {
				joinEUI = new(types.EUI64)
				if err := joinEUI.UnmarshalText([]byte(joinEUIHex)); err != nil {
					return errInvalidJoinEUI.WithCause(err)
				}
			}
			if devEUIHex, _ := cmd.Flags().GetString("source-dev-eui"); devEUIHex != "" {
				devEUI = new(types.EUI64)
				if err := devEUI.UnmarshalText([]byte(devEUIHex)); err != nil {
					return errInvalidDevEUI.WithCause(err)
				}
			}
			if joinEUI != nil && devEUI != nil {
				authenticationCode, _ := cmd.Flags().GetString("source-authentication-code")
				req.SourceDevice = &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers_{
					AuthenticatedIdentifiers: &ttnpb.ClaimEndDeviceRequest_AuthenticatedIdentifiers{
						JoinEui:            joinEUI.Bytes(),
						DevEui:             devEUI.Bytes(),
						AuthenticationCode: authenticationCode,
					},
				}
			} else {
				if joinEUI != nil || devEUI != nil {
					logger.Warn("Either target JoinEUI or DevEUI specified but need both, not considering any and using scan mode")
				}
				rd, ok := io.BufferedPipe(os.Stdin)
				if !ok {
					logger.Info("Scan QR code")
					rd = bufio.NewReader(os.Stdin)
				}
				qrCode, err := rd.ReadBytes('\n')
				if err != nil {
					return err
				}
				qrCode = qrCode[:len(qrCode)-1]
				logger.WithField("code", string(qrCode)).Debug("Scanned QR code")
				req.SourceDevice = &ttnpb.ClaimEndDeviceRequest_QrCode{
					QrCode: qrCode,
				}
			}

			req.TargetDeviceId, _ = cmd.Flags().GetString("target-device-id")
			if netIDHex, _ := cmd.Flags().GetString("target-net-id"); netIDHex != "" {
				var netID types.NetID
				if err := netID.UnmarshalText([]byte(netIDHex)); err != nil {
					return err
				}
				req.TargetNetId = netID.Bytes()
			}
			if config.NetworkServerEnabled {
				req.TargetNetworkServerAddress = config.NetworkServerGRPCAddress
			}
			req.TargetNetworkServerKekLabel, _ = cmd.Flags().GetString("target-network-server-kek-label")
			if config.ApplicationServerEnabled {
				req.TargetApplicationServerAddress = config.ApplicationServerGRPCAddress
			}
			req.TargetApplicationServerKekLabel, _ = cmd.Flags().GetString("target-application-server-kek-label")
			req.TargetApplicationServerId, _ = cmd.Flags().GetString("target-application-server-id")
			req.InvalidateAuthenticationCode, _ = cmd.Flags().GetBool("invalidate-authentication-code")

			dcs, err := api.Dial(ctx, config.DeviceClaimingServerGRPCAddress)
			if err != nil {
				return err
			}
			ids, err := ttnpb.NewEndDeviceClaimingServerClient(dcs).Claim(ctx, req)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, ids)
		},
	}
	endDevicesListQRCodeFormatsCommand = &cobra.Command{
		Use:     "list-qr-formats",
		Aliases: []string{"ls-qr-formats", "listqrformats", "lsqrformats", "lsqrfmts", "lsqrfmt", "qr-formats"},
		Short:   "List QR code formats (EXPERIMENTAL)",
		RunE: func(cmd *cobra.Command, args []string) error {
			qrg, err := api.Dial(ctx, config.QRCodeGeneratorGRPCAddress)
			if err != nil {
				return err
			}

			res, err := ttnpb.NewEndDeviceQRCodeGeneratorClient(qrg).ListFormats(ctx, ttnpb.Empty)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	endDevicesGenerateQRCommand = &cobra.Command{
		Use:     "generate-qr [application-id] [device-id]",
		Aliases: []string{"genqr"},
		Short:   "Generate an end device QR code (EXPERIMENTAL)",
		Long: `Generate an end device QR code (EXPERIMENTAL)

This command saves a QR code in PNG format in the given folder. The filename is
the device ID.

This command may take end device identifiers from stdin.`,
		Example: `
  To generate a QR code for a single end device:
    $ ttn-lw-cli end-devices generate-qr app1 dev1

  To generate a QR code for multiple end devices:
    $ ttn-lw-cli end-devices list app1 \
      | ttn-lw-cli end-devices generate-qr`,
		RunE: asBulk(func(cmd *cobra.Command, args []string) error {
			var ids *ttnpb.EndDeviceIdentifiers
			if inputDecoder != nil {
				var dev ttnpb.EndDevice
				if err := inputDecoder.Decode(&dev); err != nil {
					return err
				}
				if dev.GetIds().GetApplicationIds().GetApplicationId() == "" {
					return errNoApplicationID.New()
				}
				if dev.Ids.DeviceId == "" {
					return errNoEndDeviceID.New()
				}
				ids = dev.Ids
			} else {
				var err error
				ids, err = getEndDeviceID(cmd.Flags(), args, true)
				if err != nil {
					return err
				}
			}

			formatID, _ := cmd.Flags().GetString("format-id")

			qrg, err := api.Dial(ctx, config.QRCodeGeneratorGRPCAddress)
			if err != nil {
				return err
			}
			client := ttnpb.NewEndDeviceQRCodeGeneratorClient(qrg)
			format, err := client.GetFormat(ctx, &ttnpb.GetQRCodeFormatRequest{
				FormatId: formatID,
			})
			if err != nil {
				return err
			}

			isPaths, nsPaths, asPaths, jsPaths := splitEndDeviceGetPaths(format.FieldMask.GetPaths()...)
			if len(nsPaths) > 0 {
				isPaths = append(isPaths, "network_server_address")
			}
			if len(asPaths) > 0 {
				isPaths = append(isPaths, "application_server_address")
			}
			if len(jsPaths) > 0 {
				isPaths = append(isPaths, "join_server_address")
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			logger.WithField("paths", isPaths).Debug("Get end device from Identity Server")
			device, err := ttnpb.NewEndDeviceRegistryClient(is).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIds: ids,
				FieldMask:    ttnpb.FieldMask(isPaths...),
			})
			if err != nil {
				return err
			}

			if device.ClaimAuthenticationCode.GetValue() != "" {
				// ClaimAuthenticationCode is already retrieved from the IS. We can unset the related JS paths.
				jsPaths = ttnpb.ExcludeFields(jsPaths, claimAuthenticationCodePaths...)
			}

			nsMismatch, asMismatch, jsMismatch := compareServerAddressesEndDevice(device, config)
			if len(nsPaths) > 0 && nsMismatch {
				return errAddressMismatchEndDevice.New()
			}
			if len(asPaths) > 0 && asMismatch {
				return errAddressMismatchEndDevice.New()
			}
			if len(jsPaths) > 0 && jsMismatch {
				return errAddressMismatchEndDevice.New()
			}
			dev, err := getEndDevice(device.Ids, nsPaths, asPaths, jsPaths, true)
			if err != nil {
				return err
			}
			if err := device.SetFields(dev, append(append(nsPaths, asPaths...), jsPaths...)...); err != nil {
				return err
			}
			size, _ := cmd.Flags().GetUint32("size")
			res, err := client.Generate(ctx, &ttnpb.GenerateEndDeviceQRCodeRequest{
				FormatId:  formatID,
				EndDevice: device,
				Image: &ttnpb.GenerateEndDeviceQRCodeRequest_Image{
					ImageSize: size,
				},
			})
			if err != nil {
				return err
			}

			folder, _ := cmd.Flags().GetString("folder")
			if folder == "" {
				folder, err = os.Getwd()
				if err != nil {
					return err
				}
			}

			var ext string
			if exts, err := mime.ExtensionsByType(res.Image.Embedded.MimeType); err == nil && len(exts) > 0 {
				ext = exts[0]
			}
			filename := path.Join(folder, device.Ids.DeviceId+ext)
			if err := os.WriteFile(filename, res.Image.Embedded.Data, 0o644); err != nil { //nolint:gas
				return err
			}

			logger.WithFields(log.Fields(
				"value", res.Text,
				"filename", filename,
			)).Info("Generated QR code")
			return nil
		}),
	}
	endDevicesExternalJSCommand = &cobra.Command{
		Use:     "use-external-join-server [application-id] [device-id]",
		Aliases: []string{"use-external-js", "use-ext-js"},
		Short:   "Disassociate and delete the device from Join Server",
		RunE: func(cmd *cobra.Command, args []string) error {
			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			if !config.JoinServerEnabled {
				return errJoinServerDisabled.New()
			}

			is, err := api.Dial(ctx, config.IdentityServerGRPCAddress)
			if err != nil {
				return err
			}
			dev, err := ttnpb.NewEndDeviceRegistryClient(is).Get(ctx, &ttnpb.GetEndDeviceRequest{
				EndDeviceIds: devID,
				FieldMask:    ttnpb.FieldMask("join_server_address"),
			})
			if err != nil {
				return err
			}
			if _, _, nok := compareServerAddressesEndDevice(dev, config); nok {
				return errAddressMismatchEndDevice.New()
			}

			js, err := api.Dial(ctx, config.JoinServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewJsEndDeviceRegistryClient(js).Delete(ctx, devID)
			if err != nil {
				return err
			}

			_, err = ttnpb.NewEndDeviceRegistryClient(is).Update(ctx, &ttnpb.UpdateEndDeviceRequest{
				EndDevice: &ttnpb.EndDevice{
					Ids: devID,
				},
				FieldMask: ttnpb.FieldMask("join_server_address"),
			})
			return err
		},
	}
	endDevicesGetDefaultMACSettingsCommand = &cobra.Command{
		Use:               "get-default-mac-settings",
		Short:             "Get Network Server default MAC settings for frequency plan and LoRaWAN version",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, args []string) error {
			if !config.NetworkServerEnabled {
				return errNetworkServerDisabled.New()
			}

			req := &ttnpb.GetDefaultMACSettingsRequest{}
			_, err := req.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewNsClient(ns).GetDefaultMACSettings(ctx, req)
			if err != nil {
				return err
			}
			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	endDevicesGetNetIDCommand = &cobra.Command{
		Use:               "get-net-id",
		Short:             "Get Network Server configured Net ID",
		PersistentPreRunE: preRun(),
		RunE: func(_ *cobra.Command, _ []string) error {
			if !config.NetworkServerEnabled {
				return errNetworkServerDisabled.New()
			}

			ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewNsClient(ns).GetNetID(ctx, ttnpb.Empty)
			if err != nil {
				return err
			}
			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	endDevicesGetDevAddrPrefixesCommand = &cobra.Command{
		Use:               "get-dev-addr-prefixes",
		Short:             "Get Network Server configured device address prefixes",
		PersistentPreRunE: preRun(),
		RunE: func(_ *cobra.Command, _ []string) error {
			if !config.NetworkServerEnabled {
				return errNetworkServerDisabled.New()
			}

			ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewNsClient(ns).GetDeviceAddressPrefixes(ctx, ttnpb.Empty)
			if err != nil {
				return err
			}
			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	endDevicesListBandsCommand = &cobra.Command{
		Use:               "list-bands",
		Short:             "List available band definitions",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !config.NetworkServerEnabled {
				return errNetworkServerDisabled.New()
			}

			req := &ttnpb.ListBandsRequest{}
			_, err := req.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewConfigurationClient(ns).ListBands(ctx, req)
			if err != nil {
				return err
			}
			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	endDevicesListPhyVersionsCommand = &cobra.Command{
		Use:               "list-phy-versions",
		Aliases:           []string{"get-phy-versions"},
		Short:             "List supported phy versions",
		PersistentPreRunE: preRun(),
		RunE: func(cmd *cobra.Command, _ []string) error {
			if !config.NetworkServerEnabled {
				return errNetworkServerDisabled.New()
			}

			req := &ttnpb.GetPhyVersionsRequest{}
			_, err := req.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			ns, err := api.Dial(ctx, config.NetworkServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewConfigurationClient(ns).GetPhyVersions(ctx, req)
			if err != nil {
				return err
			}
			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
)

func init() {
	ttnpb.AddSetFlagsForLocation(endDeviceLocationFlags, "location", false)
	ttnpb.AddSetFlagsForGetDefaultMACSettingsRequest(getDefaultMACSettingsFlags, "", false)
	ttnpb.AddSelectFlagsForEndDevice(allEndDeviceSelectFlags, "", false)
	ttnpb.AddSetFlagsForEndDevice(allEndDeviceSetFlags, "", false)
	ttnpb.AddSetFlagsForListBandsRequest(listBandsFlags, "", false)
	ttnpb.AddSetFlagsForGetPhyVersionsRequest(listPhyVersionFlags, "", false)

	allEndDeviceSelectFlags.VisitAll(func(flag *pflag.Flag) {
		fieldName := toUnderscore.Replace(flag.Name)
		f1 := *flag
		f2 := *flag
		selectEndDeviceListFlags.AddFlag(&f1)
		selectEndDeviceFlags.AddFlag(&f2)
		if !ttnpb.ContainsField(fieldName, getEndDeviceFromIS) {
			util.HideFlag(selectEndDeviceListFlags, flag.Name)
			if !ttnpb.ContainsField(fieldName, getEndDeviceFromNS) &&
				!ttnpb.ContainsField(fieldName, getEndDeviceFromAS) &&
				!ttnpb.ContainsField(fieldName, getEndDeviceFromJS) {
				util.HideFlag(selectEndDeviceFlags, flag.Name)
			}
		}
	})

	addDeprecatedDeviceFlags(selectEndDeviceListFlags)
	addDeprecatedDeviceFlags(selectEndDeviceFlags)

	allEndDeviceSetFlags.VisitAll(func(flag *pflag.Flag) {
		fieldName := toUnderscore.Replace(flag.Name)
		setEndDeviceFlags.AddFlag(flag)
		if !ttnpb.ContainsField(fieldName, setEndDeviceToIS) &&
			!ttnpb.ContainsField(fieldName, setEndDeviceToNS) &&
			!ttnpb.ContainsField(fieldName, setEndDeviceToAS) &&
			!ttnpb.ContainsField(fieldName, setEndDeviceToJS) {
			util.HideFlag(setEndDeviceFlags, flag.Name)
		}
	})

	addDeprecatedDeviceFlags(setEndDeviceFlags)

	endDevicePictureFlags.String("picture", "", "upload the end device picture from this file")

	endDevicesListFrequencyPlans.Flags().Uint32("base-frequency", 0, "base frequency in MHz for hardware support (433, 470, 868 or 915)")
	endDevicesCommand.AddCommand(endDevicesListFrequencyPlans)
	endDevicesListCommand.Flags().AddFlagSet(applicationIDFlags())
	endDevicesListCommand.Flags().AddFlagSet(selectEndDeviceListFlags)
	endDevicesListCommand.Flags().AddFlagSet(selectAllEndDeviceFlags)
	endDevicesListCommand.Flags().AddFlagSet(paginationFlags())
	endDevicesListCommand.Flags().AddFlagSet(orderFlags())
	endDevicesCommand.AddCommand(endDevicesListCommand)
	ttnpb.AddSetFlagsForSearchEndDevicesRequest(endDevicesSearchCommand.Flags(), "", false)
	endDevicesSearchCommand.Flags().AddFlagSet(selectEndDeviceFlags)
	endDevicesSearchCommand.Flags().AddFlagSet(selectAllEndDeviceFlags)
	endDevicesCommand.AddCommand(endDevicesSearchCommand)
	endDevicesGetCommand.Flags().AddFlagSet(endDeviceIDFlags())
	endDevicesGetCommand.Flags().AddFlagSet(selectEndDeviceFlags)
	endDevicesGetCommand.Flags().AddFlagSet(selectAllEndDeviceFlags)
	endDevicesCommand.AddCommand(endDevicesGetCommand)
	endDevicesCreateCommand.Flags().AddFlagSet(endDeviceIDFlags())
	endDevicesCreateCommand.Flags().AddFlagSet(setEndDeviceFlags)
	endDevicesCreateCommand.Flags().AddFlagSet(payloadFormatterParameterFlags("formatters"))
	endDevicesCreateCommand.Flags().Bool("defaults", true, "configure end device with defaults")
	endDevicesCreateCommand.Flags().Bool("with-root-keys", false, "generate OTAA root keys")
	endDevicesCreateCommand.Flags().Bool("abp", false, "configure end device as ABP")
	endDevicesCreateCommand.Flags().Bool("with-session", false, "generate ABP session DevAddr and keys")
	endDevicesCreateCommand.Flags().Bool("with-claim-authentication-code", false, "generate claim authentication code of 4 bytes")
	endDevicesCreateCommand.Flags().Bool("request-dev-eui", false, "request a new DevEUI")
	endDevicesCreateCommand.Flags().AddFlagSet(endDevicePictureFlags)
	endDevicesCreateCommand.Flags().AddFlagSet(endDeviceLocationFlags)
	endDevicesCommand.AddCommand(endDevicesCreateCommand)
	endDevicesSetCommand.Flags().AddFlagSet(endDeviceIDFlags())
	endDevicesSetCommand.Flags().AddFlagSet(setEndDeviceFlags)
	endDevicesSetCommand.Flags().AddFlagSet(payloadFormatterParameterFlags("formatters"))
	endDevicesSetCommand.Flags().Bool("touch", false, "set in all registries even if no fields are specified")
	endDevicesSetCommand.Flags().AddFlagSet(endDevicePictureFlags)
	endDevicesSetCommand.Flags().AddFlagSet(endDeviceLocationFlags)
	endDevicesSetCommand.Flags().AddFlagSet(util.UnsetFlagSet())
	endDevicesCommand.AddCommand(endDevicesSetCommand)
	endDevicesProvisionCommand.Flags().AddFlagSet(applicationIDFlags())
	endDevicesProvisionCommand.Flags().AddFlagSet(dataFlags("", ""))
	endDevicesProvisionCommand.Flags().String("provisioner-id", "", "provisioner service")
	endDevicesProvisionCommand.Flags().String("join-eui", "", "(hex)")
	endDevicesProvisionCommand.Flags().String("start-dev-eui", "", "starting DevEUI to provision (hex)")
	endDevicesCommand.AddCommand(endDevicesProvisionCommand)
	endDevicesResetCommand.Flags().AddFlagSet(endDeviceIDFlags())
	endDevicesResetCommand.Flags().AddFlagSet(selectEndDeviceFlags)
	endDevicesResetCommand.Flags().AddFlagSet(selectAllEndDeviceFlags)
	endDevicesCommand.AddCommand(endDevicesResetCommand)
	endDevicesDeleteCommand.Flags().AddFlagSet(endDeviceIDFlags())
	endDevicesCommand.AddCommand(endDevicesDeleteCommand)
	endDevicesClaimCommand.Flags().AddFlagSet(applicationIDFlags())
	endDevicesClaimCommand.Flags().String("source-join-eui", "", "(hex)")
	endDevicesClaimCommand.Flags().String("source-dev-eui", "", "(hex)")
	endDevicesClaimCommand.Flags().String("source-authentication-code", "", "(hex)")
	endDevicesClaimCommand.Flags().String("target-device-id", "", "")
	endDevicesClaimCommand.Flags().String("target-net-id", "", "(hex)")
	endDevicesClaimCommand.Flags().String("target-network-server-kek-label", "", "")
	endDevicesClaimCommand.Flags().String("target-application-server-kek-label", "", "")
	endDevicesClaimCommand.Flags().String("target-application-server-id", "", "")
	endDevicesClaimCommand.Flags().Bool("invalidate-authentication-code", true, "invalidate the claim authentication code to block subsequent claiming attempts")
	endDevicesCommand.AddCommand(endDevicesClaimCommand)
	endDevicesCommand.AddCommand(endDevicesListQRCodeFormatsCommand)
	endDevicesGenerateQRCommand.Flags().AddFlagSet(endDeviceIDFlags())
	endDevicesGenerateQRCommand.Flags().String("format-id", "", "")
	endDevicesGenerateQRCommand.Flags().Uint32("size", 300, "size of the image in pixels")
	endDevicesGenerateQRCommand.Flags().String("folder", "", "folder to write the QR code image to")
	endDevicesCommand.AddCommand(endDevicesGenerateQRCommand)
	endDevicesExternalJSCommand.Flags().AddFlagSet(endDeviceIDFlags())
	endDevicesCommand.AddCommand(endDevicesExternalJSCommand)

	endDevicesCommand.AddCommand(applicationsDownlinkCommand)

	endDevicesGetDefaultMACSettingsCommand.Flags().AddFlagSet(getDefaultMACSettingsFlags)
	endDevicesCommand.AddCommand(endDevicesGetDefaultMACSettingsCommand)

	endDevicesGetNetIDCommand.Flags().AddFlagSet(getNetIDFlags)
	endDevicesCommand.AddCommand(endDevicesGetNetIDCommand)

	endDevicesGetDevAddrPrefixesCommand.Flags().AddFlagSet(getDevAddrPrefixesFlags)
	endDevicesCommand.AddCommand(endDevicesGetDevAddrPrefixesCommand)

	endDevicesListBandsCommand.Flags().AddFlagSet(listBandsFlags)
	endDevicesCommand.AddCommand(endDevicesListBandsCommand)

	endDevicesListPhyVersionsCommand.Flags().AddFlagSet(listPhyVersionFlags)
	endDevicesCommand.AddCommand(endDevicesListPhyVersionsCommand)

	Root.AddCommand(endDevicesCommand)

	endDeviceTemplatesExtendCommand.Flags().AddFlagSet(setEndDeviceFlags)
	endDeviceTemplatesCreateCommand.Flags().AddFlagSet(selectEndDeviceFlags)
	endDeviceTemplatesExecuteCommand.Flags().AddFlagSet(setEndDeviceFlags)
}

var errAddressMismatchEndDevice = errors.DefineAborted("end_device_server_address_mismatch", "Network/Application/Join Server address mismatch")

func compareServerAddressesEndDevice(device *ttnpb.EndDevice, config *Config) (nsMismatch, asMismatch, jsMismatch bool) {
	nsHost, asHost, jsHost := getHost(config.NetworkServerGRPCAddress), getHost(config.ApplicationServerGRPCAddress), getHost(config.JoinServerGRPCAddress)
	if host := getHost(device.NetworkServerAddress); config.NetworkServerEnabled && host != "" && host != nsHost {
		nsMismatch = true
		logger.WithFields(log.Fields(
			"configured", nsHost,
			"registered", host,
		)).Warnf("Registered Network Server address of end device %q does not match CLI configuration", device.GetIds().GetDeviceId())
	}
	if host := getHost(device.ApplicationServerAddress); config.ApplicationServerEnabled && host != "" && host != asHost {
		asMismatch = true
		logger.WithFields(log.Fields(
			"configured", asHost,
			"registered", host,
		)).Warnf("Registered Application Server address of end device %q does not match CLI configuration", device.GetIds().GetDeviceId())
	}
	if host := getHost(device.JoinServerAddress); config.JoinServerEnabled && host != "" && host != jsHost {
		jsMismatch = true
		logger.WithFields(log.Fields(
			"configured", jsHost,
			"registered", host,
		)).Warnf("Registered Join Server address of end device %q does not match CLI configuration", device.GetIds().GetDeviceId())
	}
	return
}
