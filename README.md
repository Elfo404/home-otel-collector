# home-otel-collector

An OpenTelemetry Collector for my (and maybe your) home.

## Available receivers

- [Nanoleaf](/pkg/receiver/nanoleafreceiver/README.md)

## Development

### Configuring the collector

The collector is configured using the `config.yaml` file.
An example configuration can be found in `config.yaml.example`, copy the file to `config.yaml` and adjust to your needs.

```bash
cp config.example.yaml config.yaml
```

#### Generate a new token for the Nanoleaf receiver

Put your Nanoleaf in pairing mode by pressing the power button for 5-7 seconds until the panel(s) starts flashing, then send a POST request to the following endpoint:

```bash
curl --location --request POST 'http://<YOUR_NANOLEAF_IP_ADDRESS>:16021/api/v1/new'
```

To get the IP address of your Nanoleaf, you can check your router's DHCP table.

### Building

```bash
make metadata && make build
```

### Running

```bash
make run
```

## Update package level binaries

See [Package Level Binaries](./PACKAGE_LEVEL_BINARIES.md)
