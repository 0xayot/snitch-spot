import * as sst from "@serverless-stack/resources";

export default class GoStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);
    this.setDefaultFunctionProps({
      runtime: "go1.x",
    });

    // Create a HTTP API
    const api = new sst.Api(this, "Api", {
      routes: {
        "GET /go": "src/go/welcome",
      },
    });

    // Show the endpoint in the output
    this.addOutputs({
      ApiEndpoint: api.url,
    });
  }
}
