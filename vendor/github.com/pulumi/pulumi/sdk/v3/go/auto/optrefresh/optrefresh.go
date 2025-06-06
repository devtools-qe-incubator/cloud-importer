// Copyright 2016-2020, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package optrefresh contains functional options to be used with stack refresh operations
// github.com/sdk/v2/go/x/auto Stack.Refresh(...optrefresh.Option)
package optrefresh

import (
	"io"

	"github.com/pulumi/pulumi/sdk/v3/go/auto/debug"
	"github.com/pulumi/pulumi/sdk/v3/go/auto/events"
)

// Parallel is the number of resource operations to run in parallel at once during the refresh
// (1 for no parallelism). Defaults to unbounded. (default 2147483647)
func Parallel(n int) Option {
	return optionFunc(func(opts *Options) {
		opts.Parallel = n
	})
}

// ExpectNoChanges will cause the preview to return an error if any changes occur
func ExpectNoChanges() Option {
	return optionFunc(func(opts *Options) {
		opts.ExpectNoChanges = true
	})
}

// ClearPendingCreates will cause the refresh to drop all pending creates from the state
func ClearPendingCreates() Option {
	return optionFunc(func(opts *Options) {
		opts.ClearPendingCreates = true
	})
}

// Message (optional) to associate with the refresh operation
func Message(message string) Option {
	return optionFunc(func(opts *Options) {
		opts.Message = message
	})
}

// Target specifies an exclusive list of resource URNs to refresh
func Target(urns []string) Option {
	return optionFunc(func(opts *Options) {
		opts.Target = urns
	})
}

// ProgressStreams allows specifying one or more io.Writers to redirect incremental refresh stdout
func ProgressStreams(writers ...io.Writer) Option {
	return optionFunc(func(opts *Options) {
		opts.ProgressStreams = writers
	})
}

// ErrorProgressStreams allows specifying one or more io.Writers to redirect incremental refresh stderr
func ErrorProgressStreams(writers ...io.Writer) Option {
	return optionFunc(func(opts *Options) {
		opts.ErrorProgressStreams = writers
	})
}

// EventStreams allows specifying one or more channels to receive the Pulumi event stream
func EventStreams(channels ...chan<- events.EngineEvent) Option {
	return optionFunc(func(opts *Options) {
		opts.EventStreams = channels
	})
}

// DebugLogging provides options for verbose logging to standard error, and enabling plugin logs.
func DebugLogging(debugOpts debug.LoggingOptions) Option {
	return optionFunc(func(opts *Options) {
		opts.DebugLogOpts = debugOpts
	})
}

// UserAgent specifies the agent responsible for the update, stored in backends as "environment.exec.agent"
func UserAgent(agent string) Option {
	return optionFunc(func(opts *Options) {
		opts.UserAgent = agent
	})
}

// Color allows specifying whether to colorize output. Choices are: always, never, raw, auto (default "auto")
func Color(color string) Option {
	return optionFunc(func(opts *Options) {
		opts.Color = color
	})
}

// ShowSecrets configures whether to show config secrets when they appear in the config.
func ShowSecrets(show bool) Option {
	return optionFunc(func(opts *Options) {
		opts.ShowSecrets = &show
	})
}

// Suppress display of periodic progress dots
func SuppressProgress() Option {
	return optionFunc(func(opts *Options) {
		opts.SuppressProgress = true
	})
}

// Suppress display of stack outputs (in case they contain sensitive values)
func SuppressOutputs() Option {
	return optionFunc(func(opts *Options) {
		opts.SuppressOutputs = true
	})
}

// ConfigFile specifies a file to use for configuration values rather than detecting the file name
func ConfigFile(path string) Option {
	return optionFunc(func(opts *Options) {
		opts.ConfigFile = path
	})
}

// Option is a parameter to be applied to a Stack.Refresh() operation
type Option interface {
	ApplyOption(*Options)
}

// ---------------------------------- implementation details ----------------------------------

// Options is an implementation detail
type Options struct {
	// Parallel is the number of resource operations to run in parallel at once
	// (1 for no parallelism). Defaults to unbounded. (default 2147483647)
	Parallel int
	// Message (optional) to associate with the refresh operation
	Message string
	// Return an error if any changes occur during this preview
	ExpectNoChanges bool
	// Clear all pending creates, dropping them from the state
	ClearPendingCreates bool
	// Specify an exclusive of resource URNs to refresh
	Target []string
	// ProgressStreams allows specifying one or more io.Writers to redirect incremental refresh stdout
	ProgressStreams []io.Writer
	// ErrorProgressStreams allows specifying one or more io.Writers to redirect incremental refresh stderr
	ErrorProgressStreams []io.Writer
	// EventStreams allows specifying one or more channels to receive the Pulumi event stream
	EventStreams []chan<- events.EngineEvent
	// DebugLogOpts specifies additional settings for debug logging
	DebugLogOpts debug.LoggingOptions
	// UserAgent specifies the agent responsible for the update, stored in backends as "environment.exec.agent"
	UserAgent string
	// Colorize output. Choices are: always, never, raw, auto (default "auto")
	Color string
	// Show config secrets when they appear.
	ShowSecrets *bool
	// Suppress display of periodic progress dots
	SuppressProgress bool
	// Suppress display of stack outputs (in case they contain sensitive values)
	SuppressOutputs bool
	// Run using the configuration values in the specified file rather than detecting the file name
	ConfigFile string
}

type optionFunc func(*Options)

// ApplyOption is an implementation detail
func (o optionFunc) ApplyOption(opts *Options) {
	o(opts)
}
