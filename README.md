# Lighthouse (WIP)

This is a simple project showcasiing reporting fraudulent actors in a "GDPR compliant manner".

THe idea is that a fintech or bank cannot report fraudulent actions to a central and non govt body without breaching data protection laws and exposing thier loss rate to competitors.

However, vetted organisations can encrypt the details of a fraudulent transaction and the offender in a uniform manner to a central directory managed and funded as a public good.

this project features two report inclined end points.
`/unsafe_report` : this is used to report fraud transaction details in plain text.
Authenticated: x-api-key

```

{
  "bankAccountName": "John Doe",
  "bankAccountNumber": "1234567890",
  "bankName": "Bank XYZ",
  "deviceId": "device_id",
  "email": "john.doe@example.com",
  "name": "John Doe",
  "offence": "Fraud",
  "metaData": "additional metadata",
  "resolved": "refunded",
  "showVictim": false,
  "amount": 1000.50,
  "incidentDate": "2023-07-05T10:00:00Z",
}

```

Response:
Errors:
Success:

```
{
    "message": "Succcess Snitching"
}
```

`/unsafe_lookup` : this is used to ascertain if a given piece of information has any relations to fraud
Authenticated: x-api-key

```

{
  "bankAccountName": "John Doe",
  "bankAccountNumber": "1234567890",
  "bankName": "Bank XYZ",
  "deviceId": "device_id",
  "email": "john.doe@example.com",
  "name": "John Doe",
}

Response:
200 {
    "message": "This user has been reported for fraud",
    "isSuspiciousReport": true
}

```

THe project leverages AWS lambdas managed and orchestrated by SST.
This project was bootstrapped with [Create Serverless Stack](https://docs.serverless-stack.com/packages/create-serverless-stack).

Start by installing the dependencies.

```bash
$ npm install
```

## Commands

### `npm run start`

Starts the local Lambda development environment.

### `npm run build`

Build your app and synthesize your stacks.

Generates a `.build/` directory with the compiled files and a `.build/cdk.out/` directory with the synthesized CloudFormation stacks.

### `npm run deploy [stack]`

Deploy all your stacks to AWS. Or optionally deploy, a specific stack.

### `npm run remove [stack]`

Remove all your stacks and all of their resources from AWS. Or optionally removes, a specific stack.

### `npm run test`

Runs your tests using Jest. Takes all the [Jest CLI options](https://jestjs.io/docs/en/cli).

## Documentation

Learn more about the Serverless Stack.

- [Docs](https://docs.serverless-stack.com)
- [@serverless-stack/cli](https://docs.serverless-stack.com/packages/cli)
- [@serverless-stack/resources](https://docs.serverless-stack.com/packages/resources)

## Community

[Follow us on Twitter](https://twitter.com/ServerlessStack) or [post on our forums](https://discourse.serverless-stack.com).
