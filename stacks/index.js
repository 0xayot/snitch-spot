import GoStack from "./GoStack";
import JsStack from "./JsStack";

export default function main(app) {
  new GoStack(app, "go-stack", {
    runtime: "go1.x",
  });
  // Add more stacks
  new JsStack(app, "js-stack", {
    runtime: "nodejs14.x",
  });
}
