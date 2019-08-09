Repro

```
!? curl -d "foobar" -X POST https://zeitbug.aboodman.now.sh/bug.go
Got 6 bytes of request body:
foobar

!? curl -d "foobar" -X POST -H "Content-Length:" https://zeitbug.aboodman.now.sh/bug.go
Got 0 bytes of request body:
```
