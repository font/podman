package images

import (
	"github.com/containers/podman/v3/cmd/podman/registry"
	"github.com/containers/podman/v3/cmd/podman/validate"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/spf13/cobra"
)

var (
	// Command: podman image _sigstore_
	sigstoreCmd = &cobra.Command{
		Use:   "sigstore",
		Short: "Sign/Verify images using sigstore",
		Long:  "Sign/Verify images using sigstore",
		RunE:  validate.SubCommandExists,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Mode:    []entities.EngineMode{entities.ABIMode, entities.TunnelMode},
		Command: sigstoreCmd,
		Parent:  imageCmd,
	})
}
