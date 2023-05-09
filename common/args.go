package common

import "github.com/urfave/cli"

var (
	SrcRepo = cli.StringFlag{
		Name:     "src-repo",
		Value:    "admin@127.0.0.1:datastore",
		Usage:    "Source repo url",
		Required: true,
	}
	SrcNamespace = cli.StringFlag{
		Name:  "src-namespace",
		Usage: "Source datastore namespace",
	}
	SrcPassword = cli.StringFlag{
		Name:     "src-password",
		Usage:    "Source repo password",
		Required: true,
	}
	SrcFingerprint = cli.StringFlag{
		Name:     "src-fingerprint",
		Usage:    "Source repo fingerprint",
		Required: true,
	}
	SrcKeyFile = cli.StringFlag{
		Name:  "src-keyfile",
		Usage: "Source key file",
	}
	SrcKeyPassword = cli.StringFlag{
		Name:  "src-keypassword",
		Usage: "Source key password",
	}

	DstRepo = cli.StringFlag{
		Name:     "dst-repo",
		Value:    "admin@127.0.0.1:datastore",
		Usage:    "Destination repo url",
		Required: true,
	}
	DstNamespace = cli.StringFlag{
		Name:  "dst-namespace",
		Usage: "Destination datastore namespace",
	}
	DstPassword = cli.StringFlag{
		Name:     "dst-password",
		Usage:    "Destination repo password",
		Required: true,
	}
	DstFingerprint = cli.StringFlag{
		Name:     "dst-fingerprint",
		Usage:    "Destination repo fingerprint",
		Required: true,
	}
	DstKeyFile = cli.StringFlag{
		Name:  "dst-keyfile",
		Usage: "Destination key file",
	}
	DstKeyPassword = cli.StringFlag{
		Name:  "dst-keypassword",
		Usage: "Destination key password",
	}

	BackupId = cli.StringFlag{
		Name:     "backup-id",
		Usage:    "backup id",
		Required: true,
	}
	BackupTime = cli.Uint64Flag{
		Name:     "backup-time",
		Value:    0,
		Usage:    "backup time",
		Required: true,
	}
)
