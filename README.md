# dynamarshall

```sh
$ echo '{"foo":"bar"}' | dynamarshall | jq
{
  "foo": {
    "B": null,
    "BOOL": null,
    "BS": null,
    "L": null,
    "M": null,
    "N": null,
    "NS": null,
    "NULL": null,
    "S": "bar",
    "SS": null
  }
}
```
