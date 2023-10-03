## Hidden (need to create) : "private" directory and "private/private.yaml" file, in "config" folder

Path = ***/config/private/private.yaml***

Content =

```yaml
---
lcd:
  url: xxxxxx
  get_timeout_in_seconds: 30
  nb_of_attempts_if_failure: 5
  nb_minutes_of_break_between_attempts: 5
bdd:
  host_name: xxxxxx
  db_name: xxxxxx
  user_name: xxxxxx
  password: xxxxxx
  port: 3306
email:
  host_name: xxxxxx
  smtp_port: 465
  from: xxxxxx
  pwd: xxxxxx
  to: xxxxxx
```