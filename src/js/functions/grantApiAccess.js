import ApiAccessRequestModel from "../models/apiAccessRequestModel";
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

    console.log("pop", apiAccessRequest);

    if (!apiAccessRequest)
      return failure(
        {
          status: "failed",
          message: "Not found",
        },
        400
      );
    // if apiAccessRequest

    await apiAccessRequest.updateOne(updateData);
    await apiAccessRequest.userId.updateOne({
      apiEnabled: !updateData.granted,
    });

    const user = apiAccessRequest.userId;

    body.granted
      ? sendEmailNotification(user.email, "You can now request API Keys")
      : null;

    return success({ message: "successful" });
  } catch (error) {
    return verboseResponse({ status: false, message: error.message });
  }
}
