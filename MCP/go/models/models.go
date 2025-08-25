package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Spin represents the Spin schema from the OpenAPI specification
type Spin struct {
	Work string `json:"work,omitempty"` // For classical music, the main compositional work the track (see `song` field) belongs to
	Conductor string `json:"conductor,omitempty"` // For classical music, conducor of the ensemble
	Image string `json:"image,omitempty"` // Cover art
	Performers string `json:"performers,omitempty"` // For classical music, featured performing artists, e.g. soloists
	Isrc string `json:"isrc,omitempty"`
	Label string `json:"label,omitempty"` // Record label, i.e. publisher of the sound recording
	Local bool `json:"local,omitempty"` // Is the artist local to the station?
	Request bool `json:"request,omitempty"` // Was the spin requested by a listener?
	Artist string `json:"artist,omitempty"`
	End string `json:"end,omitempty"` // UTC datetime ISO-8601
	Classical bool `json:"classical,omitempty"` // Is the track's metadata schema "classical" rather than "popular"?
	Release string `json:"release,omitempty"`
	Medium string `json:"medium,omitempty"` // Media format in which the sound recording was reased
	Timezone string `json:"timezone,omitempty"` // Station's time zone
	Ensemble string `json:"ensemble,omitempty"` // For classical music, orchestra, performing ensemble, choir, etc.
	Id int `json:"id,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	Composer string `json:"composer,omitempty"`
	Note string `json:"note,omitempty"` // HTML-formatted DJ annotation of the spin, can include images etc.
	Catalog_number string `json:"catalog-number,omitempty"` // Reference number in the record label's catalog
	Release_custom string `json:"release-custom,omitempty"` // Station-specific custom field relating to the release
	Upc string `json:"upc,omitempty"` // Universal Product Code of the release
	NewField bool `json:"new,omitempty"` // Is this a recent release?
	Genre string `json:"genre,omitempty"`
	Artist_custom string `json:"artist-custom,omitempty"` // Station-specific custom field relating to the song's artist
	Start string `json:"start,omitempty"` // UTC datetime ISO-8601
	Released int `json:"released,omitempty"` // Year of initial release
	Va bool `json:"va,omitempty"` // Is it a "Various Artists" release?
	Iswc string `json:"iswc,omitempty"`
	Duration int `json:"duration,omitempty"` // Duration in seconds
	Playlist_id int `json:"playlist_id,omitempty"`
	Song string `json:"song,omitempty"` // Title of the song or track
	Label_custom string `json:"label-custom,omitempty"` // Station-specific custom field relating to the record label
}

// Error represents the Error schema from the OpenAPI specification
type Error struct {
	Code int `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Name string `json:"name,omitempty"`
	Status int `json:"status,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Href string `json:"href,omitempty"`
}

// Show represents the Show schema from the OpenAPI specification
type Show struct {
	Hide_dj bool `json:"hide_dj,omitempty"` // Should the client application hide information about the show's DJs/hosts?
	Image string `json:"image,omitempty"`
	One_off bool `json:"one_off,omitempty"` // Is the show a one-off in the schedule instead of repeating?
	Title string `json:"title,omitempty"` // Program/show title
	Url string `json:"url,omitempty"` // URL to web site for the program/show
	Links map[string]interface{} `json:"_links,omitempty"`
	Category string `json:"category,omitempty"` // Program/show category
	End string `json:"end,omitempty"` // UTC datetime ISO-8601
	Id int `json:"id,omitempty"`
	Since int `json:"since,omitempty"` // Since what year has the program/show existed?
	Description string `json:"description,omitempty"` // HTML-formatted description of the playlist or program/show
	Duration int `json:"duration,omitempty"` // Duration in seconds
	Start string `json:"start,omitempty"` // UTC datetime ISO-8601
	Timezone string `json:"timezone,omitempty"` // Station's time zone
}

// BaseIndexResponse represents the BaseIndexResponse schema from the OpenAPI specification
type BaseIndexResponse struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	Meta Pagination `json:"_meta,omitempty"`
}

// ValidationError represents the ValidationError schema from the OpenAPI specification
type ValidationError struct {
	Field string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

// Pagination represents the Pagination schema from the OpenAPI specification
type Pagination struct {
	Pagecount int `json:"pageCount,omitempty"`
	Perpage int `json:"perPage,omitempty"`
	Totalcount int `json:"totalCount,omitempty"`
	Currentpage int `json:"currentPage,omitempty"`
}

// Persona represents the Persona schema from the OpenAPI specification
type Persona struct {
	Image string `json:"image,omitempty"`
	Name string `json:"name,omitempty"` // On-air DJ/host name
	Since int `json:"since,omitempty"` // Since what year has the DJ/host been at the station?
	Website string `json:"website,omitempty"` // URL to web site for the DJ/host
	Links map[string]interface{} `json:"_links,omitempty"`
	Bio string `json:"bio,omitempty"` // HTML-formatted biography
	Email string `json:"email,omitempty"` // DJ/host's email address
	Id int `json:"id,omitempty"`
}

// Playlist represents the Playlist schema from the OpenAPI specification
type Playlist struct {
	Automation bool `json:"automation,omitempty"` // Was the playlist created playlists created by a radio station automation system?
	Image string `json:"image,omitempty"`
	Links map[string]interface{} `json:"_links,omitempty"`
	Episode_description string `json:"episode_description,omitempty"` // HTML-formatted description of the episode
	Timezone string `json:"timezone,omitempty"` // Station's time zone
	Id int `json:"id,omitempty"`
	Show_id int `json:"show_id,omitempty"`
	End string `json:"end,omitempty"` // UTC datetime ISO-8601
	Start string `json:"start,omitempty"` // UTC datetime ISO-8601
	Hide_dj bool `json:"hide_dj,omitempty"` // Should the client application hide information about the playlist's DJ/host?
	Since int `json:"since,omitempty"` // Since what year has the program/show existed?
	Url string `json:"url,omitempty"` // URL to web site for the playlist or program/show
	Persona_id int `json:"persona_id,omitempty"`
	Category string `json:"category,omitempty"` // Program/show category
	Description string `json:"description,omitempty"` // HTML-formatted description of the playlist or program/show
	Episode_name string `json:"episode_name,omitempty"` // Title of this episode of the program/show
	Title string `json:"title,omitempty"` // Program/show title
	Duration int `json:"duration,omitempty"` // Duration in seconds
}
