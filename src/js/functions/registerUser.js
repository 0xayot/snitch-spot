import UserModel from "../models/userModel";
import db from "../utils/dbconnector";
import { success, verboseResponse } from "../utils/responsUtils";
import { retrieveUserId } from "../utils/userUtils";

export async function handler(event) {
  try {
    await db.connect();
    const body = event.body;

    // Sanitize body
    const userData = {
      email: body.email,
      organisation: body.organisation,
      name: body.name,
      cognitoId: retrieveUserId(),
    };

    const savedUser = await UserModel.create(userData);

    // request api access

    return success(savedUser);
  } catch (error) {
    return verboseResponse({ status: false, message: error.message }, 404);
  }
}
