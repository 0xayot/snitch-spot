import db from "../utils/dbconnector";

export async function handler(event) {
  await db.connect();
  return {
    statusCode: 200,
    headers: { "Content-Type": "text/plain" },
    body: `Hello, World! Your request was received at ${event.requestContext.time}.`,
  };
}
