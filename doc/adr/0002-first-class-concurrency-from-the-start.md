# 2. First class concurrency from the start

Date: 2024-04-20

## Status

Accepted



## Context

Masa is intended to become an integrated system to manage thoughts. People have
thoughts everywhere, and generally don't carry their desktop PC around. Which
means Masa will need to be able to synchronise between multiple devices, laptop,
desktop, phone, and any other devices that it may be possible to use.

Masa is a library as well as providing a CLI tool. The intention is to allow
multiple plugins and extensions, and potentially multiple clients. Which means
that even on a single device it might be necesery at some point to be able to
synchronise between a browser plugin, an IDE plugin, and a CLI tool for example.

## Decision

Go routines will be used where possible, and functionality tested with
concurrency in mind. Go routines are run in the same thread, different threads
and maybe even different processes depending on the go runtime and the device
it's running on. By using go routines Masa functionality will be able to make
efficient use of the compute resources of the device, be that a single core
embedded system (using tinyGo) or the latest desktop with 20+ cores and 30+
threads.

The system design must take into account the intention to support multiple
clients on a single device and ultimately multi-device with syncronisation. This
is especially important for the core Masa library, even thought initially it
will only be used in a simple single process app.

Pragmatic approach should be taken, adding complexity as needed. The key
takeaway from this ADR is to make the overall direction clear. Tactical
solutions may go against this in the short term, but where this is the case they
must be documented with a plan to resolve the tech dept as appropriate.

## Consequences

Implemented well this decision will make it easier to extend Masa later to
multiple devices. Although this does add a bit more complexity upfront when
designing the initial system.

To avoid over complicating the design initially, it will be assumed that
there is only one Masa server running at any point in time per device. Multi
device syncronisation methods are yet to be determined, but will likely
initially start with a single server model and later maybe a transaction log
style syncronisation to allow for offline operation where connections to a
centeralised server are not necesery for basic operations.

The CLI command will initially be the "server" and "client" in the same process.
However as the project progreses this is intended to split. Further decissions
about how that is managed will be made later, but should be as seamless to the
user as possible.

