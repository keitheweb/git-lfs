package commands

import (
	"github.com/github/git-lfs/lfs"
	"github.com/spf13/cobra"
)

// uninstallCmd removes any configuration and hooks set by Git LFS.
func uninstallCommand(cmd *cobra.Command, args []string) {
	if err := lfs.UninstallFilters(); err != nil {
		Error(err.Error())
	}

	Print("Global Git LFS configuration has been removed.")

	if lfs.InRepo() {
		uninstallHooksCommand(cmd, args)
	}
}

// uninstallHooksCmd removes any hooks created by Git LFS.
func uninstallHooksCommand(cmd *cobra.Command, args []string) {
	if err := lfs.UninstallHooks(); err != nil {
		Error(err.Error())
	}

	Print("Hooks for this repository have been removed.")
}

func init() {
	RegisterSubcommand(func() *cobra.Command {
		cmd := &cobra.Command{
			Use:    "uninstall",
			PreRun: resolveLocalStorage,
			Run:    uninstallCommand,
		}

		cmd.AddCommand(&cobra.Command{
			Use:    "hooks",
			PreRun: resolveLocalStorage,
			Run:    uninstallHooksCommand,
		})
		return cmd
	})
}
