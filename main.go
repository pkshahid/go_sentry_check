package main

import (
  "log"
  "time"
  "github.com/getsentry/sentry-go"
  "errors"
)

func main() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://3cc86389ab265486c0a4801fb15968ef@o4506279358496768.ingest.sentry.io/4506279368720384",
		Debug : true,

		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,

    // Configures whether SDK should generate and attach stack traces to pure
	  // capture message calls.
    AttachStacktrace : true,

    // The sample rate for event submission in the range [0.0, 1.0]. By default,
    // all events are sent. Thus, as a historical special case, the sample rate
    // 0.0 is treated as if it was 1.0. To drop all events, set the DSN to the
    // empty string.
    // SampleRate : 1.0,

    // Enable performance tracing.
    EnableTracing : true,

    // The sample rate for sampling traces in the range [0.0, 1.0].
    TracesSampleRate : 1.0,

    // Used to customize the sampling of traces, overrides TracesSampleRate.
    // TracesSampler TracesSampler

    // The sample rate for profiling traces in the range [0.0, 1.0].
    // This is relative to TracesSampleRate - it is a ratio of profiled traces out of all sampled traces.
    // ProfilesSampleRate : 1.0,

    // List of regexp strings that will be used to match against event's message
    // and if applicable, caught errors type and value.
    // If the match is found, then a whole event will be dropped.
    // IgnoreErrors : make([]string),

    // List of regexp strings that will be used to match against a transaction's
    // name. If a match is found, then the transaction  will be dropped.
    // IgnoreTransactions : make([]string),

    // If this flag is enabled, certain personally identifiable information (PII) is added by active integrations.
    // By default, no such data is sent.
    SendDefaultPII : true,

    // BeforeSend is called before error events are sent to Sentry.
    // Use it to mutate the event or return nil to discard the event.
    // BeforeSend func(event *Event, hint *EventHint) *Event

    // BeforeSendTransaction is called before transaction events are sent to Sentry.
    // Use it to mutate the transaction or return nil to discard the transaction.
    // BeforeSendTransaction func(event *Event, hint *EventHint) *Event

    // Before breadcrumb add callback.
    // BeforeBreadcrumb func(breadcrumb *Breadcrumb, hint *BreadcrumbHint) *Breadcrumb

    // Integrations to be installed on the current Client, receives default
    // integrations.
    // Integrations func([]Integration) []Integration

    // io.Writer implementation that should be used with the Debug mode.
    // DebugWriter io.Writer

    // The transport to use. Defaults to HTTPTransport.
    // Transport Transport

    // The server name to be reported.
    ServerName : "Test Server",

    // The release to be sent with events.
    //
    // Some Sentry features are built around releases, and, thus, reporting
    // events with a non-empty release improves the product experience. See
    // https://docs.sentry.io/product/releases/.
    //
    // If Release is not set, the SDK will try to derive a default value
    // from environment variables or the Git repository in the working
    // directory.
    //
    // If you distribute a compiled binary, it is recommended to set the
    // Release value explicitly at build time. As an example, you can use:
    //
    // 	go build -ldflags='-X main.release=VALUE'
    //
    // That will set the value of a predeclared variable 'release' in the
    // 'main' package to 'VALUE'. Then, use that variable when initializing
    // the SDK:
    //
    // 	sentry.Init(ClientOptions{Release: release})
    //
    // See https://golang.org/cmd/go/ and https://golang.org/cmd/link/ for
    // the official documentation of -ldflags and -X, respectively.
    Release : "v1.0",

    // The dist to be sent with events.
    Dist : "Check",

    // The environment to be sent with events.
    Environment : "Test",

    // Maximum number of breadcrumbs
    // when MaxBreadcrumbs is negative then ignore breadcrumbs.
    MaxBreadcrumbs : 20,

    // Maximum number of spans.
    //
    // See https://develop.sentry.dev/sdk/envelopes/#size-limits for size limits
    // applied during event ingestion. Events that exceed these limits might get dropped.
    // MaxSpans : 20,

    // An optional pointer to http.Client that will be used with a default
    // HTTPTransport. Using your own client will make HTTPTransport, HTTPProxy,
    // HTTPSProxy and CaCerts options ignored.
    // HTTPClient *http.Client

    // An optional pointer to http.Transport that will be used with a default
    // HTTPTransport. Using your own transport will make HTTPProxy, HTTPSProxy
    // and CaCerts options ignored.
    // HTTPTransport http.RoundTripper

    // An optional HTTP proxy to use.
    // This will default to the HTTP_PROXY environment variable.
    // HTTPProxy string

    // An optional HTTPS proxy to use.
    // This will default to the HTTPS_PROXY environment variable.
    // HTTPS_PROXY takes precedence over HTTP_PROXY for https requests.
    // HTTPSProxy string

    // An optional set of SSL certificates to use.
    // CaCerts *x509.CertPool

    // MaxErrorDepth is the maximum number of errors reported in a chain of errors.
    // This protects the SDK from an arbitrarily long chain of wrapped errors.
    //
    // An additional consideration is that arguably reporting a long chain of errors
    // is of little use when debugging production errors with Sentry. The Sentry UI
    // is not optimized for long chains either. The top-level error together with a
    // stack trace is often the most useful information.
    MaxErrorDepth : 20,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	// Flush buffered events before the program terminates.
	defer sentry.Flush(2 * time.Second)

	err = errors.New("Test Error. It's Working.")
	sentry.WithScope(func(scope *sentry.Scope) {

		// set context character
		scope.SetContext("character", map[string]interface{}{
			"name":        "Mighty Fighter",
			"age":         19,
			"attack_type": "melee",
		})

		//add extra data additional
		scope.SetExtra("extra_key", "extra_value")
		//set user id
		scope.SetUser(sentry.User{ID: "test_user"})

		scope.SetLevel(sentry.LevelError)

		sentry.CaptureException(err)
	})  

}