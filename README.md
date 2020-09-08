# deconz-go
A library for interfacing with the deCONZ REST API.

This library is designed to be a very basic wrapper around the API offered by the deconz REST API - it is not designed to maintain state. It simply exposes the objects and their operations in an idiomatic Go way. It provides a convenience handler for the deconz websocket to make it easy to build additional software which is maintaining a state-aware view of the gateway.

Currently implemented and tested functionality includes:
1. All methods on the groups endpoint
2. All methods on the lights endpoint
3. All methods on the scenes endpoint
4. Read methods on the sensors endpoint
5. Some methods on the configuration endpoint
6. The websocket endpoint

Adding support for rules, schedules and touchlink should be fairly straightforward, however this work has not yet been undertaken.

The currently supported pieces of the configuration API allow for the creation & deletion of API keys and retrieval of gateway state.

It is possible to see small CLI tools which exercise the above API endpoints in the examples/ directory.
