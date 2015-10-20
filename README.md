# PacktpubFreeAlert
## Prepare to use
* register https://www.pushbullet.com
* get your access token at https://www.pushbullet.com/#settings
* put your access token to configAlert.gcfg
` apikey = Bearer [your access token]`
* set cron with crontab style

```
Field name   | Mandatory? | Allowed values  | Allowed special characters
----------   | ---------- | --------------  | --------------------------
Seconds      | Yes        | 0-59            | * / , -
Minutes      | Yes        | 0-59            | * / , -
Hours        | Yes        | 0-23            | * / , -
Day of month | Yes        | 1-31            | * / , - ?
Month        | Yes        | 1-12 or JAN-DEC | * / , -
Day of week  | Yes        | 0-6 or SUN-SAT  | * / , - ?
```
