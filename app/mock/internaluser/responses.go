package internaluser

const InternalUserListJSON = `
{
 "SearchResult": {
   "total": 3,
   "resources": [
     {
       "id": "exampleID1",
       "name": "name1",
       "description": "description1",
       "link": {
         "rel": "self",
         "href": "https://<ise-ip>:9060/ers/config/ers/config/internaluser/exampleID1",
         "type": "application/json"
       }
     },
     {
       "id": "exampleID2",
       "name": "name2",
       "description": "description2",
       "link": {
         "rel": "self",
         "href": "https://<ise-ip>:9060/ers/config/ers/config/internaluser/exampleID2",
         "type": "application/json"
       }
     },
     {
       "id": " exampleID3",
       "name": "name3",
       "description": "description3",
       "link": {
         "rel": "self",
         "href": "https://<ise-ip>:9060/ers/config/ers/config/internaluser/exampleID",
         "type": "application/json"
       }
     }
   ]
 }
}
`

const InternalUserJSON = `
{
  "InternalUser": {
    "id": "%s",
    "name": "name",
    "description": "description",
    "enabled": true,
    "email": "email@domain.com",
    "accountNameAlias": "user123",
    "password": "*******",
    "firstName": "firstName",
    "lastName": "lastName",
    "changePassword": true,
    "identityGroups": "identityGroups",
    "passwordNeverExpires": false,
    "daysForPasswordExpiration": 60,
    "expiryDateEnabled": true,
    "expiryDate": "2021-05-19",
    "enablePassword": "enablePassword",
    "dateModified": "2021-01-03",
    "dateCreated": "2021-01-01",
    "customAttributes": {
      "key1": "value1",
      "key2": "value2"
    },
    "passwordIDStore": "Internal Users",
    "link": {
      "rel": "self",
      "href": "https://<ise-ip>:9060/ers/config/ers/config/internaluser/b24b09e2-16f8-4e48-ace1-2c262a9449f1",
      "type": "application/json"
    }
  }
}
`
