package soundcloud

type SoundCloud interface {
	Search(query string) ([]Video, error)
}

type soundcloud struct {
	id string
}

func New() (SoundCloud, error) {
	id, err := GetClientID()
	if err != nil {
		return nil, err
	}

	return &soundcloud{id: id}, nil
}
