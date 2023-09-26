package dashlist

// Option represents an option that can be used to configure a graph panel.
type Option func(dashlist *DashList) error

// DashList represents a dashboard list panel.
type DashList struct {
	Builder *sdk.Panel
}

// New creates a new dashboard list panel.
func New(title string, options ...Option) (*DashList, error) {
	panel := &DashList{Builder: sdk.NewDashlist(title)}
	panel.Builder.IsNew = false
	panel.Builder.DashlistPanel.Limit = 10

	for _, opt := range append(defaults(), options...) {
		if err := opt(panel); err != nil {
			return nil, err
		}
	}

	return panel, nil
}

// NoTitle removes the default text and sets to an empty string
func NoTitle() Option {
	return func(text *Text) error {
		text.Builder.Title = ""
		return nil
	}
}

// ShowHeadings toggles headings to be shown
func ShowHeadings() Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.Headings = true
		return nil
	}
}

// ShowSearch toggles dashboards from search to be shown
func ShowSearch() Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.Search = true
		return nil
	}
}

// ShowStarred toggles starred dashboards to be shown
func ShowStarred() Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.Starred = true
		return nil
	}
}

// ShowRecent toggles recently viewed dashboards to be shown
func ShowRecent() Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.Recent = true
		return nil
	}
}

// MaxItems sets the max number of dashboards visible
func MaxItems(limit int) Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.Limit = limit
		return nil
	}
}

// Query sets query string
func Query(query string) Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.Query = query
		return nil
	}
}

// Folder sets folder to operate on
func Folder(id int) Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.FolderID = id
		return nil
	}
}

// Tags sets dashboard tags to show
func Tags(tags []string) Option {
	return func(text *Text) error {
		text.Builder.DashlistPanel.Tags = tags
		return nil
	}
}
