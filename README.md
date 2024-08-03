# web-crawler

#### A tiny web-crawler which looks for the links, extract and prints them concurrently to the Terminal output

The web-crawler is designed in an asynchronous way: based on channels, wait group, mutexes and goroutines

The web-crawler icon is generated by AI.
![crawler-icon](https://i.ibb.co/F60L7Kn/create-an-icon-of-web-crawler-written-in-golang-programming-language.jpg)

### Usage

Run with the default target url:
- ``make run``

Run with one of two pre-defined target urls:
- ``make run-target-1`` or ``make run-target-2``

Run with your custom url:
- ``make run -url=https://www.my-url.com`` -- !important: you have to stick to the format as in the example

Download a binary [here][bin]

#### Other
There's a posibility to build it on your own:
- ``make build``

Or run linter:
- ``make lint``


[bin]: https://github.com/RSheremeta/web-crawler/releases/tag/1.0