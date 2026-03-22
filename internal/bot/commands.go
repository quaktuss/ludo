package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

var slashCommands = []*discordgo.ApplicationCommand{
	{
		Name:        "create-server",
		Description: "Create a new game server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "game",
				Description: "Game type",
				Required:    true,
				Choices: []*discordgo.ApplicationCommandOptionChoice{
					{Name: "Project Zomboid", Value: "project-zomboid"},
				},
			},
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Server name",
				Required:    true,
			},
		},
	},
	{
		Name:        "list-servers",
		Description: "List all servers in this guild",
	},
	{
		Name:        "server-status",
		Description: "Get status of a specific server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Server name",
				Required:    true,
			},
		},
	},
	{
		Name:        "start-server",
		Description: "Start a server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Server name",
				Required:    true,
			},
		},
	},
	{
		Name:        "stop-server",
		Description: "Stop a server",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Server name",
				Required:    true,
			},
		},
	},
	{
		Name:        "delete-server",
		Description: "Delete a server permanently",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "name",
				Description: "Server name",
				Required:    true,
			},
		},
	},
}

func (b *Bot) registerCommands() error {
	for _, cmd := range slashCommands {
		if _, err := b.session.ApplicationCommandCreate(b.session.State.User.ID, b.config.Discord.GuildID, cmd); err != nil {
			return fmt.Errorf("failed to create command %s: %w", cmd.Name, err)
		}
	}
	log.Println("Successfully registered all slash commands")
	return nil
}
