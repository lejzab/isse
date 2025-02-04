## End points

### Pobranie listy urządzeń
request GET /ers/config/networkdevice/

response 

```json
{
  "SearchResult": {
    "total": 1,
    "resources": [
      {
        "id": "79aac430-7cc8-11eb-ad58-005056926583",
        "name": "ISE_EST_Local_Host",
        "description": "example nd",
        "link": {
          "rel": "self",
          "href": "https://<ise-ip>:9060/ers/config/ers/config/networkdevice/79aac430-7cc8-11eb-ad58-005056926583",
          "type": "application/json"
        }
      }
    ]
  }
}
```

### Wyszukanie urządzenia

### Pobranie urządzenia
request GET /ers/config/networkdevice/{device_id}

response
```json
{
  "NetworkDevice": {
    "id": "device_id",
    "name": "ISE_EST_Local_Host",
    "description": "example nd",
    "authenticationSettings": {
      "networkProtocol": "RADIUS",
      "radiusSharedSecret": "VHEGKOCCUHYB",
      "enableKeyWrap": true,
      "dtlsRequired": true,
      "enableMultiSecret": true,
      "keyEncryptionKey": 1234567890123456,
      "messageAuthenticatorCodeKey": 12345678901234567000,
      "keyInputFormat": "ASCII"
    },
    "snmpsettings": {
      "version": "ONE",
      "roCommunity": "aaa",
      "pollingInterval": 3600,
      "linkTrapQuery": true,
      "macTrapQuery": true,
      "originatingPolicyServicesNode": "Auto"
    },
    "trustsecsettings": {
      "deviceAuthenticationSettings": {
        "sgaDeviceId": "networkDevice1",
        "sgaDevicePassword": "aaa"
      },
      "sgaNotificationAndUpdates": {
        "downlaodEnvironmentDataEveryXSeconds": 86400,
        "downlaodPeerAuthorizationPolicyEveryXSeconds": 86400,
        "reAuthenticationEveryXSeconds": 86400,
        "downloadSGACLListsEveryXSeconds": 86400,
        "otherSGADevicesToTrustThisDevice": false,
        "sendConfigurationToDevice": false,
        "sendConfigurationToDeviceUsing": "ENABLE_USING_COA",
        "coaSourceHost": "IseNodeName"
      },
      "deviceConfigurationDeployment": {
        "includeWhenDeployingSGTUpdates": true,
        "enableModePassword": "aaa",
        "execModePassword": "aaa",
        "execModeUsername": "aaa"
      },
      "pushIdSupport": false
    },
    "tacacsSettings": {
      "sharedSecret": "aaa",
      "connectModeOptions": "ON_LEGACY"
    },
    "profileName": "Cisco",
    "coaPort": 0,
    "dtlsDnsName": "ISE213.il.com",
    "NetworkDeviceIPList": [
      {
        "ipaddress": "127.0.0.1",
        "mask": 32
      }
    ],
    "NetworkDeviceGroupList": [
      "Location#All Locations",
      "Device Type#All Device Types"
    ],
    "link": {
      "rel": "self",
      "href": "https://<ise-ip>:9060/ers/config/ers/config/networkdevice/79aac430-7cc8-11eb-ad58-005056926583",
      "type": "application/json"
    }
  }
}
```