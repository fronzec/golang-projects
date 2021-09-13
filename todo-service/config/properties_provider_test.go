package config

import (
	"os"
	"reflect"
	"testing"
)

func TestNewPropertiesConfigProvider(t *testing.T) {
	err := os.Setenv("PROPERTIES_FILE", "./example/app.properties")
	if err != nil {
		return
	}
	tests := []struct {
		name    string
		want    Provider
		wantErr bool
	}{
		{"read successfully with all configs", &PropertiesConfigProvider{
			Configs{
				configs: map[string]interface{}{
					"db.url":      "",
					"db.name":     "",
					"db.username": "",
					"db.password": "",
				},
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				got, err := NewPropertiesConfigProvider()
				if (err != nil) != tt.wantErr {
					t.Errorf("NewPropertiesConfigProvider() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("NewPropertiesConfigProvider() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestPropertiesConfigProvider_GetFloat(t *testing.T) {
	type fields struct {
		Configs Configs
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				provider := &PropertiesConfigProvider{
					Configs: tt.fields.Configs,
				}
				got, err := provider.GetFloat(tt.args.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetFloat() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("GetFloat() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestPropertiesConfigProvider_GetInt(t *testing.T) {
	type fields struct {
		Configs Configs
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				provider := &PropertiesConfigProvider{
					Configs: tt.fields.Configs,
				}
				got, err := provider.GetInt(tt.args.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetInt() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("GetInt() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestPropertiesConfigProvider_GetInt64(t *testing.T) {
	type fields struct {
		Configs Configs
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				provider := &PropertiesConfigProvider{
					Configs: tt.fields.Configs,
				}
				got, err := provider.GetInt64(tt.args.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetInt64() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("GetInt64() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestPropertiesConfigProvider_GetString(t *testing.T) {
	type fields struct {
		Configs Configs
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				provider := &PropertiesConfigProvider{
					Configs: tt.fields.Configs,
				}
				got, err := provider.GetString(tt.args.key)
				if (err != nil) != tt.wantErr {
					t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				if got != tt.want {
					t.Errorf("GetString() got = %v, want %v", got, tt.want)
				}
			},
		)
	}
}

func TestPropertiesConfigProvider_loadConfig(t *testing.T) {
	type fields struct {
		Configs Configs
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				provider := &PropertiesConfigProvider{
					Configs: tt.fields.Configs,
				}
				if err := provider.loadConfig(); (err != nil) != tt.wantErr {
					t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
				}
			},
		)
	}
}

func TestPropertiesConfigProvider_logConfigs(t *testing.T) {
	type fields struct {
		Configs Configs
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(
			tt.name, func(t *testing.T) {
				_ = &PropertiesConfigProvider{
					Configs: tt.fields.Configs,
				}
			},
		)
	}
}
