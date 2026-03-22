package bot

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
	"ludo-bot/internal/config"
	"ludo-bot/internal/templates"
)

type Bot struct {
	session         *discordgo.Session
	config          *config.Config
	templateManager *templates.TemplateManager
}

func New(cfg *config.Config, configDir string) (*Bot, error) {
	session, err := discordgo.New("Bot " + cfg.Discord.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to create Discord session: %w", err)
	}

	templateManager := templates.NewTemplateManager(configDir)
	if err := templateManager.LoadTemplates(); err != nil {
		return nil, fmt.Errorf("failed to load game templates: %w", err)
	}

	bot := &Bot{
		session:         session,
		config:          cfg,
		templateManager: templateManager,
	}

	session.AddHandler(bot.onReady)
	session.AddHandler(bot.onInteractionCreate)

	return bot, nil
}

func (b *Bot) Start() error {
	if err := b.session.Open(); err != nil {
		return fmt.Errorf("failed to open Discord session: %w", err)
	}
	if err := b.registerCommands(); err != nil {
		return fmt.Errorf("failed to register commands: %w", err)
	}
	return nil
}

func (b *Bot) Stop() {
	if b.session != nil {
		b.session.Close()
	}
}

func (b *Bot) onReady(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Ludo bot is ready! Logged in as %s", r.User.String())

	err := s.UpdateStatusComplex(discordgo.UpdateStatusData{
		Status: "online",
		Activities: []*discordgo.Activity{
			{
				Name: "Managing game servers",
				Type: discordgo.ActivityTypeWatching,
			},
		},
	})
	if err != nil {
		log.Printf("Error setting bot status: %v", err)
	}
}
