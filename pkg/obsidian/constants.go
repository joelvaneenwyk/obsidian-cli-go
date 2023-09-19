package obsidian

const (
	ExecuteUriError            = "failed to execute Obsidian URI"
	NoteDoesNotExistError      = "cannot find note in vault"
	VaultAccessError           = "failed to access vault directory"
	VaultReadError             = "failed to read notes in vault"
	VaultWriteError            = "failed to write to update notes in vault"
	CliConfigReadError         = "cannot find vault config, please use set-default command to set default vault or use --vault flag"
	CliConfigParseError        = "could not parse vault config file, please use set-default command to set default vault or use --vault flag"
	CliConfigDirWriteError     = "failed to create vault config directory. Please ensure you have the correct permissions."
	CliConfigGenerateJsonError = "failed to generate vault config file. Please ensure vault name does not contain any special characters."
	CliConfigWriteError        = "failed to write vault config file. Please ensure you have correct permissions."
	ConfigReadError            = "failed to read Obsidian config file. Please ensure vault has been set up in Obsidian."
	ConfigParseError           = "failed to parse Obsidian config file. Please ensure vault has been set up in Obsidian."
	ConfigVaultNotFoundError   = "vault not found in Obsidian config file. Please ensure vault has been set up in Obsidian."
)
