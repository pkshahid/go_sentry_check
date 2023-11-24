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