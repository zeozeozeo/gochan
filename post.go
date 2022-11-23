package gochan

// Post represents a thread on a board.
type Post struct {
	No              int      `json:"no"`             // The numeric post ID
	Resto           int      `json:"resto"`          // For replies: this is the ID of the thread being replied to. For OP: this value is zero
	Sticky          jsonBool `json:"sticky"`         // If the thread is being pinned to the top of the page
	Closed          jsonBool `json:"closed"`         // If the thread is closed to replies
	Now             string   `json:"now"`            // MM/DD/YY(Day)HH:MM (:SS on some boards), EST/EDT timezone
	Time            jsonTime `json:"time"`           // Time when the post was created
	Name            string   `json:"name"`           // Name user posted with. Defaults to "Anonymous"
	Trip            string   `json:"trip"`           // The user's tripcode, in format: !tripcode or !!securetripcode
	ID              string   `json:"id"`             // The poster's ID (if exists)
	Capcode         string   `json:"capcode"`        // The capcode identifier for a post
	Country         string   `json:"country"`        // Poster's ISO 3166-1 alpha-2 country code
	CountryName     string   `json:"country_name"`   // Poster's country name
	Subject         string   `json:"sub"`            // OP Subject text
	Comment         string   `json:"com"`            // Comment (HTML escaped)
	ImageTime       jsonTime `json:"tim"`            // Time that an image was uploaded
	Filename        string   `json:"filename"`       // Filename as it appeared on the poster's device
	Extension       string   `json:"ext"`            // File extension (filetype)
	Filesize        int64    `json:"fsize"`          // Size of uploaded file in bytes
	FileMD5         string   `json:"md5"`            // 24 character, packed base64 MD5 hash of file
	ImageWidth      int64    `json:"w"`              // Image width dimension
	ImageHeight     int64    `json:"h"`              // Image height dimension
	ThumbnailWidth  int64    `json:"tn_w"`           // Thumbnail image width dimension
	ThumbnailHeight int64    `json:"tn_h"`           // Thumbnail image height dimension
	FileDeleted     jsonBool `json:"filedeleted"`    // If the file was deleted from the post
	Spoiler         jsonBool `json:"spoiler"`        // If the image was spoilered or not
	CustomSpoiler   int      `json:"custom_spoiler"` // The custom spoiler ID for a spoilered image (1-10)
	OmittedPosts    int      `json:"omitted_posts"`  // Number of replies minus the number of previewed replies
	OmittedImages   int      `json:"omitted_images"` // Number of image replies minus the number of previewed image replies
	Replies         int      `json:"replies"`        // Total number of replies to a thread
	Images          int      `json:"images"`         // Total number of image replies to a thread
	BumpLimit       jsonBool `json:"bumplimit"`      // If a thread has reached bumplimit, it will no longer bump
	ImageLimit      jsonBool `json:"imagelimit"`     // If an image has reached image limit, no more image replies can be made
	LastModified    jsonTime `json:"last_modified"`  // The last time the thread was modified (post added/modified/deleted, thread closed/sticky settings modified)
	Tag             string   `json:"tag"`            // The category of .swf upload ("Game", "Loop", etc..)
	SemanticURL     string   `json:"semantic_url"`   // SEO URL slug for thread
	Since4Pass      int      `json:"since4pass"`     // Year 4chan pass bought (any 4 digit year)
	UniqueIPs       int      `json:"unique_ips"`     // Number of unique posters in a thread
	MobileOptimized jsonBool `json:"m_img"`          // Mobile optimized image exists for post
	LastReplies     []Post   `json:"last_replies"`   // Most recent replies to a thread
}
