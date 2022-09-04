# Zero To Mastery: Golang
This repo contains the source files to accompany the Go language course found at [zerotomastery.io](https://zerotomastery.io/).

## Demos / Exercises
To follow along with the demos & work on the exercises, open the `src/lectures` directory in your IDE.
Demo source files are available in `src/lectures/demo` and coding excercises are available in `src/lectures/exercise`.
When your terminal is in the `src/lectures` directory, demos can be ran using `go run ./demo/demo-name` and exercises can be ran using `go run ./exercise/exercise-name`.

## Solutions
To view solutions for the exercises and projects, checkout the `solutions` branch using `git checkout solutions` (you may need to commit or stash your changes to switch branches).
The `solutions` branch also includes the code seen in the demos.

## Slides
Offline copies of all the slides used in the course are available for download [here](https://cdn.jayson.info/ztm/zerotomastery.io-golang-slides.zip).
File names correspond to the titles of the lectures.

## Pixl
To work on the Pixl project, open the `projects/pixl` directory in your IDE.
You can then execute `go run ./pixl` to run the project.

### Pixl Prerequisites
The Pixl project requires a working installation of `gcc`.
Installation instructions for various operating systems are covered in the course.
`gcc` is a widely available tool, so if your operating system is _not_ detailed in the course (or if you are having trouble), please ping `Jayson#6730` in Discord for installation assistance or ask in `#go` or `#helpme`.

## Keiko
[k6](https://k6.io/) is used to benchmark the performance of the server for this project. To install `k6`, run `go install go.k6.io/k6@latest`.

After installation completes, you can run a benchmark with `k6 run bench.js`.

`Keiko` also uses `SQLite`, which requires `gcc`. If you can run `Pixl`, then you are good to Go :)
