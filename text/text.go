package text

import (
  "bytes"
  "regexp"
  "strings"
  "unicode"

  "golang.org/x/text/runes"
  "golang.org/x/text/transform"
  "golang.org/x/text/unicode/norm"
)

var space = regexp.MustCompile(`\s+`)

// normalize make a normalization process on a string.
func Normalize(s string) string {

  // A transformation to remove non-spacing marks.
  markT := func(r rune) bool { return unicode.Is(unicode.Mn, r) }

  // A transformation to remove clean non-letter runes.
  mappingT := func(r rune) rune {
    if !validRune[r] {
      return ' '
    }
    return r
  }

  // A chain of transformation for a string.
  t := transform.Chain(
    norm.NFKD,
    transform.RemoveFunc(markT),
    runes.Map(mappingT),
  )

  r := transform.NewReader(strings.NewReader(s), t)
  buf := new(bytes.Buffer)
  buf.ReadFrom(r)

  trimmed := strings.Trim(space.ReplaceAllString(buf.String(), " "), " ")

  return strings.ToLower(trimmed)
}