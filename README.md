# Golang-RSA-Playground
- Encrypt & Decrypt message by use public key & private key RSA-4096 (*.asc)  
- Encode & Decode encrypted message by use Base64

### Result-Test

	In practice, someone want to send a message to Tam will use Tam's public key to encrypt the message.
	Then, Bring encrypted message to Tam by use Base64 to encode the message. 
	After that, Tam will decode the message by use Base64 and decrypt the message by use Tam's private key.
	Finaly, Tam will get the message.
	
```
Message : Test-RSA Decryption with Tam's Private Key

-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDbHQDPnqbbiJM3NTkeG+4LeFOEVb6wr8magesbCMbP15tJuSFw....
-----END RSA PRIVATE KEY-----

-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBANsdAM+eptuIkzc1OR4b7gt4U4RVvrCvyZqB6xsIxs/Xm0m5IXCp2Eks....
-----END RSA PUBLIC KEY-----

-----BEGIN MESSAGE-----
ZuTnpW5rA9qPdEmmDGE35V+g3iMUpDDo5FTeyvlGGhRmxUG6CnIDTIun9KhbLqJD....
-----END MESSAGE-----


-------------Begin With Base64------------------

Encode message from RSA By Base64 :
LS0tLS1CRUdJTiBNRVNTQUdFLS0tLS0KW....

Decode message from RSA By Base64 :
-----BEGIN MESSAGE-----
ZuTnpW5rA9qPdEmmDGE35V+g3iMUpDDo5FTeyvlGGhRmxUG6CnIDTIun9KhbLqJD....
...
-----END MESSAGE-----

-------------End With Base64------------------

RSA decrypted : Test-RSA Decryption with Tam's Private Key
```

### Conference
- [Dataset (COVID-19 in Thailand 2021)](https://data.go.th/en/dataset/covid-19-daily)