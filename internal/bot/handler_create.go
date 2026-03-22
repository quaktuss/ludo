package bot

import (
	"fmt"
	"log"
	"strings"

	"github.com/bwmarrin/discordgo"
	"ludo-bot/internal/templates"
)

func (b *Bot) handleCreateServer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	gameType := data.Options[0].StringValue()
	serverName := data.Options[1].StringValue()

	template, err := b.templateManager.GetTemplate(gameType)
	if err != nil {
		b.respondError(s, i, fmt.Sprintf("Game template not found: %v", err))
		return
	}

	embed := b.createTemplateEmbed(template, serverName)
	components := b.createTemplateComponents(gameType, serverName)

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds:     []*discordgo.MessageEmbed{embed},
			Components: components,
		},
	})
	if err != nil {
		log.Printf("Error sending template configuration: %v", err)
	}
}

func (b *Bot) createTemplateEmbed(template *templates.GameTemplate, serverName string) *discordgo.MessageEmbed {
	var imageURL string
	switch template.Name {
	case "project-zomboid":
		imageURL = "https://cdn.cloudflare.steamstatic.com/steam/apps/108600/header.jpg"
	default:
		imageURL = "https://via.placeholder.com/400x200/2f3136/ffffff?text=Game+Server"
	}

	return &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("🎮 Configure %s Server", template.DisplayName),
		Description: template.Description,
		Color:       0x00ff00,
		Image:       &discordgo.MessageEmbedImage{URL: imageURL},
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "📊 Resources",
				Value:  fmt.Sprintf("• CPU: %s\n• Memory: %s\n• Storage: %s", template.Resources.CPU, template.Resources.Memory, template.Resources.Storage),
				Inline: true,
			},
			{
				Name:   "🌐 Network",
				Value:  b.formatPorts(template.Network.Ports),
				Inline: true,
			},
			{
				Name:   "⚙️ Game Settings",
				Value:  b.formatSettings(template.Settings),
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: fmt.Sprintf("Server: %s", serverName),
		},
	}
}

func (b *Bot) formatPorts(ports []templates.PortConfig) string {
	var parts []string
	for _, port := range ports {
		parts = append(parts, fmt.Sprintf("• %d/%s", port.Port, port.Protocol))
	}
	return strings.Join(parts, "\n")
}

func (b *Bot) formatSettings(settings map[string]templates.Setting) string {
	var parts []string
	for name, setting := range settings {
		parts = append(parts, fmt.Sprintf("• %s: %v", name, setting.Default))
	}
	return strings.Join(parts, "\n")
}

func (b *Bot) createTemplateComponents(gameType, serverName string) []discordgo.MessageComponent {
	return []discordgo.MessageComponent{
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "📊 Resources",
					Style:    discordgo.SecondaryButton,
					CustomID: fmt.Sprintf("modify_resources_%s_%s", gameType, serverName),
				},
				discordgo.Button{
					Label:    "🌐 Network",
					Style:    discordgo.SecondaryButton,
					CustomID: fmt.Sprintf("modify_network_%s_%s", gameType, serverName),
				},
				discordgo.Button{
					Label:    "⚙️ Settings 1/2",
					Style:    discordgo.SecondaryButton,
					CustomID: fmt.Sprintf("modify_settings1_%s_%s", gameType, serverName),
				},
				discordgo.Button{
					Label:    "⚙️ Settings 2/2",
					Style:    discordgo.SecondaryButton,
					CustomID: fmt.Sprintf("modify_settings2_%s_%s", gameType, serverName),
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.SelectMenu{
					CustomID:    fmt.Sprintf("choose_setting_%s_%s", gameType, serverName),
					Placeholder: "🔧 Choose specific setting to modify",
					Options: []discordgo.SelectMenuOption{
						{Label: "Max Players", Value: "max_players", Description: "Number of players allowed"},
						{Label: "Server Password", Value: "server_password", Description: "Server access password"},
						{Label: "Server Name", Value: "server_name", Description: "Display name of server"},
						{Label: "PvP Mode", Value: "pvp", Description: "Enable/disable PvP"},
						{Label: "Admin Password", Value: "admin_password", Description: "Admin access password"},
						{Label: "Pause When Empty", Value: "pause_empty", Description: "Pause when no players"},
						{Label: "Save Interval", Value: "save_world_every_minutes", Description: "Auto-save frequency"},
						{Label: "Public Server", Value: "public_server", Description: "Visible in public list"},
					},
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "📝 All Settings (Advanced)",
					Style:    discordgo.SecondaryButton,
					CustomID: fmt.Sprintf("modify_all_settings_%s_%s", gameType, serverName),
				},
			},
		},
		discordgo.ActionsRow{
			Components: []discordgo.MessageComponent{
				discordgo.Button{
					Label:    "✅ Create Server",
					Style:    discordgo.SuccessButton,
					CustomID: fmt.Sprintf("create_confirm_%s_%s", gameType, serverName),
				},
				discordgo.Button{
					Label:    "❌ Cancel",
					Style:    discordgo.DangerButton,
					CustomID: "create_cancel",
				},
				discordgo.Button{
					Label:    "📋 View YAML",
					Style:    discordgo.PrimaryButton,
					CustomID: fmt.Sprintf("view_yaml_%s_%s", gameType, serverName),
				},
			},
		},
	}
}
