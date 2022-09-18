package kua

import (
  browser "github.com/EDDYCJY/fake-useragent"
  "github.com/mileusna/useragent"
)

func Random() string {
  return browser.Random()
}

func Parse(ua string) useragent.UserAgent {
  return useragent.Parse(ua)
}
