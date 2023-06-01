import ApiAccessRequestModel from "../models/apiAccessRequestModel";
import db from "../utils/dbconnector";
import { retrieveUserRecordByEvent } from "../utils/userUtils";

export async function handler(event) {
  await db.connect();
  const user = retrieveUserRecordByEvent(event);

  if (user.apiEnabled)
    return {
      statusCode: 200,
      headers: { "Content-Type": "text/plain" },
      body: { message: "success" },
    };

  await ApiAccessRequestModel.create({ userId: user.id });

  return {
    statusCode: 200,
    headers: { "Content-Type": "text/plain" },
    body: { message: "success" },
  };
}
