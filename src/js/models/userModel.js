import { Schema, model } from "mongoose";

const userSchema = new Schema(
  {
    email: { type: String, required: true },
    organisation: {
      type: String,
      required: true,
    },
    cognitoId: {
      type: String,
      required: true,
      index: true,
    },
    apiEnabled: {
      type: Boolean,
      required: true,
      default: false,
    },
    name: {
      type: String,
      required: true,
    },
    encryptedApiKey: {
      type: String,
      required: true,
      index: true,
    },
  },
  { timestamps: true }
);
const UserModel = model("User", userSchema);
export default UserModel;
