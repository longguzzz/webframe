```
// curl "http://localhost:54321/login" -X POST -d 'username=user&password=111'
// 404 NOT FOUND: /login
// curl "http://localhost:54321/LOGIN" -X POST -d 'username=user&password=111'
// {"password":"111","username":"user"}

// curl "http://localhost:54321/hello?name=testname"
// para: testname%
// zsh有问题，会比预期的多个百分号，测试不同的平台、简化操作……换成bash后正常
// curl "http://localhost:54321/"
// <h1>root</h1>%
// zsh有问题，会比预期的多个百分号，测试不同的平台、简化操作……换成bash后正常
// curl -i "http://localhost:54321/"
// HTTP/1.1 200 OK
// Content-Type: text/html
// Date: Mon, 08 May 2023 04:36:00 GMT
// Content-Length: 13

// <h1>root</h1>%
```