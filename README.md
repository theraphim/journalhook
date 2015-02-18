# journalhook

systemd-journald hook for logrus

[![wercker status](https://app.wercker.com/status/22d05fb9c9b3442e90aed3ec814e4469/m "wercker status")](https://app.wercker.com/project/bykey/22d05fb9c9b3442e90aed3ec814e4469)

## Use

```go
import "github.com/wercker/journalhook"

journalhook.Enable()
```

Note that this will discard all journal output. Generally when logging to the
journal, your application will be managed by systemd, which will automatically
capture command output and send it to the journal. This way we preempt that
potential log message duplication.

If you'd like not to do that, it's as easy as:

```go

import (
    "github.com/wercker/journalhook"
    "github.com/Sirupsen/logrus"
)

logrus.AddHook(&journalhook.JournalHook{})
```
