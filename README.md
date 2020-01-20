# IPFS to JDCloud Edge Storage Client

A command line client (cli) to migrate an IPFS CID to an object in JDCloud's object storage (oss)

## Configure JDCloud

- Step1: register JDCloud account.
- Step2: generate access key (https://uc.jdcloud.com/account/accesskey).
- Step3: enable OSS Service & create a bucket (https://oss-console.jdcloud.com/space)
- Step4: upload an object to bucket.

## Build ipfs-to-oss
```
go build -mod vendor .
```

## Configure CLI
View cli help text for details on flags and environment variables
```
ipfs-to-oss -help
```

### Example
Copy cid `QmYVo8wfic7TvDUL9fcdM8bze5KKrL4QMdKn5LE2jwzMp9`to oss bucket at path `guides/examples/pinning`

```
ipfs-to-oss -cid=QmYVo8wfic7TvDUL9fcdM8bze5KKrL4QMdKn5LE2jwzMp9 -output=guides/examples/pinning
```

## Check result

Visit page with link `https://oss-console.jdcloud.com/edgeStorage?tab=files` to check task result.

## Credit

Thanks to small_fish__ for initial code example
