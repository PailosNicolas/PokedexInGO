package commands

import (
	"os"

	"github.com/PailosNicolas/PokedexInGO/structs"
)

func CommandExit(cfg *structs.Config, args ...string) error {
	os.Exit(0)
	return nil
}
