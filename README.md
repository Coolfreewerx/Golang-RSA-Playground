# Golang-RSA-Playground
- Encrypt & Decrypt message by use RSA-4096 
- Encode & Decode encrypted message by use Base64

### Result-Test
```
Message : Test-RSA Decryption with Tam's Private Key

-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDbHQDPnqbbiJM3NTkeG+4LeFOEVb6wr8magesbCMbP15tJuSFw
qdhJLHzq1/ZINovneoAU6hzolEFoJSKHqmLGMgKn2Aoqo1tuYuLVobjFeKIv9Gfy
PUX4X9rKWMSKcP2wgu8EWGQJM3C5TWaR6fuqhud+DmpEfolIVflM+LK56QIDAQAB
AoGBALz5FJv7jpjD/4ObyXkoSXNnAVCeyumDknJJHkWgmibjIrDAlBCgy0LAtbOt
MmExehPX2OMwzmHgi1CQxmkPEEOaMoIo3J2exoyscyYRFZRNdBioMmJhp8Z5+Ww/
soj1tQe9DU//z1p2AEyOok1fhqVbpYZ7gP5lZ4DIPfIQaU2pAkEA9zs+OoTT+Qlb
cEoS+48vGNnVmqOBs9ZJgX/qAtTOi3dNU4FZAMaTez4Ohj+3+6DgGmCi2tCFQuIP
hwS9EftrkwJBAOLidZo6qKNBF4t3qJZB6XsAWgFOodnIotu126pgILvbNDU65rJt
vSXo+W2PYqsvtsQN5goyaLSHevjmzQ2SChMCQQDMQZzi+hSq/okWF4zhWuWdUXOB
pC6nZpYqMIUku669WN6A7C+dTXJRcu7LCV+2u64K/OXvhDFlFaGgelC+x4qTAkBI
eNnj5/TSIQqqTBCX3nn8BkK5xCpC7KnI0LBYHdiW6RAmKKhkOlV+9IooCvMTh2wz
99SRUUTM1bHcK604NnvtAkAFa8SGiqaBjiZsddTyOWfCioicymkuRfbH/3xuG7o3
QLFltXb1s9y7VPpGpyGuDlq2hXaGnd/9LUUyK9Dt7gsg
-----END RSA PRIVATE KEY-----

-----BEGIN RSA PUBLIC KEY-----
MIGJAoGBANsdAM+eptuIkzc1OR4b7gt4U4RVvrCvyZqB6xsIxs/Xm0m5IXCp2Eks
fOrX9kg2i+d6gBTqHOiUQWglIoeqYsYyAqfYCiqjW25i4tWhuMV4oi/0Z/I9Rfhf
2spYxIpw/bCC7wRYZAkzcLlNZpHp+6qG534OakR+iUhV+Uz4srnpAgMBAAE=
-----END RSA PUBLIC KEY-----

-----BEGIN MESSAGE-----
ZuTnpW5rA9qPdEmmDGE35V+g3iMUpDDo5FTeyvlGGhRmxUG6CnIDTIun9KhbLqJD
0gyfQJhJ6gKq0f3RdLn+0+/GCauP5EJAanBhAt9INiDcJbJwx2nhbBi3qllH/JtA
yRa0tEfP8CEAg64ASnkVw/AFDAxOCsujWGsZf/Q2V6E=
-----END MESSAGE-----


-------------Begin With Base64------------------

Encode message from RSA By Base64 :
LS0tLS1CRUdJTiBNRVNTQUdFLS0tLS0KWnVUbnBXNXJBOXFQZEVtbURHRTM1VitnM2lNVXBERG81RlRleXZsR0doUm14VUc2Q25JRFRJdW45S2hiTHFKRAowZ3lmUUpoSjZnS3EwZjNSZExuKzArL0dDYXVQNUVKQWFuQmhBdDlJTmlEY0piSnd4Mm5oYkJpM3FsbEgvSnRBCnlSYTB0RWZQOENFQWc2NEFTbmtWdy9BRkRBeE9Dc3VqV0dzWmYvUTJWNkU9Ci0tLS0tRU5EIE1FU1NBR0UtLS0tLQo=

Decode message from RSA By Base64 :
-----BEGIN MESSAGE-----
ZuTnpW5rA9qPdEmmDGE35V+g3iMUpDDo5FTeyvlGGhRmxUG6CnIDTIun9KhbLqJD
0gyfQJhJ6gKq0f3RdLn+0+/GCauP5EJAanBhAt9INiDcJbJwx2nhbBi3qllH/JtA
yRa0tEfP8CEAg64ASnkVw/AFDAxOCsujWGsZf/Q2V6E=
-----END MESSAGE-----

-------------End With Base64------------------

RSA decrypted : Test-RSA Decryption with Tam's Private Key
```

### Conference
- [Dataset (COVID-19 in Thailand 2021)](https://data.go.th/en/dataset/covid-19-daily)