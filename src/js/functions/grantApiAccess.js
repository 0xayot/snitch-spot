import ApiAccessRequestModel from "../models/apiAccessRequestModel";
// eslint-disable-next-line @typescript-eslint/no-unused-vars
import UserModel from "../models/userModel";
import db from "../utils/dbconnector";
import { failure, success, verboseResponse } from "../utils/responseUtils";
import { sendEmailNotification } from "../utils/utils";

export async function handler(event) {
  try {
    await db.connect();
    const body = JSON.parse(event.body);

    const updateData = { granted: body.granted, processed: true };

    const apiAccessRequest = await ApiAccessRequestModel.findById(
      body.id
    ).populate("userId");

    if (!apiAccessRequest)
      return failure(
        {
          status: "failed",
          message: "Not found",
        },
        400
      );

    await apiAccessRequest.updateOne(updateData);
    await apiAccessRequest.userId.updateOne(
      {
        apiEnabled: updateData.granted,
      },
      { upsert: true }
    );

    const user = apiAccessRequest.userId;

    body.granted
      ? sendEmailNotification(user.email, "You can now request API Keys")
      : null;

    return success({ message: "successful" });
  } catch (error) {
    console.log("GrantApiAccessError", error);
    return verboseResponse({ status: false, message: error.message });
  }
}
