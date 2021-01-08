package keypairs

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/ci/gcloud"
	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
)

// Load will attempt to load a Keypair using private and public-key files from
// the user's file system; specifically from the config dir. It is assumed that
// this keypair file has no passphrase, even if it is encrypted.
func Load(keyName string) (Keypair, error) {
	keyFilename := LocalKeyFilename(keyName)
	if err := validateKeyFile(keyFilename); err != nil {
		return nil, err
	}
	return loadAndParseKeypair(keyFilename)
}

// Save will save the unencrypted and encoded private key to a local config
// file. The filename will be the value of `keyName` and suffixed with `.key`.
func Save(kp Keypair, keyName string) error {
	err := ioutil.WriteFile(LocalKeyFilename(keyName), []byte(kp.EncodePrivateKey()), 0600)
	if err != nil {
		return errs.Wrap(err, "WriteFile failed")
	}
	return nil
}

// Delete will delete an unencrypted and encoded private key from the local
// config directory. The base filename (sans suffix) must be provided.
func Delete(keyName string) error {
	filename := LocalKeyFilename(keyName)
	if fileutils.FileExists(filename) {
		if err := os.Remove(filename); err != nil {
			return errs.Wrap(err, "os.Remove %s failed", filename)
		}
	}
	return nil
}

// LoadWithDefaults will call Load with the default key name (i.e.
// constants.KeypairLocalFileName). If the key override is set
// (constants.PrivateKeyEnvVarName), that value will be parsed directly.
func LoadWithDefaults() (Keypair, error) {
	key, err := gcloud.GetSecret(constants.PrivateKeyEnvVarName)
	if err != nil && !errors.Is(err, gcloud.ErrNotAvailable{}) {
		return nil, errs.Wrap(err, "gcloud.GetSecret failed")
	}
	if err == nil && key != "" {
		logging.Debug("Using private key sourced from gcloud")
		return ParseRSA(key)
	}

	if key := os.Getenv(constants.PrivateKeyEnvVarName); key != "" {
		logging.Debug("Using private key sourced from environment")
		return ParseRSA(key)
	}

	return Load(constants.KeypairLocalFileName)
}

// SaveWithDefaults will call Save with the provided keypair and the default
// key name (i.e. constants.KeypairLocalFileName). The operation will fail when
// the key override is set (constants.PrivateKeyEnvVarName).
func SaveWithDefaults(kp Keypair) error {
	if hasKeyOverride() {
		return locale.NewInputError("keypairs_err_override_with_save")
	}

	return Save(kp, constants.KeypairLocalFileName)
}

// DeleteWithDefaults will call Delete with the default key name (i.e.
// constants.KeypairLocalFileName). The operation will fail when the key
// override is set (constants.PrivateKeyEnvVarName).
func DeleteWithDefaults() error {
	if hasKeyOverride() {
		return locale.NewInputError("keypairs_err_override_with_delete")
	}

	return Delete(constants.KeypairLocalFileName)
}

// LocalKeyFilename returns the full filepath for the given key name
func LocalKeyFilename(keyName string) string {
	return filepath.Join(config.Get().ConfigPath(), keyName+".key")
}

func loadAndParseKeypair(keyFilename string) (Keypair, error) {
	keyFileBytes, err := ioutil.ReadFile(keyFilename)
	if err != nil {
		return nil, errs.Wrap(err, "ReadFile %s failed", keyFilename)
	}
	return ParseRSA(string(keyFileBytes))
}

func hasKeyOverride() bool {
	if os.Getenv(constants.PrivateKeyEnvVarName) != "" {
		return true
	}

	tkn, err := gcloud.GetSecret(constants.PrivateKeyEnvVarName)
	if err != nil && !errors.Is(err, gcloud.ErrNotAvailable{}) {
		logging.Error("Could not retrieve gcloud secret: %v", err)
	}
	if err == nil && tkn != "" {
		return true
	}

	return false
}
