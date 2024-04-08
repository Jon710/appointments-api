# appointments-api

This is an API that allows patients scheduling appointment with their doctor of choice.
I will be simulating patients trying to schedule appointments concurrenctly and solve this problem in both optimistic (version-based/timestamp in database) and pessimistic (go routines, locks) ways.
I intend to use artillery (or something like that) to load test the API.

## references
- https://stackoverflow.com/questions/48930732/how-to-test-unlikely-concurrent-scenarios?rq=3
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency
- https://aprendagolang.com.br/2022/07/20/mutex-ou-channels-como-resolver-race-condition/
- https://www.youtube.com/watch?v=Ya5KRFrwPug mutexes x channels anthony gg
- https://www.youtube.com/watch?v=URwmzTeuHdk postgres locks
- https://www.youtube.com/watch?v=4F-WiPFrPsA pessimistic with Arpit
- https://betterstack.com/community/guides/scaling-nodejs/nodejs-caching-redis/ (it has an example of artillery)
- https://github.com/arpitbbhayani/concurrency-in-depth/blob/master/05-concurrent-thread-safe-queue/main.go
- https://medium.com/tech-at-wildlife-studios/write-backend-systems-50aae8db849e
- https://getstream.io/blog/how-we-test-go-at-stream/
- https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/
- https://www.youtube.com/watch?v=7zDl-aPW9sg&list=PL0xRBLFXXsP5yp0V2vHLrW9Wftt3uU34H