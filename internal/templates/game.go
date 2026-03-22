package templates

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type GameTemplate struct {
	Name        string             `json:"name"`
	DisplayName string             `json:"display_name"`
	Description string             `json:"description"`
	Resources   ResourceConfig     `json:"resources"`
	Network     NetworkConfig      `json:"network"`
	Settings    map[string]Setting `json:"settings"`
	HelmChart   HelmConfig         `json:"helm_chart"`
	Volumes     []VolumeConfig     `json:"volumes"`
}

type ResourceConfig struct {
	CPU       string `json:"cpu"`
	Memory    string `json:"memory"`
	Storage   string `json:"storage"`
	MinCPU    string `json:"min_cpu,omitempty"`
	MinMemory string `json:"min_memory,omitempty"`
}

type NetworkConfig struct {
	Ports []PortConfig `json:"ports"`
}

type PortConfig struct {
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Name     string `json:"name,omitempty"`
}

type Setting struct {
	Type        string      `json:"type"` // string, int, bool, select
	Default     interface{} `json:"default"`
	Description string      `json:"description"`
	Required    bool        `json:"required"`
	Options     []string    `json:"options,omitempty"` // For select type
	Min         *int        `json:"min,omitempty"`     // For int type
	Max         *int        `json:"max,omitempty"`     // For int type
}

type HelmConfig struct {
	Chart      string                 `json:"chart"`
	Repository string                 `json:"repository,omitempty"`
	Version    string                 `json:"version,omitempty"`
	Values     map[string]interface{} `json:"values,omitempty"`
}

type VolumeConfig struct {
	Name      string `json:"name"`
	MountPath string `json:"mount_path"`
	Size      string `json:"size"`
	Type      string `json:"type"` // pvc, hostPath, etc.
}

type TemplateManager struct {
	templates map[string]*GameTemplate
	configDir string
}

func NewTemplateManager(configDir string) *TemplateManager {
	return &TemplateManager{
		templates: make(map[string]*GameTemplate),
		configDir: configDir,
	}
}

func (tm *TemplateManager) LoadTemplates() error {
	entries, err := os.ReadDir(tm.configDir)
	if err != nil {
		return fmt.Errorf("failed to read template directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ".json" {
			if err := tm.loadTemplate(filepath.Join(tm.configDir, entry.Name())); err != nil {
				return fmt.Errorf("failed to load template %s: %w", entry.Name(), err)
			}
		}
	}

	return nil
}

func (tm *TemplateManager) loadTemplate(filePath string) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	var template GameTemplate
	if err := json.Unmarshal(data, &template); err != nil {
		return err
	}

	tm.templates[template.Name] = &template
	return nil
}

func (tm *TemplateManager) GetTemplate(name string) (*GameTemplate, error) {
	template, exists := tm.templates[name]
	if !exists {
		return nil, fmt.Errorf("template %s not found", name)
	}

	// Return a copy to avoid modifications
	templateCopy := *template
	return &templateCopy, nil
}

func (tm *TemplateManager) ListTemplates() []string {
	var names []string
	for name := range tm.templates {
		names = append(names, name)
	}
	return names
}

func (tm *TemplateManager) ValidateTemplate(template *GameTemplate) error {
	if template.Name == "" {
		return fmt.Errorf("template name is required")
	}
	if template.Resources.CPU == "" || template.Resources.Memory == "" {
		return fmt.Errorf("CPU and Memory resources are required")
	}
	if len(template.Network.Ports) == 0 {
		return fmt.Errorf("at least one port is required")
	}
	return nil
}
