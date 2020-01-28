# Who is on-call?

This is a mocking exercise. We have a seemingly working software, that can announce who is on-call this week, and it can add some drama.

We have an end-to-end test in place, but it has some limitations. It works with the "real" PagerDuty API, so it's very fragile (we have to change the expectations every week), and it's not trivial to test all the features of the software.

Your task is to introduce unit testing, decouple it from the real PagerDuty client, and test all the important features of the software.

These are:
- it announces the one on-call (uppercased) with some introductory text
- it gives an error if it can't fetch the one on-call
- it pauses for a bit after every line
- and it takes a bigger break after the final line to add more drama

## Setup

Get the project

```
go get github.com/gaborszakacs/who-is-on-call
```

See the program running

```
go run main.go
```

Run the tests including the E2E (`-count=1` needed to disable cache)

```
go test ./... -tags=e2e -count=1
```

Run only the unit tests you're writing

```
go test ./...
```

Auto-run tests on file change and add red/green color, aliased as `gt` (`brew install fswatch` needed)

```
alias gt=$'fswatch -e ".*" -i "\\.go$" . | xargs -n1 -I {} sh -c "clear && printf \'\\e[3J\';go test ./... | sed \'\'/ok/s//\C-[[32mPASS\C-[[0m/\'\' | sed \'\'/FAIL/s//\C-[[31mFAIL\C-[[0m/\'\'"'
```
