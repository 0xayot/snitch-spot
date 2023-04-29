import GoStack from "./GoStack";
import JsStack from "./JsStack";

export default function main(app) {
  // API Entry Point
  new GoStack(app, "go-stack", {
    runtime: "go1.x",
  });
  // UI API
  new JsStack(app, "js-stack", {
    runtime: "nodejs14.x",
  });
}
