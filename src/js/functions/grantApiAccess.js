import ApiAccessRequestModel from "../models/apiAccessRequestModel";
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

    await apiAccessRequest.update(updateData);
    const user = await UserModel.findByIdAndUpdate(
      apiAccessRequest.userId,
      { apiEnabled: body.granted },
      { new: true }
    );

    body.granted
      ? sendEmailNotification(user.email, "You can now request API Keys")
      : null;

    return success({ message: "successful" });
  } catch (error) {
    return verboseResponse({ status: false, message: error.message });
  }
}
