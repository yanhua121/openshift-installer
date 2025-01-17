package bootkube

import (
	"context"
	"os"
	"path/filepath"

	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/templates/content"
)

const (
	machineConfigServerCAConfigMapFileName = "machine-config-server-ca-configmap.yaml.template"
)

var _ asset.WritableAsset = (*MachineConfigServerCAConfigMap)(nil)

// MachineConfigServerCAConfigMap is the constant to represent contents of machine-config-server-ca-configmap.yaml.template file.
type MachineConfigServerCAConfigMap struct {
	FileList []*asset.File
}

// Dependencies returns all of the dependencies directly needed by the asset.
func (t *MachineConfigServerCAConfigMap) Dependencies() []asset.Asset {
	return []asset.Asset{}
}

// Name returns the human-friendly name of the asset.
func (t *MachineConfigServerCAConfigMap) Name() string {
	return "MachineConfigServerCAConfigMap"
}

// Generate generates the actual files by this asset.
func (t *MachineConfigServerCAConfigMap) Generate(_ context.Context, parents asset.Parents) error {
	fileName := machineConfigServerCAConfigMapFileName
	data, err := content.GetBootkubeTemplate(fileName)
	if err != nil {
		return err
	}
	t.FileList = []*asset.File{
		{
			Filename: filepath.Join(content.TemplateDir, fileName),
			Data:     data,
		},
	}
	return nil
}

// Files returns the files generated by the asset.
func (t *MachineConfigServerCAConfigMap) Files() []*asset.File {
	return t.FileList
}

// Load returns the asset from disk.
func (t *MachineConfigServerCAConfigMap) Load(f asset.FileFetcher) (bool, error) {
	file, err := f.FetchByName(filepath.Join(content.TemplateDir, machineConfigServerCAConfigMapFileName))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	t.FileList = []*asset.File{file}
	return true, nil
}
