# ludo
Bot Discord for creation of game server

## Quick Start

1. **Create Discord Bot**:
   - Go to https://discord.com/developers/applications
   - Create new application → Bot → Copy token

2. **Create .env file**:
   ```bash
   cp .env.example .env
   # Edit .env with your values
   ```

3. **Build and run**:
   ```bash
   docker build -t ludo-bot .
   docker run --env-file .env ludo-bot
   ```

## Commands

- `/create-server <game> <name>` - Create a new game server
- `/list-servers` - List all servers in this guild
- `/server-status <name>` - Get server status
- `/start-server <name>` - Start a server
- `/stop-server <name>` - Stop a server
- `/delete-server <name>` - Delete a server

## Supported Games

- Project Zomboid