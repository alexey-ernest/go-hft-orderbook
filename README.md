[![Build Status](https://travis-ci.com/alexey-ernest/go-hft-orderbook.svg?branch=master)](https://travis-ci.com/alexey-ernest/go-hft-orderbook)

# go-hft-orderbook
Golang implementation of a Limit Order Book (LOB) for high frequency trading in crypto exchanges. Inspired by [this](https://web.archive.org/web/20110219163448/http://howtohft.wordpress.com/2011/02/15/how-to-build-a-fast-limit-order-book/) article.

## Operations

* Add – O(log M) for the first order at a limit, O(1) for all others
* Cancel – O(1)
* GetBestBid/Offer – O(1)
* GetVolumeAtLimit – O(1)

## Performance
* Random generated insertion with limited number of price levels (10K levels) on average MacBook Pro: ~200ns/op or ~5M op/s

## TODO
* Object pool (Done)
* Real data for benchmarks


## Radicle URN
rad:git:hwd1yregaqj5mrp5dgc3gyuu6exszg84zr71su8j1f7t6oe9czxee1zzyzr