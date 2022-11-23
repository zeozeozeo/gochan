# gochan

Fully-featured [4chan API](https://github.com/4chan/4chan-API) wrapper for Go with zero dependencies.

# Info

- The API ratelimit is 1 second ([API rules](https://github.com/4chan/4chan-API#api-rules))
- Thread updating ratelimit is 15 seconds by default, the minimum is 10 seconds ([API rules](https://github.com/4chan/4chan-API#api-rules))
- Uses [If-Modified-Since](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Modified-Since) for update requests ([API rules](https://github.com/4chan/4chan-API#api-rules))
- SSL is enabled by default, can be disabled with `session.SSL = false` ([API rules](https://github.com/4chan/4chan-API#api-rules))