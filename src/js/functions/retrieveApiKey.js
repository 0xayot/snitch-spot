import db from "../utils/dbconnector";
import { failure, success } from "../utils/responseUtils";
import { retrieveUserRecordByEvent } from "../utils/userUtils";
import { generateHash, generateRandomString } from "../utils/utils";

export async function handler(event) {
  try {
    await db.connect();

    const user = await retrieveUserRecordByEvent(event);

    if (user.apiEnabled) {
      const apiKey = generateRandomString();
      user.encryptedApiKey = generateHash(apiKey);

      await user.save();

      return success({ message: "success", apiKey: apiKey });
    } else {
      throw new Error("You are not api Enabled");
    }
  } catch (error) {
    return failure(error.message);
  }
}
