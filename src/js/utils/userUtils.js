import UserModel from "../models/userModel";

export function retrieveUserId(event, fail = true) {
  if (process.env.IS_LOCAL) return event?.requestContext.apiId;
  const userId =
    event?.requestContext?.authorizer?.iam?.cognitoIdentity?.identityId;
  if (userId) return userId;
  if (fail && !userId) throw new Error("User not logged in");
}

export async function retrieveUserRecordByEvent(event, fail = true) {
  let cognitoId;
  if (process.env.IS_LOCAL) {
    cognitoId =
      event?.requestContext?.authorizer?.iam?.cognitoIdentity?.identityId ||
      event?.requestContext.apiId;
  } else {
    cognitoId =
      event?.requestContext?.authorizer?.iam?.cognitoIdentity?.identityId;
  }
  if (fail && !cognitoId) throw new Error("User not logged in");
  const user = await UserModel.findOne({
    cognitoId,
  });
  if (user) {
    return user;
  } else if (fail) {
    throw new Error("User not found");
  }
}
