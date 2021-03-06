package types

import (
	"fmt"
	"time"
)

const (
	// DockerPrefix is the URL-like schema prefix used for Docker image references.
	DockerPrefix = "docker://"
)

// Registry is a service providing repositories.
type Registry interface {
	Repositories() []Repository
	Repository(ref string) Repository
	Lookup(term string) []Image // docker registry v1 only AFAICT, v2 can be built hacking with Images()
}

// Repository is a set of images.
type Repository interface {
	Images() []Image
	Image(ref string) Image // ref == image name w/o registry part
}

// Image is a Docker image in a repository.
type Image interface {
	// ref to repository?
	Layers(layers ...string) error // configure download directory? Call it DownloadLayers?
	Manifest() (ImageManifest, error)
	RawManifest(version string) ([]byte, error)
	DockerTar() ([]byte, error) // ??? also, configure output directory
}

// ImageManifest is the interesting subset of metadata about an Image.
// TODO(runcom)
type ImageManifest interface {
	String() string
}

// DockerImageManifest is a set of metadata describing Docker images and their manifest.json files.
// Note that this is not exactly manifest.json, e.g. some fields have been added.
type DockerImageManifest struct {
	Name          string
	Tag           string
	Digest        string
	RepoTags      []string
	Created       time.Time
	DockerVersion string
	Labels        map[string]string
	Architecture  string
	Os            string
	Layers        []string
}

func (m *DockerImageManifest) String() string {
	return fmt.Sprintf("%s:%s", m.Name, m.Tag)
}
