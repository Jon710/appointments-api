# appointments-api

This is an API that allows patients scheduling appointment with their doctor of choice.
I will simulate patients trying to schedule appointments concurrenctly and solve it in both optimistic (version-based/timestamp in database) and pessimistic (goroutines, db or go locks) ways.
I intend to use artillery (or something like that) to load test the API.

## references
- https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency
- https://www.youtube.com/watch?v=URwmzTeuHdk postgres locks
- https://www.youtube.com/watch?v=4F-WiPFrPsA pessimistic with Arpit
- https://betterstack.com/community/guides/scaling-nodejs/nodejs-caching-redis/ (it has an example of artillery)
- https://github.com/arpitbbhayani/concurrency-in-depth/blob/master/05-concurrent-thread-safe-queue/main.go
- https://medium.com/tech-at-wildlife-studios/write-backend-systems-50aae8db849e
- https://getstream.io/blog/how-we-test-go-at-stream/
- https://getstream.io/blog/building-a-performant-api-using-go-and-cassandra/