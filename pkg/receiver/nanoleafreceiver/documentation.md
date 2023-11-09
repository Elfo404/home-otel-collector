[comment]: <> (Code generated by mdatagen. DO NOT EDIT.)

# nanoleafreceiver

## Default Metrics

The following metrics are emitted by default. Each of them can be disabled by applying the following configuration:

```yaml
metrics:
  <metric_name>:
    enabled: false
```

### brightness

Brightness

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| percentage | Gauge | Double |

### color_mode

Color mode.

`0`: HS (Hue-Saturation), `1`: CT (Color-Temperature), `2`: Effect

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| mode | Gauge | Int |

### color_temperature

Color temperature.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| kelvin | Gauge | Int |

### hue

Hue

`0` when the panel is off, `1` when the panel is on.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| degrees | Gauge | Double |

### saturation

Color saturation

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| percentage | Gauge | Double |

### status

Status of the panel

`0` when the panel is off, `1` when the panel is on.

| Unit | Metric Type | Value Type |
| ---- | ----------- | ---------- |
| status | Gauge | Int |