import UserModel from "../models/userModel";
import db from "../utils/dbconnector";
import { retrieveUserId } from "../utils/userUtils";

export async function handler(event) {
  await db.connect();
  const body = event.body;

  // Sanitize body
  const userData = {
    email: body.email,
    organisation: body.organisation,
    name: body.name,
    cognitoId: retrieveUserId(),
  };

  const savedUser = UserModel.create(userData);

  // request api access

  return {
    statusCode: 200,
    headers: { "Content-Type": "text/plain" },
    body: savedUser,
  };
}
