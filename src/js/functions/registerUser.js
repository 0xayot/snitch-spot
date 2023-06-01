import UserModel from "../models/userModel";
import db from "../utils/dbconnector";

export async function handler(event) {
  await db.connect();
  const body = event.body;

  // Sanitize body
  const userData = {
    email: body.email,
    organisation: body.organisation,
    name: body.name,
  };

  const savedUser = UserModel.create(userData);

  // request api access

  return {
    statusCode: 200,
    headers: { "Content-Type": "text/plain" },
    body: savedUser,
  };
}
