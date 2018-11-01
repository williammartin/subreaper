# Subreaper

A simple Go library to allow setting, unsetting and fetching of the subreaper attribute. You can read about subreapers at man7, but a quick summary is:

> A subreaper fulfills the role of init(1) for its descendant processes. When a process becomes orphaned (i.e., its immediate parent terminates) then that process will be reparented to the nearest still living ancestor subreaper. Subsequently, calls to getppid() in the orphaned process will now return the PID of the subreaper process, and when the orphan terminates, it is the subreaper process that will receive a SIGCHLD signal and will be able to wait(2) on the process to discover its termination status

# Running the tests

Clone the repo and:

```
ginkgo
```

# Example usage

```go
package main

import (
  "fmt"

  "github.com/williammartin/subreaper"
)

func main() {
  if err := subreaper.Set(); err != nil {
    panic(err)
  }

  isSubreaper, err := subreaper.IsSubreaper()
  if err != nil {
    panic(err)
  }

  fmt.Printf("Subreaper? %t\n", isSubreaper)

  if err := subreaper.Unset(); err != nil {
    panic(err)
  }

  isSubreaper, err = subreaper.IsSubreaper()
  if err != nil {
    panic(err)
  }

  // Should print "Subreaper? false"
  fmt.Printf("Subreaper? %t\n", isSubreaper)
}
```
