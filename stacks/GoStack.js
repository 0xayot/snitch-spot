import * as sst from "@serverless-stack/resources";

export default class GoStack extends sst.Stack {
  constructor(scope, id, props) {
    super(scope, id, props);
    this.setDefaultFunctionProps({
      runtime: "go1.x",
      environment: {
        MONGODB_URI: process.env.MONGODB_URI,
      },
    });

    // Create a HTTP API
    const api = new sst.Api(this, "GoApi", {
      routes: {
        "GET /go": "src/go/welcome/main.go",
        "POST /report": "src/go/functions/report_incident.go",
      },
    });

    // Show the endpoint in the output
    this.addOutputs({
      ApiEndpoint: api.url,
    });
  }
}
