package main

import (
	"github.com/spinitron-v2-api/mcp-server/config"
	"github.com/spinitron-v2-api/mcp-server/models"
	tools_persona "github.com/spinitron-v2-api/mcp-server/tools/persona"
	tools_playlist "github.com/spinitron-v2-api/mcp-server/tools/playlist"
	tools_show "github.com/spinitron-v2-api/mcp-server/tools/show"
	tools_spin "github.com/spinitron-v2-api/mcp-server/tools/spin"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_persona.CreateGet_personas_idTool(cfg),
		tools_playlist.CreateGet_playlistsTool(cfg),
		tools_playlist.CreateGet_playlists_idTool(cfg),
		tools_show.CreateGet_showsTool(cfg),
		tools_show.CreateGet_shows_idTool(cfg),
		tools_spin.CreateGet_spinsTool(cfg),
		tools_spin.CreateGet_spins_idTool(cfg),
		tools_persona.CreateGet_personasTool(cfg),
	}
}
