# HttpCooker

HttpCooker is a command line tool to perform load testing on REST API endpoints at scale. HTTP calls can be orchestrated to emulate simple peaks in traffic or a heavy request load period. Sets of requests are configurable to be performed synchronously or asynchronously.

## Flags

### url
__Required__
> URL of the endpoint to make calls to.

### method
__Not Required__
Default - "GET"
> HTTP method to make calls to [url](#url) with.

### bearer
__Not Required__
> Authorization bearer token to include in the _Authorization_ header for each HTTP [call](#calls). If not specified, request's _Authorization_ header will not be added.

### groups
__Not Required__
Default - 1
> Largest body of work, represents a batch of [sets](#sets) of HTTP [calls](#calls) to make to a [url](#url).

### sets
__Not Required__
Default - 1
> Number of sets in a [group](#groups)

### calls
__Not Required__
Default - 0
> Number of HTTP requests (calls) to make per set.

### con
__Not Required__
Default - False
> Specifies if [calls](#calls) in a given [set](#sets) should be ran concurrently or not. When ran conncurrently (con == true) all the [calls](#calls) will be ran asynchronously.

### wg
__Not Required__
Default - 0
> Specifies the wait time, in seconds, to wait after each [group](#groups) before running the next group.
