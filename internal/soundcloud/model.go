package soundcloud

import "time"

type SearchResult struct {
	Collection   []Collection `json:"collection"`
	TotalResults int          `json:"total_results"`
	NextHref     string       `json:"next_href"`
	QueryUrn     string       `json:"query_urn"`
}

type Format struct {
	Protocol string `json:"protocol"`
	MimeType string `json:"mime_type"`
}

type Transcodings struct {
	URL                 string `json:"url"`
	Preset              string `json:"preset"`
	Duration            int    `json:"duration"`
	Snipped             bool   `json:"snipped"`
	Format              Format `json:"format"`
	Quality             string `json:"quality"`
	IsLegacyTranscoding bool   `json:"is_legacy_transcoding"`
}

type Media struct {
	Transcodings []Transcodings `json:"transcodings"`
}

type Product struct {
	ID string `json:"id"`
}

type CreatorSubscriptions struct {
	Product Product `json:"product"`
}

type CreatorSubscription struct {
	Product Product `json:"product"`
}

type Visuals struct {
	Urn      string    `json:"urn"`
	Enabled  bool      `json:"enabled"`
	Visuals  []Visuals `json:"visuals"`
	Tracking any       `json:"tracking"`
}

type Badges struct {
	Pro            bool `json:"pro"`
	CreatorMidTier bool `json:"creator_mid_tier"`
	ProUnlimited   bool `json:"pro_unlimited"`
	Verified       bool `json:"verified"`
}

type User struct {
	AvatarURL            string                 `json:"avatar_url"`
	City                 string                 `json:"city"`
	CommentsCount        int                    `json:"comments_count"`
	CountryCode          string                 `json:"country_code"`
	CreatedAt            time.Time              `json:"created_at"`
	CreatorSubscriptions []CreatorSubscriptions `json:"creator_subscriptions"`
	CreatorSubscription  CreatorSubscription    `json:"creator_subscription"`
	Description          string                 `json:"description"`
	FollowersCount       int                    `json:"followers_count"`
	FollowingsCount      int                    `json:"followings_count"`
	FirstName            string                 `json:"first_name"`
	FullName             string                 `json:"full_name"`
	GroupsCount          int                    `json:"groups_count"`
	ID                   int                    `json:"id"`
	Kind                 string                 `json:"kind"`
	LastModified         time.Time              `json:"last_modified"`
	LastName             string                 `json:"last_name"`
	LikesCount           int                    `json:"likes_count"`
	PlaylistLikesCount   int                    `json:"playlist_likes_count"`
	Permalink            string                 `json:"permalink"`
	PermalinkURL         string                 `json:"permalink_url"`
	PlaylistCount        int                    `json:"playlist_count"`
	RepostsCount         any                    `json:"reposts_count"`
	TrackCount           int                    `json:"track_count"`
	URI                  string                 `json:"uri"`
	Urn                  string                 `json:"urn"`
	Username             string                 `json:"username"`
	Verified             bool                   `json:"verified"`
	Visuals              Visuals                `json:"visuals"`
	Badges               Badges                 `json:"badges"`
	StationUrn           string                 `json:"station_urn"`
	StationPermalink     string                 `json:"station_permalink"`
}

type PublisherMetadata struct {
	ID              int    `json:"id"`
	Urn             string `json:"urn"`
	Artist          string `json:"artist"`
	AlbumTitle      string `json:"album_title"`
	ContainsMusic   bool   `json:"contains_music"`
	UpcOrEan        string `json:"upc_or_ean"`
	Isrc            string `json:"isrc"`
	Explicit        bool   `json:"explicit"`
	PLine           string `json:"p_line"`
	PLineForDisplay string `json:"p_line_for_display"`
	CLine           string `json:"c_line"`
	CLineForDisplay string `json:"c_line_for_display"`
	ReleaseTitle    string `json:"release_title"`
}

type Collection struct {
	ArtworkURL         string            `json:"artwork_url"`
	Caption            any               `json:"caption"`
	Commentable        bool              `json:"commentable"`
	CommentCount       int               `json:"comment_count"`
	CreatedAt          time.Time         `json:"created_at"`
	Description        string            `json:"description"`
	Downloadable       bool              `json:"downloadable"`
	DownloadCount      int               `json:"download_count"`
	Duration           int               `json:"duration"`
	FullDuration       int               `json:"full_duration"`
	EmbeddableBy       string            `json:"embeddable_by"`
	Genre              string            `json:"genre"`
	HasDownloadsLeft   bool              `json:"has_downloads_left"`
	ID                 int               `json:"id"`
	Kind               string            `json:"kind"`
	LabelName          any               `json:"label_name"`
	LastModified       time.Time         `json:"last_modified"`
	License            string            `json:"license"`
	LikesCount         int               `json:"likes_count"`
	Permalink          string            `json:"permalink"`
	PermalinkURL       string            `json:"permalink_url"`
	PlaybackCount      int               `json:"playback_count"`
	Public             bool              `json:"public"`
	PublisherMetadata  PublisherMetadata `json:"publisher_metadata,omitempty"`
	PurchaseTitle      any               `json:"purchase_title"`
	PurchaseURL        any               `json:"purchase_url"`
	ReleaseDate        time.Time         `json:"release_date"`
	RepostsCount       int               `json:"reposts_count"`
	SecretToken        any               `json:"secret_token"`
	Sharing            string            `json:"sharing"`
	State              string            `json:"state"`
	Streamable         bool              `json:"streamable"`
	TagList            string            `json:"tag_list"`
	Title              string            `json:"title"`
	URI                string            `json:"uri"`
	Urn                string            `json:"urn"`
	UserID             int               `json:"user_id"`
	Visuals            any               `json:"visuals"`
	WaveformURL        string            `json:"waveform_url"`
	DisplayDate        time.Time         `json:"display_date"`
	Media              Media             `json:"media"`
	StationUrn         string            `json:"station_urn"`
	StationPermalink   string            `json:"station_permalink"`
	TrackAuthorization string            `json:"track_authorization"`
	MonetizationModel  string            `json:"monetization_model"`
	Policy             string            `json:"policy"`
	User               User              `json:"user"`
}
