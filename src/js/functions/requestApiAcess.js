import ApiAccessRequestModel from "../models/apiAccessRequestModel";
import db from "../utils/dbconnector";
import { success, verboseResponse } from "../utils/responseUtils";
import { retrieveUserRecordByEvent } from "../utils/userUtils";

export async function handler(event) {
  try {
    await db.connect();
    const user = await retrieveUserRecordByEvent(event);

    if (user.apiEnabled)
      return success({ message: "You currently have api access " });

    await ApiAccessRequestModel.create({ userId: user.id });

    return success({ message: "We have received your request" });
  } catch (error) {
    return verboseResponse({ status: false, message: error.message });
  }
}
