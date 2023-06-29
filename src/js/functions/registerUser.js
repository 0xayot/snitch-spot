import UserModel from "../models/userModel";
import db from "../utils/dbconnector";
import { success, verboseResponse } from "../utils/responseUtils";
import { retrieveUserId } from "../utils/userUtils";

export async function handler(event) {
  try {
    await db.connect();
    const body = JSON.parse(event.body);

    // Sanitize body
    const userData = {
      email: body.email,
      organisation: body.organisation,
      name: body.name,
      cognitoId: retrieveUserId(event),
    };

    const savedUser = await UserModel.create(userData);

    // request api access

    return success(savedUser.toJSON());
  } catch (error) {
    return verboseResponse({ status: false, message: error.message }, 404);
  }
}
