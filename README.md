## Alertmanager

See original alertmanager doc at : https://github.com/prometheus/alertmanager

## What is new?
### MongoDB Configuration 

You can add MongoDB in alertmanger.yaml.
For example:
```
receivers:
  - name: "webhook"
    webhook_configs:
      - url: "https://webhook.site/fb1c57e7-f628-429d-9d82-388bd8beec34"
  - name: "mongodb"
    mongodb_configs:
      url: "10.136.134.218"
      port: "27017"
      database: "test"
      collection: "alertmanager"
      username: "root"
      password: "123456"
      search: true
  - name: "default"

```
Then every alert will be saved as document in MongoDB.

### Related API
* /api/v2/stored_alerts
  
  Get stored alerts in MongoDB

* /api/v2/stored_alerts/severity_days
  
  Get counts of alerts in days

See more deatails in [oepnapi.yaml](https://github.com/edwardzhanged/alertmanager-plus/blob/main/api/v2/openapi.yaml)
