package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) onInteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		b.handleSlashCommand(s, i)
	case discordgo.InteractionMessageComponent:
		b.handleButtonInteraction(s, i)
	case discordgo.InteractionModalSubmit:
		b.handleModalSubmit(s, i)
	}
}

func (b *Bot) handleSlashCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ApplicationCommandData()

	if !b.hasPermission(i.Member, data.Name) {
		b.respondError(s, i, "You don't have permission to use this command")
		return
	}

	switch data.Name {
	case "create-server":
		b.handleCreateServer(s, i)
	case "list-servers":
		b.handleListServers(s, i)
	case "server-status":
		b.handleServerStatus(s, i)
	case "start-server":
		b.handleStartServer(s, i)
	case "stop-server":
		b.handleStopServer(s, i)
	case "delete-server":
		b.handleDeleteServer(s, i)
	default:
		b.respondError(s, i, "Unknown command")
	}
}

func (b *Bot) handleButtonInteraction(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()
	customID := data.CustomID

	switch data.ComponentType {
	case discordgo.SelectMenuComponent:
		b.handleSelectMenu(s, i)
	case discordgo.ButtonComponent:
		b.routeButtonInteraction(s, i, customID)
	}
}

func (b *Bot) handleModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	customID := i.ModalSubmitData().CustomID

	switch {
	case hasPrefix(customID, "resources_modal_"):
		b.handleResourcesModalSubmit(s, i)
	case hasPrefix(customID, "settings1_modal_"):
		b.handleSettings1ModalSubmit(s, i)
	case hasPrefix(customID, "settings2_modal_"):
		b.handleSettings2ModalSubmit(s, i)
	case hasPrefix(customID, "single_setting_modal_"):
		b.handleSingleSettingModalSubmit(s, i)
	case hasPrefix(customID, "all_settings_modal_"):
		b.handleAllSettingsModalSubmit(s, i)
	case hasPrefix(customID, "settings_modal_"): // Legacy support
		b.handleSettingsModalSubmit(s, i)
	}
}

func (b *Bot) hasPermission(member *discordgo.Member, commandName string) bool {
	// TODO: Implement proper role-based permissions
	return true
}

func (b *Bot) respondError(s *discordgo.Session, i *discordgo.InteractionCreate, message string) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "❌ " + message,
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}
