package mysql

import (
	"github.com/spf13/cobra"
	"github.com/wal-g/tracelog"
	"github.com/wal-g/wal-g/internal"
	"github.com/wal-g/wal-g/internal/databases/mysql"
	"github.com/wal-g/wal-g/utility"
)

const (
	backupListShortDescription = "Prints available backups"
	PrettyFlag                 = "pretty"
	JSONFlag                   = "json"
	DetailFlag                 = "detail"
)

var (
	// backupListCmd represents the backupList command
	backupListCmd = &cobra.Command{
		Use:   "backup-list",
		Short: backupListShortDescription,
		Args:  cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			folder, err := internal.ConfigureFolder()
			tracelog.ErrorLogger.FatalOnError(err)
			if detail {
				mysql.HandleDetailedBackupList(folder.GetSubFolder(utility.BaseBackupPath), pretty, json)
			} else {
				internal.HandleDefaultBackupList(folder.GetSubFolder(utility.BaseBackupPath), pretty, json)
			}
		},
	}
	json   = false
	pretty = false
	detail = false
)

func init() {
	cmd.AddCommand(backupListCmd)

	// TODO: Merge similar backup-list functionality
	// to avoid code duplication in command handlers
	backupListCmd.Flags().BoolVar(&pretty, PrettyFlag, false, "Prints more readable output")
	backupListCmd.Flags().BoolVar(&json, JSONFlag, false, "Prints output in json format")
	backupListCmd.Flags().BoolVar(&detail, DetailFlag, false, "Prints extra backup details")
}
