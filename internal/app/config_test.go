package app

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.HTTPAddr != ":8080" {
		t.Fatalf("expected default HTTP address %q, got %q", ":8080", cfg.HTTPAddr)
	}
}

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		cfg     Config
		wantErr bool
	}{
		{
			name: "valid configuration",
			cfg: Config{
				HTTPAddr: ":8080",
			},
			wantErr: false,
		},
		{
			name:    "empty HTTP address",
			cfg:     Config{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.cfg.Validate()

			if tt.wantErr && err == nil {
				t.Fatal("expected validation error")
			}

			if !tt.wantErr && err != nil {
				t.Fatalf("expected no validation error, got %v", err)
			}
		})
	}
}

func TestNewConfiguredHTTPServer(t *testing.T) {
	cfg := Config{
		HTTPAddr: ":9090",
	}

	server := NewConfiguredHTTPServer(cfg)

	if server.Addr != ":9090" {
		t.Fatalf("expected server address %q, got %q", ":9090", server.Addr)
	}

	if server.Handler == nil {
		t.Fatal("expected server handler to be configured")
	}
}
