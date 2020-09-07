# Light example

This example shows how to interact with a specified light resource. To turn on the light at resource ID 2 (i.e. the first one available on the gateway), you would run:

```
$ go run main.go --host=<IP of your gateway> --apiKey=<API key of the gateway> --id=2 --setState=true --isOn=true
```

If you wish to retrieve all lights, simply omit the 'id' parameter.
If you wish just to retrieve a light, omit the 'setState/setConfig' parameters.

If necessary, it is possible to get both the IP and the API key by following the documentation [here](https://dresden-elektronik.github.io/deconz-rest-doc/getting_started/).
