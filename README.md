# Scrapingo
Let's make exceptional web crawler/scraper framework to create datasets for ai but not limited to that \
 **Take one feature at a time**


# Features
- Easy to Use
- Clean API
## Phase 1
- Fast (>1k request/sec on a single core)
- Manages request delays and maximum concurrency per domain
- Cookies, Middlewares, robots.txt
- Automatic response decoding to UTF-8 or Automatic encoding of non-unicode responses
- Automatic Data Exporting (JSON, CSV, or custom)
 
## Phase 2
- pagination
- Automatic cookie and session handling
- Sync/async/parallel scraping
- Limit Concurrency (Global/Per Domain)
- Request Delays (Constant/Randomized)

## Phase 3
- JS Rendering
- Caching (Memory/Disk/LevelDB)
- Metrics (Prometheus, Expvar, or custom)
- No CAPTCHA
- Templating to build the workflow faster // let's say you have to collect same kind of data from different 100 sites but store in same db
- notifications
- rotating proxy

# Future Goals
- Distributed scraping | redis & kafka like things ( check out [distrubuted system crawler](https://github.com/henrylee2cn/pholcus) )
- easy integation with 3rd party tools ( like gsheets etc)
- create api for visited site??
- web interface
