# Hello-Web-in-Go
- Hello World level
- with simple test case example
- integrate with CI/CD
- test the PR

## setup golang dev env
- [在 Mac book / Centos 7 上設置 Golang 的開發環境 | YIIDTW](https://yiidtw.github.io/blog/2018-05-23-golang-dev-setup/#more)

## usage
```
# how to run
$ git clone https://github.com/yiidtw/hello-web-in-go
$ cd hello-web-in-go
$ go build
$ ./hello-web-in-go &

# open browser or curl the get the result
$ curl localhost:5000
Hello Web

# how to test
$ go test
PASS
ok      hello-web-in-go 0.006s
```
