package bot

import (
	"github.com/bwmarrin/discordgo"
)

func (b *Bot) handleListServers(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "📋 No servers found in this guild.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleServerStatus(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "📊 Server status feature coming soon!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleStartServer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🟢 Server start feature coming soon!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleStopServer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🔴 Server stop feature coming soon!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleDeleteServer(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🗑️ Server delete feature coming soon!",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleCreateConfirm(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// TODO: Actually create the server in Kubernetes
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "🚀 Server creation started! This will take a few minutes...",
		},
	})
}

func (b *Bot) handleCreateCancel(s *discordgo.Session, i *discordgo.InteractionCreate) {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "❌ Server creation cancelled.",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}

func (b *Bot) handleViewYAML(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// TODO: Generate and show Kubernetes YAML
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "```yaml\n# Generated Kubernetes YAML would appear here\n```",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
}
