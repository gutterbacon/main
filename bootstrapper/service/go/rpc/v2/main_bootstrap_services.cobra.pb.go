// Code generated by protoc-gen-cobra. DO NOT EDIT.

package v2

import (
	client "github.com/getcouragenow/protoc-gen-cobra/client"
	flag "github.com/getcouragenow/protoc-gen-cobra/flag"
	iocodec "github.com/getcouragenow/protoc-gen-cobra/iocodec"
	cobra "github.com/spf13/cobra"
	grpc "google.golang.org/grpc"
	proto "google.golang.org/protobuf/proto"
	io "io"
)

func BSServiceClientCommand(options ...client.Option) *cobra.Command {
	cfg := client.NewConfig(options...)
	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("BSService"),
		Short:  "BSService service client",
		Long:   "",
		Hidden: false,
	}
	cfg.BindFlags(cmd.PersistentFlags())
	cmd.AddCommand(
		_BSServiceNewBootstrapCommand(cfg),
		_BSServiceGetBootstrapCommand(cfg),
		_BSServiceListBootstrapCommand(cfg),
		_BSServiceExecuteBootstrapCommand(cfg),
		_BSServiceDeleteBootstrapCommand(cfg),
	)
	return cmd
}

func _BSServiceNewBootstrapCommand(cfg *client.Config) *cobra.Command {
	req := &NewBSRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("NewBootstrap"),
		Short:  "NewBootstrap RPC client",
		Long:   "hide",
		Hidden: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService", "NewBootstrap"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewBSServiceClient(cc)
				v := &NewBSRequest{}

				stm, err := cli.NewBootstrap(cmd.Context())
				if err != nil {
					return err
				}
				for {
					if err := in(v); err != nil {
						if err == io.EOF {
							_ = stm.CloseSend()
							break
						}
						return err
					}
					proto.Merge(v, req)
					if err = stm.Send(v); err != nil {
						return err
					}
				}

				res, err := stm.CloseAndRecv()
				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.FileExtension, cfg.FlagNamer("FileExtension"), "", "either file_path or file content (in bytes)")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.BsRequest, cfg.FlagNamer("BsRequest"), "")

	return cmd
}

func _BSServiceGetBootstrapCommand(cfg *client.Config) *cobra.Command {
	req := &GetBSRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("GetBootstrap"),
		Short:  "GetBootstrap RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService", "GetBootstrap"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewBSServiceClient(cc)
				v := &GetBSRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.GetBootstrap(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.FileId, cfg.FlagNamer("FileId"), "", "")

	return cmd
}

func _BSServiceListBootstrapCommand(cfg *client.Config) *cobra.Command {
	req := &ListBSRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("ListBootstrap"),
		Short:  "ListBootstrap RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService", "ListBootstrap"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewBSServiceClient(cc)
				v := &ListBSRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ListBootstrap(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().Int64Var(&req.PerPageEntries, cfg.FlagNamer("PerPageEntries"), 0, "")
	cmd.PersistentFlags().StringVar(&req.OrderBy, cfg.FlagNamer("OrderBy"), "", "")
	cmd.PersistentFlags().StringVar(&req.CurrentPageId, cfg.FlagNamer("CurrentPageId"), "", "")
	cmd.PersistentFlags().BoolVar(&req.IsDescending, cfg.FlagNamer("IsDescending"), false, "")
	flag.BytesBase64Var(cmd.PersistentFlags(), &req.Filters, cfg.FlagNamer("Filters"), "")

	return cmd
}

func _BSServiceExecuteBootstrapCommand(cfg *client.Config) *cobra.Command {
	req := &GetBSRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("ExecuteBootstrap"),
		Short:  "ExecuteBootstrap RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService", "ExecuteBootstrap"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewBSServiceClient(cc)
				v := &GetBSRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.ExecuteBootstrap(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.FileId, cfg.FlagNamer("FileId"), "", "")

	return cmd
}

func _BSServiceDeleteBootstrapCommand(cfg *client.Config) *cobra.Command {
	req := &GetBSRequest{}

	cmd := &cobra.Command{
		Use:    cfg.CommandNamer("DeleteBootstrap"),
		Short:  "DeleteBootstrap RPC client",
		Long:   "",
		Hidden: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			if cfg.UseEnvVars {
				if err := flag.SetFlagsFromEnv(cmd.Parent().PersistentFlags(), true, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService"); err != nil {
					return err
				}
				if err := flag.SetFlagsFromEnv(cmd.PersistentFlags(), false, cfg.EnvVarNamer, cfg.EnvVarPrefix, "BSService", "DeleteBootstrap"); err != nil {
					return err
				}
			}
			return client.RoundTrip(cmd.Context(), cfg, func(cc grpc.ClientConnInterface, in iocodec.Decoder, out iocodec.Encoder) error {
				cli := NewBSServiceClient(cc)
				v := &GetBSRequest{}

				if err := in(v); err != nil {
					return err
				}
				proto.Merge(v, req)

				res, err := cli.DeleteBootstrap(cmd.Context(), v)

				if err != nil {
					return err
				}

				return out(res)

			})
		},
	}

	cmd.PersistentFlags().StringVar(&req.FileId, cfg.FlagNamer("FileId"), "", "")

	return cmd
}
