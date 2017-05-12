[![GoDoc](https://godoc.org/github.com/mdirkse/i3ipc-go?status.svg)](http://godoc.org/github.com/mdirkse/i3ipc-go/)
[![Build Status](https://travis-ci.org/mdirkse/i3ipc-go.svg?branch=master)](https://travis-ci.org/mdirkse/i3ipc-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/mdirkse/i3ipc-go)](https://goreportcard.com/report/github.com/mdirkse/i3ipc-go)
[![codecov](https://codecov.io/gh/mdirkse/i3ipc-go/branch/master/graph/badge.svg)](https://codecov.io/gh/mdirkse/i3ipc-go)
[![BCH compliance](https://bettercodehub.com/edge/badge/mdirkse/i3ipc-go?branch=master)](https://bettercodehub.com/)

i3ipc
=====

Overview
--------
i3ipc is a golang library for convenient access to the IPC API of the [i3 window manager](http://i3wm.org). If you want to take a look at the documentation then head to [http://godoc.org/github.com/mdirkse/i3ipc-go/][doc].

Compatibility
-------------
This library can be used with the i3 IPC as it is as of at least version [4.13](https://github.com/i3/i3/releases/tag/4.13). However, according to the i3 maintainers:
> The IPC isn't versioned. It can change with every release and usually does in one way or another. We only try to avoid breaking changes to the documented values since we consider those stable, but with good enough reason even those can change.

We'll do our best to make this library track the i3 IPC as closely as possible, but if you find anything missing (or broken) please file an issue or (even better) submit a pull request.

Usage
-----
Thanks to Go's built-in git support, you can start using i3ipc with a simple

    import "github.com/mdirkse/i3ipc"

For everything except subscriptions, you will want to create an IPCSocket over which the communication will take place. This object has methods for all message types that i3 will accept, though some might be split into multiple methods (eg. *Get_Bar_Config*). You can create such a socket quite easily:

    ipcsocket, err := i3ipc.GetIPCSocket()

As a simple example of what you could do next, let's get the version of i3 over our new socket:

    version, err := ipcsocket.GetVersion()

For further commands, refer to `go doc` or use the aforementioned [website][doc].

### Subscriptions
i3ipc handles subscriptions in a convenient way: you don't have to think about managing the socket or watch out for unordered replies. The appropriate method simply returns a channel from which you can read Event objects.

Here's a simple example - we start the event listener, we subscribe to workspace events, then simple print all of them as we receive them:

    i3ipc.StartEventListener()
    ws_events, err := i3ipc.Subscribe(i3ipc.I3WorkspaceEvent)
    for {
        event := <-ws_events
        fmt.Printf("Received an event: %v\n", event)
    }

i3ipc currently has no way of subscribing to multiple event types over a single channel. If you want this, you can simply create multiple subscriptions, then demultiplex those channels yourself - `select` is your friend.

Credits
-------
Many thanks to [proxypoke](https://github.com/proxypoke) for originally starting this project.

[doc]: http://godoc.org/github.com/mdirkse/i3ipc-go/
