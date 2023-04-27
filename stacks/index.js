import GoStack from "./GoStack";
import JsStack from "./JsStack";

export default function main(app) {
  new GoStack(app, "go-stack");

  // Add more stacks
  new JsStack(app, "js-stack");
}
