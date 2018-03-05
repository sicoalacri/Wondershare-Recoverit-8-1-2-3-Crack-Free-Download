package bosh

import (
	"database/sql"
	"io"

	"github.com/EngineerBetter/concourse-up/config"
	"github.com/EngineerBetter/concourse-up/director"
	"github.com/EngineerBetter/concourse-up/terraform"
)

const cloudConfigFilename = "cloud-config.yml"

// StateFilename is default name for bosh-init state file
const StateFilename = "director-state.json"

// CredsFilename is default name for bosh-init creds file
const CredsFilename = "director-creds.yml"

// Client is a concrete implementation of the IClient interface
type Client struct {
	config   *config.Config
	metadata *terraform.Metadata
	director director.IClient
	db       *sql.DB
	stdout   io.Writer
	stderr   io.Writer
}

// IClient is a client for performing bosh-init commands
type IClient interface {
	Deploy([]byte, []byte, bool) ([]byte, []byte, error)
	Delete([]byte) ([]byte, error)
	Cleanup() error
	Instances() ([]Instance, error)
}

// ClientFactory creates a new IClient
type ClientFactory func(config *config.Config, metadata *terraform.Metadata, director director.IClient, db *sql.DB, stdout, stderr io.Writer) IClient

// NewClient creates a new Client
func NewClient(config *config.Config, metadata *terraform.Metadata, director director.IClient, db *sql.DB, stdout, stderr io.Writer) IClient {
	return &Client{
		config:   config,
		metadata: metadata,
		director: director,
		db:       db,
		stdout:   stdout,
		stderr:   stderr,
	}
}

// Cleanup cleans up temporary files associated with bosh init
func (client *Client) Cleanup() error {
	return client.director.Cleanup()
}
