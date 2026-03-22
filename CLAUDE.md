# CLAUDE.md

This file provides guidance to Claude Code when working with code in this repository.

## Project Overview

**Ludo** is a Discord bot written in Go for creating and managing game servers (Kubernetes-backed) via Discord slash commands and interactive UI (buttons, modals, select menus).

## Technology Stack

- **Language**: Go 1.21+
- **Discord**: `github.com/bwmarrin/discordgo`
- **Kubernetes**: `k8s.io/client-go`
- **Database**: AWS DynamoDB via `github.com/aws/aws-sdk-go`
- **Config format**: JSON (game templates)

## Project Structure

```
cmd/
  bot/
    main.go               # Entry point — run with: go run ./cmd/bot

internal/
  config/
    config.go             # Env-based config loading (Discord, AWS, K8s, Bot)
  bot/
    bot.go                # Bot struct, New(), Start(), Stop(), onReady()
    commands.go           # Slash command definitions + registerCommands()
    handlers.go           # Interaction router (slash / button / modal)
    handler_create.go     # /create-server flow, embed & UI components
    handler_modify.go     # Modify buttons, select menu, settings modals
    handler_modals.go     # Modal submit handlers
    handler_servers.go    # Server management commands + confirm/cancel/yaml
  templates/
    game.go               # GameTemplate structs + TemplateManager (loads JSON)
  k8s/
    client.go             # Kubernetes client stub (TODO: lifecycle methods)
  database/
    client.go             # DynamoDB client stub (TODO: CRUD methods)

configs/
  games/
    project-zomboid.json  # Game template definition (resources, ports, settings)

k8s/
  helm-charts/            # Helm chart overrides (empty — TODO)
  manifests/              # Raw K8s manifests (empty — TODO)

main.go                   # Ignored (//go:build ignore) — real entry: cmd/bot/main.go
Dockerfile                # Multi-stage Alpine build, runs as non-root, port 8080
.env.example              # Required environment variables
```

## Running the Bot

```bash
cp .env.example .env
# Fill in DISCORD_TOKEN and DISCORD_GUILD_ID at minimum
go run ./cmd/bot
```

## Development Notes

### What works
- Full Discord interaction routing (commands, buttons, select menus, modals)
- Interactive server creation flow with resource/settings configuration
- Game template loading from `configs/games/*.json`
- Config validation (DISCORD_TOKEN is required)

### What's stubbed / TODO
- `internal/k8s/client.go` — add Deploy, Delete, Scale, Status methods
- `internal/database/client.go` — add SaveServer, GetServer, ListServers, DeleteServer
- `internal/bot/handler_servers.go` — all handlers return placeholders
- `internal/bot/handlers.go` — `hasPermission()` always returns true
- `internal/bot/handler_create.go` — `handleCreateConfirm()` doesn't create anything yet
- Helm chart integration

### Adding a new game
1. Create `configs/games/<game-name>.json` following the Project Zomboid template
2. Add a new `Choices` entry in `slashCommands` in `internal/bot/commands.go`
