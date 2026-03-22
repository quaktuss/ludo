package bot

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"ludo-bot/internal/templates"
)

func (b *Bot) routeButtonInteraction(s *discordgo.Session, i *discordgo.InteractionCreate, customID string) {
	switch {
	case strings.HasPrefix(customID, "modify_resources_"):
		b.handleModifyResources(s, i)
	case strings.HasPrefix(customID, "modify_settings1_"):
		b.handleModifySettings1(s, i)
	case strings.HasPrefix(customID, "modify_settings2_"):
		b.handleModifySettings2(s, i)
	case strings.HasPrefix(customID, "modify_all_settings_"):
		b.handleModifyAllSettings(s, i)
	case strings.HasPrefix(customID, "create_confirm_"):
		b.handleCreateConfirm(s, i)
	case customID == "create_cancel":
		b.handleCreateCancel(s, i)
	case strings.HasPrefix(customID, "view_yaml_"):
		b.handleViewYAML(s, i)
	}
}

func (b *Bot) handleModifyResources(s *discordgo.Session, i *discordgo.InteractionCreate) {
	parts := strings.Split(i.MessageComponentData().CustomID, "_")
	if len(parts) < 4 {
		b.respondError(s, i, "Invalid button data")
		return
	}

	gameType := parts[2]
	serverName := parts[3]

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: fmt.Sprintf("resources_modal_%s_%s", gameType, serverName),
			Title:    "Modify Resources",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					discordgo.TextInput{CustomID: "cpu", Label: "CPU (cores)", Style: discordgo.TextInputShort, Placeholder: "1", Required: true, MaxLength: 10},
				}},
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					discordgo.TextInput{CustomID: "memory", Label: "Memory (e.g., 2Gi)", Style: discordgo.TextInputShort, Placeholder: "2Gi", Required: true, MaxLength: 10},
				}},
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					discordgo.TextInput{CustomID: "storage", Label: "Storage (e.g., 5Gi)", Style: discordgo.TextInputShort, Placeholder: "5Gi", Required: true, MaxLength: 10},
				}},
			},
		},
	})
}

func (b *Bot) handleModifySettings1(s *discordgo.Session, i *discordgo.InteractionCreate) {
	parts := strings.Split(i.MessageComponentData().CustomID, "_")
	if len(parts) < 4 {
		b.respondError(s, i, "Invalid button data")
		return
	}

	gameType := parts[2]
	serverName := parts[3]
	template, err := b.templateManager.GetTemplate(gameType)
	if err != nil {
		b.respondError(s, i, "Template not found")
		return
	}

	mainSettings := []string{"max_players", "server_password", "server_name", "pvp", "admin_password"}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID:   fmt.Sprintf("settings1_modal_%s_%s", gameType, serverName),
			Title:      "Main Game Settings (1/2)",
			Components: b.createSettingsComponents(template, mainSettings),
		},
	})
}

func (b *Bot) handleModifySettings2(s *discordgo.Session, i *discordgo.InteractionCreate) {
	parts := strings.Split(i.MessageComponentData().CustomID, "_")
	if len(parts) < 4 {
		b.respondError(s, i, "Invalid button data")
		return
	}

	gameType := parts[2]
	serverName := parts[3]
	template, err := b.templateManager.GetTemplate(gameType)
	if err != nil {
		b.respondError(s, i, "Template not found")
		return
	}

	advancedSettings := []string{"pause_empty", "save_world_every_minutes", "public_server", "autosave_interval"}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID:   fmt.Sprintf("settings2_modal_%s_%s", gameType, serverName),
			Title:      "Advanced Game Settings (2/2)",
			Components: b.createSettingsComponents(template, advancedSettings),
		},
	})
}

func (b *Bot) handleSelectMenu(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	if len(data.Values) == 0 {
		return
	}

	selectedSetting := data.Values[0]
	parts := strings.Split(data.CustomID, "_")
	if len(parts) < 4 {
		b.respondError(s, i, "Invalid select data")
		return
	}

	gameType := parts[2]
	serverName := parts[3]
	template, err := b.templateManager.GetTemplate(gameType)
	if err != nil {
		b.respondError(s, i, "Template not found")
		return
	}

	setting, exists := template.Settings[selectedSetting]
	if !exists {
		b.respondError(s, i, "Setting not found")
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: fmt.Sprintf("single_setting_modal_%s_%s_%s", selectedSetting, gameType, serverName),
			Title:    fmt.Sprintf("Modify: %s", selectedSetting),
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    selectedSetting,
						Label:       setting.Description,
						Style:       discordgo.TextInputShort,
						Placeholder: fmt.Sprintf("%v", setting.Default),
						Required:    setting.Required,
						MaxLength:   200,
					},
				}},
			},
		},
	})
}

func (b *Bot) handleModifyAllSettings(s *discordgo.Session, i *discordgo.InteractionCreate) {
	parts := strings.Split(i.MessageComponentData().CustomID, "_")
	if len(parts) < 5 {
		b.respondError(s, i, "Invalid button data")
		return
	}

	gameType := parts[3]
	serverName := parts[4]
	template, err := b.templateManager.GetTemplate(gameType)
	if err != nil {
		b.respondError(s, i, "Template not found")
		return
	}

	var settingsText strings.Builder
	for name, setting := range template.Settings {
		settingsText.WriteString(fmt.Sprintf("%s=%v\n", name, setting.Default))
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseModal,
		Data: &discordgo.InteractionResponseData{
			CustomID: fmt.Sprintf("all_settings_modal_%s_%s", gameType, serverName),
			Title:    "All Game Settings (Advanced)",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:  "all_settings",
						Label:     "Settings (key=value format, one per line)",
						Style:     discordgo.TextInputParagraph,
						Value:     settingsText.String(),
						Required:  false,
						MaxLength: 4000,
					},
				}},
			},
		},
	})
}

func (b *Bot) createSettingsComponents(template *templates.GameTemplate, settingNames []string) []discordgo.MessageComponent {
	var components []discordgo.MessageComponent
	count := 0

	for _, name := range settingNames {
		if setting, exists := template.Settings[name]; exists && count < 5 {
			components = append(components, discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.TextInput{
						CustomID:    name,
						Label:       setting.Description,
						Style:       discordgo.TextInputShort,
						Placeholder: fmt.Sprintf("%v", setting.Default),
						Required:    setting.Required,
						MaxLength:   100,
					},
				},
			})
			count++
		}
	}

	return components
}
