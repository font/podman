package images

import (
	"github.com/containers/podman/v3/cmd/podman/common"
	"github.com/containers/podman/v3/cmd/podman/registry"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/spf13/cobra"
)

var (
	sigstoreSignDescription = "Sign image and create entry in sigstore."
	sigstoreSignCommand     = &cobra.Command{
		Use:               "sign [options] IMAGE [IMAGE...]",
		Short:             "Sign an image using Sigstore",
		Long:              sigstoreSignDescription,
		RunE:              sigstoreSign,
		Args:              cobra.MinimumNArgs(1),
		ValidArgsFunction: common.AutocompleteImages,
		Example:           `podman image sigstore sign imageID`,
	}
)

var (
	sigstoreSignOptions entities.SigstoreSignOptions
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Mode:    []entities.EngineMode{entities.ABIMode},
		Command: sigstoreSignCommand,
		Parent:  sigstoreCmd,
	})
}

func sigstoreSign(cmd *cobra.Command, args []string) error {
	_, err := registry.ImageEngine().SigstoreSign(registry.Context(), args, sigstoreSignOptions)
	return err
}
