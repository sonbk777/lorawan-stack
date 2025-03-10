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
	"os"
	"strconv"
	"strings"

	"github.com/TheThingsIndustries/protoc-gen-go-flags/flagsplugin"
	pbtypes "github.com/gogo/protobuf/types"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"go.thethings.network/lorawan-stack/v3/cmd/internal/io"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/api"
	"go.thethings.network/lorawan-stack/v3/cmd/ttn-lw-cli/internal/util"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
)

var (
	selectApplicationPackageAssociationsFlags        = util.NormalizedFlagSet()
	selectApplicationPackageDefaultAssociationsFlags = util.NormalizedFlagSet()

	selectAllApplicationPackageAssociationsFlags        = util.SelectAllFlagSet("application package association")
	selectAllApplicationPackageDefaultAssociationsFlags = util.SelectAllFlagSet("application package default association")

	packageNeedsData = map[string]struct{}{
		"lora-cloud-device-management-v1": {},
	}
)

func applicationPackageAssociationIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("application-id", "", "")
	flagSet.String("device-id", "", "")
	flagSet.Uint8("f-port", 0, "")
	return flagSet
}

func applicationPackageDefaultAssociationIDFlags() *pflag.FlagSet {
	flagSet := &pflag.FlagSet{}
	flagSet.String("application-id", "", "")
	flagSet.Uint8("f-port", 0, "")
	return flagSet
}

var errNoFPort = errors.DefineInvalidArgument("no_f_port", "no FPort set")

func getApplicationPackageAssociationID(flagSet *pflag.FlagSet, args []string) (*ttnpb.ApplicationPackageAssociationIdentifiers, error) {
	applicationID, _ := flagSet.GetString("application-id")
	deviceID, _ := flagSet.GetString("device-id")
	fport, _ := flagSet.GetUint8("f-port")
	switch len(args) {
	case 0:
	case 1:
	case 2:
		logger.Warn("Only single ID found in arguments, not considering arguments")
	case 3:
		applicationID = args[0]
		deviceID = args[1]
		fport64, err := strconv.ParseUint(args[2], 10, 8)
		if err != nil {
			return nil, err
		}
		fport = uint8(fport64)
	default:
		logger.Warn("multiple IDs found in arguments, considering the first")
		applicationID = args[0]
		deviceID = args[1]
		fport64, err := strconv.ParseUint(args[2], 10, 8)
		if err != nil {
			return nil, err
		}
		fport = uint8(fport64)
	}
	if applicationID == "" {
		return nil, errNoApplicationID.New()
	}
	if deviceID == "" {
		return nil, errNoEndDeviceID.New()
	}
	if fport == 0 {
		return nil, errNoFPort.New()
	}
	return &ttnpb.ApplicationPackageAssociationIdentifiers{
		EndDeviceIds: &ttnpb.EndDeviceIdentifiers{
			ApplicationIds: &ttnpb.ApplicationIdentifiers{ApplicationId: applicationID},
			DeviceId:       deviceID,
		},
		FPort: uint32(fport),
	}, nil
}

func getApplicationPackageDefaultAssociationID(flagSet *pflag.FlagSet, args []string) (*ttnpb.ApplicationPackageDefaultAssociationIdentifiers, error) {
	applicationID, _ := flagSet.GetString("application-id")
	fport, _ := flagSet.GetUint8("f-port")
	switch len(args) {
	case 0:
	case 1:
		logger.Warn("Only single ID found in arguments, not considering arguments")
	case 2:
		applicationID = args[0]
		fport64, err := strconv.ParseUint(args[1], 10, 8)
		if err != nil {
			return nil, err
		}
		fport = uint8(fport64)
	default:
		logger.Warn("multiple IDs found in arguments, considering the first")
		applicationID = args[0]
		fport64, err := strconv.ParseUint(args[1], 10, 8)
		if err != nil {
			return nil, err
		}
		fport = uint8(fport64)
	}
	if applicationID == "" {
		return nil, errNoApplicationID.New()
	}
	if fport == 0 {
		return nil, errNoFPort.New()
	}
	return &ttnpb.ApplicationPackageDefaultAssociationIdentifiers{
		ApplicationIds: &ttnpb.ApplicationIdentifiers{ApplicationId: applicationID},
		FPort:          uint32(fport),
	}, nil
}

var (
	applicationsPackagesCommand = &cobra.Command{
		Use:     "packages",
		Aliases: []string{"package", "pkg", "pkgs"},
		Short:   "Application packages commands",
	}
	applicationsPackagesListCommand = &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the available application packages for the device",
		RunE: func(cmd *cobra.Command, args []string) error {
			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).List(ctx, devID)
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackagesAssociationsCommand = &cobra.Command{
		Use:     "associations",
		Aliases: []string{"assoc", "assocs"},
		Short:   "Application packages associations commands",
	}
	applicationsPackageAssociationGetCommand = &cobra.Command{
		Use:     "get [application-id] [device-id] [f-port]",
		Aliases: []string{"info"},
		Short:   "Get the properties of an application package association",
		RunE: func(cmd *cobra.Command, args []string) error {
			assocID, err := getApplicationPackageAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPackageAssociationsFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPackageAssociationsFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationPackageRegistry/GetAssociation"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).GetAssociation(ctx, &ttnpb.GetApplicationPackageAssociationRequest{
				Ids:       assocID,
				FieldMask: ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackageAssociationsListCommand = &cobra.Command{
		Use:     "list [application-id] [device-id]",
		Aliases: []string{"ls"},
		Short:   "List application package associations",
		RunE: func(cmd *cobra.Command, args []string) error {
			devID, err := getEndDeviceID(cmd.Flags(), args, true)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPackageAssociationsFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPackageAssociationsFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationPackageRegistry/ListAssociations"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			limit, page, opt, getTotal := withPagination(cmd.Flags())
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).ListAssociations(ctx, &ttnpb.ListApplicationPackageAssociationRequest{
				Ids:       devID,
				Limit:     limit,
				Page:      page,
				FieldMask: ttnpb.FieldMask(paths...),
			}, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackageAssociationSetCommand = &cobra.Command{
		Use:     "set [application-id] [device-id] [f-port]",
		Aliases: []string{"update"},
		Short:   "Set the properties of an application package association",
		RunE: func(cmd *cobra.Command, args []string) error {
			association := &ttnpb.ApplicationPackageAssociation{}
			paths, err := association.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			assocID, err := getApplicationPackageAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			association.Ids = assocID

			reader, err := getDataReader("data", cmd.Flags())
			if err != nil {
				if _, needsData := packageNeedsData[association.PackageName]; needsData {
					logger.WithError(err).Warn("Package data not available")
				}
			} else {
				var st pbtypes.Struct
				err := jsonpb.TTN().NewDecoder(reader).Decode(&st)
				if err != nil {
					return err
				}

				association.Data = &st
				paths = append(paths, "data")
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).SetAssociation(ctx, &ttnpb.SetApplicationPackageAssociationRequest{
				Association: association,
				FieldMask:   ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackageAssociationDeleteCommand = &cobra.Command{
		Use:     "delete [application-id] [device-id] [f-port]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete an application package association",
		RunE: func(cmd *cobra.Command, args []string) error {
			assocID, err := getApplicationPackageAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewApplicationPackageRegistryClient(as).DeleteAssociation(ctx, assocID)
			if err != nil {
				return err
			}

			return nil
		},
	}
	applicationsPackagesDefaultAssociationsCommand = &cobra.Command{
		Use:     "default-associations",
		Aliases: []string{"def-assoc", "def-assocs"},
		Short:   "Application packages default associations commands",
	}
	applicationsPackageDefaultAssociationGetCommand = &cobra.Command{
		Use:     "get [application-id] [f-port]",
		Aliases: []string{"info"},
		Short:   "Get the properties of an application package default association",
		RunE: func(cmd *cobra.Command, args []string) error {
			assocID, err := getApplicationPackageDefaultAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPackageDefaultAssociationsFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPackageDefaultAssociationsFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationPackageRegistry/GetDefaultAssociation"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).GetDefaultAssociation(ctx, &ttnpb.GetApplicationPackageDefaultAssociationRequest{
				Ids:       assocID,
				FieldMask: ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackageDefaultAssociationsListCommand = &cobra.Command{
		Use:     "list [application-id]",
		Aliases: []string{"ls"},
		Short:   "List application package default associations",
		RunE: func(cmd *cobra.Command, args []string) error {
			appID := getApplicationID(cmd.Flags(), args)
			if appID == nil {
				return errNoApplicationID.New()
			}
			paths := util.SelectFieldMask(cmd.Flags(), selectApplicationPackageDefaultAssociationsFlags)
			if len(paths) == 0 {
				logger.Warn("No fields selected, will select everything")
				selectApplicationPackageDefaultAssociationsFlags.VisitAll(func(flag *pflag.Flag) {
					paths = append(paths, strings.Replace(flag.Name, "-", "_", -1))
				})
			}
			paths = ttnpb.AllowedFields(paths, ttnpb.RPCFieldMaskPaths["/ttn.lorawan.v3.ApplicationPackageRegistry/ListDefaultAssociations"].Allowed)

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			limit, page, opt, getTotal := withPagination(cmd.Flags())
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).ListDefaultAssociations(ctx, &ttnpb.ListApplicationPackageDefaultAssociationRequest{
				Ids:       appID,
				Limit:     limit,
				Page:      page,
				FieldMask: ttnpb.FieldMask(paths...),
			}, opt)
			if err != nil {
				return err
			}
			getTotal()

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackageDefaultAssociationSetCommand = &cobra.Command{
		Use:     "set [application-id] [f-port]",
		Aliases: []string{"update"},
		Short:   "Set the properties of an application package default association",
		RunE: func(cmd *cobra.Command, args []string) error {
			association := &ttnpb.ApplicationPackageDefaultAssociation{}
			paths, err := association.SetFromFlags(cmd.Flags(), "")
			if err != nil {
				return err
			}
			assocID, err := getApplicationPackageDefaultAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}
			association.Ids = assocID

			reader, err := getDataReader("data", cmd.Flags())
			if err != nil {
				if _, needsData := packageNeedsData[association.PackageName]; needsData {
					logger.WithError(err).Warn("Package data not available")
				}
			} else {
				var st pbtypes.Struct
				err := jsonpb.TTN().NewDecoder(reader).Decode(&st)
				if err != nil {
					return err
				}

				association.Data = &st
				paths = append(paths, "data")
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			res, err := ttnpb.NewApplicationPackageRegistryClient(as).SetDefaultAssociation(ctx, &ttnpb.SetApplicationPackageDefaultAssociationRequest{
				Default:   association,
				FieldMask: ttnpb.FieldMask(paths...),
			})
			if err != nil {
				return err
			}

			return io.Write(os.Stdout, config.OutputFormat, res)
		},
	}
	applicationsPackageDefaultAssociationDeleteCommand = &cobra.Command{
		Use:     "delete [application-id] [f-port]",
		Aliases: []string{"del", "remove", "rm"},
		Short:   "Delete an application package default association",
		RunE: func(cmd *cobra.Command, args []string) error {
			assocID, err := getApplicationPackageDefaultAssociationID(cmd.Flags(), args)
			if err != nil {
				return err
			}

			as, err := api.Dial(ctx, config.ApplicationServerGRPCAddress)
			if err != nil {
				return err
			}
			_, err = ttnpb.NewApplicationPackageRegistryClient(as).DeleteDefaultAssociation(ctx, assocID)
			if err != nil {
				return err
			}

			return nil
		},
	}
)

func init() {
	ttnpb.AddSelectFlagsForApplicationPackageAssociation(selectApplicationPackageAssociationsFlags, "", false)
	ttnpb.AddSelectFlagsForApplicationPackageDefaultAssociation(selectApplicationPackageDefaultAssociationsFlags, "", false)
	applicationsPackagesCommand.AddCommand(applicationsPackagesListCommand)
	applicationsPackageAssociationGetCommand.Flags().AddFlagSet(applicationPackageAssociationIDFlags())
	applicationsPackageAssociationGetCommand.Flags().AddFlagSet(selectApplicationPackageAssociationsFlags)
	applicationsPackageAssociationGetCommand.Flags().AddFlagSet(selectAllApplicationPackageAssociationsFlags)
	applicationsPackagesCommand.AddCommand(applicationsPackagesAssociationsCommand)
	applicationsPackagesAssociationsCommand.AddCommand(applicationsPackageAssociationGetCommand)
	applicationsPackageAssociationsListCommand.Flags().AddFlagSet(endDeviceIDFlags())
	applicationsPackageAssociationsListCommand.Flags().AddFlagSet(selectApplicationPackageAssociationsFlags)
	applicationsPackageAssociationsListCommand.Flags().AddFlagSet(selectAllApplicationPackageAssociationsFlags)
	applicationsPackageAssociationsListCommand.Flags().AddFlagSet(paginationFlags())
	applicationsPackagesAssociationsCommand.AddCommand(applicationsPackageAssociationsListCommand)
	ttnpb.AddSetFlagsForApplicationPackageAssociation(applicationsPackageAssociationSetCommand.Flags(), "", false)
	flagsplugin.AddAlias(applicationsPackageAssociationSetCommand.Flags(), "ids.end-device-ids.application-ids.application-id", "application-id", flagsplugin.WithHidden(false))
	flagsplugin.AddAlias(applicationsPackageAssociationSetCommand.Flags(), "ids.end-device-ids.device-id", "device-id", flagsplugin.WithHidden(false))
	flagsplugin.AddAlias(applicationsPackageAssociationSetCommand.Flags(), "ids.f-port", "f-port", flagsplugin.WithHidden(false))
	applicationsPackageAssociationSetCommand.Flags().AddFlagSet(dataFlags("data", "package data"))
	applicationsPackagesAssociationsCommand.AddCommand(applicationsPackageAssociationSetCommand)
	applicationsPackageAssociationDeleteCommand.Flags().AddFlagSet(applicationPackageAssociationIDFlags())
	applicationsPackagesAssociationsCommand.AddCommand(applicationsPackageAssociationDeleteCommand)
	applicationsPackagesCommand.AddCommand(applicationsPackagesDefaultAssociationsCommand)
	applicationsPackagesDefaultAssociationsCommand.AddCommand(applicationsPackageDefaultAssociationGetCommand)
	applicationsPackageDefaultAssociationsListCommand.Flags().AddFlagSet(applicationIDFlags())
	applicationsPackageDefaultAssociationsListCommand.Flags().AddFlagSet(selectApplicationPackageDefaultAssociationsFlags)
	applicationsPackageDefaultAssociationsListCommand.Flags().AddFlagSet(selectAllApplicationPackageDefaultAssociationsFlags)
	applicationsPackageDefaultAssociationsListCommand.Flags().AddFlagSet(paginationFlags())
	applicationsPackagesDefaultAssociationsCommand.AddCommand(applicationsPackageDefaultAssociationsListCommand)
	ttnpb.AddSetFlagsForApplicationPackageDefaultAssociation(applicationsPackageDefaultAssociationSetCommand.Flags(), "", false)
	flagsplugin.AddAlias(applicationsPackageDefaultAssociationSetCommand.Flags(), "ids.application-ids.application-id", "application-id", flagsplugin.WithHidden(false))
	flagsplugin.AddAlias(applicationsPackageDefaultAssociationSetCommand.Flags(), "ids.f-port", "f-port", flagsplugin.WithHidden(false))
	applicationsPackageDefaultAssociationSetCommand.Flags().AddFlagSet(dataFlags("data", "package data"))
	applicationsPackagesDefaultAssociationsCommand.AddCommand(applicationsPackageDefaultAssociationSetCommand)
	applicationsPackageDefaultAssociationDeleteCommand.Flags().AddFlagSet(applicationPackageDefaultAssociationIDFlags())
	applicationsPackagesDefaultAssociationsCommand.AddCommand(applicationsPackageDefaultAssociationDeleteCommand)
	applicationsCommand.AddCommand(applicationsPackagesCommand)
}
