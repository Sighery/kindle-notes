# Steps to connect and read/write BLE characteristic

### Setup

```
$ ace_bt_cli
ACEBTCLI BT Client :: start
p_data (size:27) = 1B 00 00 00 00 00 00 00 00 00 00 61 63 65 5F 62
                   74 5F 63 6C 69 00 00 00 00 00 00
ACEBTCLI BT Client :: opened session with 0x18570

>: ble regble
ACEBTCLI CLI callback : aceBtCli_bleRegCallback() status: 0
ACEBTCLI Register BLE Client Success

>: ble regGattc
ACEBTCLI Register Gatt Client Success
```


### Connection

```
>: ble connect 2C:CF:67:B8:DC:3F 2 10 true
ACEBTCLI str: 2CCF67B8DC3F
ACEBTCLI aceBtCli_aclStateChangedCallback() status:0 addr:2C:CF:67:B8:DC:3F state:0, transport:0, reason:0
ACEBTCLI CLI callback : aceBtCli_bleConnStateChangedCallback()
ACEBTCLI state 0 status 0 connHandle 0xb5f067b0 addr 3f
ACEBTCLI GATT Client Connect Success
```

### Get GATT DB

```
>: ble getdb 0xb5f067b0
ACEBTCLI CLI callback : aceBtCli_bleGattcGetDbCallback()
ACEBTCLI connHandle 0xb5f067b0 no_svc 4
ACEBTCLI Gatt Database index :0 0xb5f00808
ACEBTCLI Service 0 uuid 18 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  serviceType 0
ACEBTCLI        Gatt Characteristics 0 uuid 2a 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI Gatt Database index :1 0xb5f0083c
ACEBTCLI Service 0 uuid 18 0f 00 00 00 00 00 00 00 00 00 00 00 00 00 00  serviceType 0
ACEBTCLI        Gatt Characteristics with Notifications 0 uuid 2a 19 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI                Descriptor UUID 29 02 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI Gatt Database index :2 0xb5f00870
ACEBTCLI Service 0 uuid 18 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00  serviceType 0
ACEBTCLI        Gatt Characteristics 0 uuid 2b 2a 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI Gatt Database index :3 0xb5f008a4
ACEBTCLI Service 0 uuid ff 10 00 00 00 00 00 00 00 00 00 00 00 00 00 00  serviceType 0
ACEBTCLI        Gatt Characteristics 0 uuid ff 11 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI                Descriptor 1 UUID 29 02 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI                Descriptor 2 UUID 29 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI        Gatt Characteristics 1 uuid ff 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI                Descriptor 1 UUID 29 02 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI                Descriptor 2 UUID 29 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI Discover Gatt Service Database Failure
```

### Values in hex ASCII

Applies to read and write. Values seem to be ASCII sent in hex. Check
[this table][ascii table]. I don't think UTF8 is supported here. Examples:

* `ON` is `4f4e`
* `OFF` is `4f4646`

### Reading characteristic

Get the UUID from the `getdb` call and join it all in a single string without spaces.
Characteristic `ff 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00` becomes
`ff120000000000000000000000000000` for this call.

```
>: ble readChars ff120000000000000000000000000000
ACEBTCLI CLI callback : aceBtCli_bleGattcReadCharsCallback() status: 0
ACEBTCLI connHandle 0xb5f067b0
ACEBTCLI UUID:: ff 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI 4f
ACEBTCLI 46
ACEBTCLI 46
ACEBTCLI Read Characteristic Success
```

### Writing characteristic

Value comes at the end. It's ASCII in hex, all joined together. E.g. `ON` is `4f4e`. `OFF` is `4f4646`.

Write LED ON:
```
>: ble writeChars 1 ff120000000000000000000000000000 4f4e
ACEBTCLI CLI callback : aceBtCli_bleGattcWriteCharsCallback()
ACEBTCLI connHandle 0xb5f067b0 gatt format 255
ACEBTCLI Write Characteristic Success
```

Write LED OFF:
```
>: ble writeChars 1 ff120000000000000000000000000000 4f4646
ACEBTCLI CLI callback : aceBtCli_bleGattcWriteCharsCallback()
ACEBTCLI connHandle 0xb5f067b0 gatt format 255
ACEBTCLI Write Characteristic Success
```

### Notification

You can subscribe to a given characteristic with NOTIFY support. When you
subscribe, you'll get constant messages in stdout. You can still type to disable
it back, you'll just lose your inputted characters in a sea of text

#### Enable subscription

```
>: ble regNotify ff120000000000000000000000000000 1
ACEBTCLI CLI callback : aceBtCli_bleGattcWriteDescCallback()
ACEBTCLI connHandle 0xb5f067b0 status 0
ACEBTCLI Register Notification Success

ACEBTCLI CLI callback : aceBtCli_bleGattcNotifyCharsCallback()
ACEBTCLI connHandle 0xb5f067b0
ACEBTCLI UUID:: ff 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI 4f
ACEBTCLI 46
ACEBTCLI 46
ACEBTCLI CLI callback : aceBtCli_bleGattcNotifyCharsCallback()
ACEBTCLI connHandle 0xb5f067b0
ACEBTCLI UUID:: ff 12 00 00 00 00 00 00 00 00 00 00 00 00 00 00
ACEBTCLI 4f
ACEBTCLI 46
ACEBTCLI 46
```

#### Disable subscription

```
>: ble regNotify ff120000000000000000000000000000 0
ACEBTCLI CLI callback : aceBtCli_bleGattcWriteDescCallback()
ACEBTCLI connHandle 0xb5f067b0 status 0
ACEBTCLI Register Notification Success
```




[ascii table]: https://condor.depaul.edu/sjost/it236/documents/ascii.htm
