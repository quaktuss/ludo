package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleResourcesModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.ModalSubmitData()

	cpu := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	memory := data.Components[1].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value
	storage := data.Components[2].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput).Value

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("✅ Resources updated:\n• CPU: %s\n• Memory: %s\n• Storage: %s", cpu, memory, storage),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleSettings1ModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Main settings (1/2) updated!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleSettings2ModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Advanced settings (2/2) updated!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleSingleSettingModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Setting updated via dropdown menu!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleAllSettingsModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ All settings updated (advanced mode)!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleSettingsModalSubmit(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Legacy handler
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "✅ Game settings updated!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
