# aoc2023
https://adventofcode.com/2023/about

To use input downloader:
---
1. Login to advent of code and grab cookie with key "session"
2. ```export AOC_SESSION={sessionCookie}```
3. ```go run fetchInput.go {dayNumber}```

To run a day:
---
```go run main.go {dayNumber}```

or

```docker build -t aoc:latest {pathToDockerfile}```

```docker run -it aoc:latest {dayNumber}```

---
To generate day template in goland:
---
While days dir is highlighted: File->New->aoc_2023_day Then enter day number


---