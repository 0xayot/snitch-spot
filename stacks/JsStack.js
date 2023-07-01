import * as sst from "@serverless-stack/resources";

export default class JsStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);
    this.setDefaultFunctionProps({
      runtime: "nodejs14.x",
      environment: {
        MONGODB_URI: process.env.MONGODB_URI,
      },
    });
    // Create a HTTP API
    const api = new sst.Api(this, "Api", {
      routes: {
        "GET /js": "src/js/welcome/lambda.handler",
        "POST /js/signup": "src/js/functions/registerUser.handler",
        "GET /js/request-key": "src/js/functions/requestApiAcess.handler",
        "POST /js/permit": "src/js/functions/grantApiAccess.handler",
        "GET /js/retrieve": "src/js/functions/retrieveApiKey.handler",
      },
    });

    // Show the endpoint in the output
    this.addOutputs({
      ApiEndpoint: api.url,
    });
  }
}
