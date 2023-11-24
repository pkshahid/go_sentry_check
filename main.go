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
    // Set TracesSampleRate to 1.0 to capture 100%
    // of transactions for performance monitoring.
    // We recommend adjusting this value in production,
    TracesSampleRate: 1.0,
  })
  if err != nil {
    log.Fatalf("sentry.Init: %s", err)
  }

  // Flush buffered events before the program terminates.
  defer sentry.Flush(5 * time.Second)
  err = checkTest()
  err2 := "DummyError"
  log.Println(err)

  sentry.ConfigureScope(func(scope *sentry.Scope) {
	  scope.SetLevel(sentry.LevelError)
	//   sentry.CaptureException(err2)
	  sentry.CaptureMessage(err2)
	})

}

func checkTest()error{
	err := checkTest2()
	return err
}

func checkTest2()error{
	err := errors.New("Check Error from checkTest22")
	return err
}