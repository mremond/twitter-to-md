# TODO

- Improve profile crawler by marking crossreferenced profiles as "certified".
- [CRAWLER] Add ability to filter on expected content type (defaults to "text/html", as we are building primarily a
  HTML tool.
- Resolve twitter short url inside embedded tweets.
- Prerender Youtube links
  It should work by embedding content in a way that avoid tracking. We cannot just embed Youtube video snippet.
  The archives converter should:
  - Download the video cover image locally
  - Generate a preview that do not leak information to Youtube.
  As a first step, clicking the video will get the browser to load the Youtube page, but later on, we can
  render the video locally by loading the content from Youtube inline. It will have the same effect as loading the
  Youtube page, but inline. The "on-demand" video play without forced Youtube script embed will be compliant with
  Do not track policy (Like what Medium is doing for example).
  See: https://webdesign.tutsplus.com/tutorials/how-to-lazy-load-embedded-youtube-videos--cms-26743
  http://coffeespace.org.uk/dnt.js 
- Resolve HTML 5 / RDFa prefixes properly when parsing page.
- Generate entries for liked tweets ? They are not included in archive, so requires querying Twitter API to get them.
  We could just generate link.
- Add media types to metadata file
- Add metadata to SQLite index at root dir.
- Rename media file to shorter / more friendly filenames.
- Write initial test suite.
- Refactor / clean-up
- Fix JS header removal to make it more generic (in case there is a part1, etc.)
- Convert smileys to Emoji
- Use similarity to find duplicate post across several source of data
- Remove utm_ parameters from links (used for tracking promo campaigns)

## URL sanitization

- Fix m.engadget.com links (the domain does not exist anymore)
Example:
 http://m.engadget.com/default/article.do?artUrl=http://www.engadget.com/2011/02/08/nokia-ceo-stephen-elop-rallies-troops-in-brutally-honest-burnin/&category=classic&postPage=1
 => http://www.engadget.com/2011/02/08/nokia-ceo-stephen-elop-rallies-troops-in-brutally-honest-burnin/&category=classic&postPage=1

## Roadmap

Other possible services to support for archive cleaning and unification:

- Instagram
- Facebook
- Google+
- Hangout
- Medium
- LinkedIn
- Pinterest
- Flickr
- Quora
- Pocket
- Pinboard
- Dropbox Paper
- Evernote
- Feedly
