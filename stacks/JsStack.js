import * as sst from "@serverless-stack/resources";

export default class JsStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);
    this.setDefaultFunctionProps({
      runtime: "nodejs14.x",
    });
    // Create a HTTP API
    const api = new sst.Api(this, "Api", {
      routes: {
        "GET /js": "src/js/welcome/lambda.handler",
      },
    });

    // Show the endpoint in the output
    this.addOutputs({
      ApiEndpoint: api.url,
    });
  }
}
