import db from "./dbconnector";

export function success(body) {
  return buildResponse(200, { status: "success", ...body });
}

export function failure(body, status = 500) {
  return buildResponse(status, body);
}

export function verboseResponse(body, status = 500) {
  return buildResponse(status, body);
}

function buildResponse(statusCode, body) {
  console.log(`here's the response body:`, body);

  db.close();

  return {
    statusCode: statusCode,
    headers: {
      "Access-Control-Allow-Origin": "*",
      "Access-Control-Allow-Credentials": true,
      "Access-Control-Expose-Headers": "X-Total-Count",
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
    isBase64Encoded: false,
  };
}
