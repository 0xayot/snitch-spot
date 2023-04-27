import * as sst from "@serverless-stack/resources";

export default class GoStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);

    // Create a HTTP API
    const api = new sst.Api(this, "Api", {
      routes: {
        "GET /": "src",
      }
    });

    this.setDefaultFunctionProps({
      runtime: "go1.x"
    });

    // Show the endpoint in the output
    this.addOutputs({
      "ApiEndpoint": api.url,
    });
  }
}
