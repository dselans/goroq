goroq 
=====
Goroq is a tool for assisting with Go testing. It will watch over a number of Go project dirs and automatically run tests, whenever a change is detected.

The output of the tests can be configured to be written to any file, so that you can keep a cotinuous `tail -f` running in a separate window.

### Usage
Using goroq is super simple. Run it with no arguments and it will do all of the heavy lifting.

```
$ go get github.com/dselans/goroq
$ goroq
>> First time running goroq
>> Updated goroq configuration in ~/.goroqrc
>> Monitoring test changes in <current_dir>
>> Outputting test results to <current_dir>/goroq.log
>> Daemonizing...
$
```

Additional params:
```
$ goreq -h
Usage: ./goreq [-h|-v|] [-d directory] [-o output_file] [-c config_file]
```

### Misc
Goroq uses `inotify` for detecting file changes (both in tests and its own configuration).

While similar projects exist ([goconvey](http://goconvey.co/), [looper](https://github.com/nathany/looper)), this seemed like a fun project to get some more Go experience.
