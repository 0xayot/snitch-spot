import { Schema, model } from "mongoose";

const userSchema = new Schema(
  {
    email: { type: String, unique: true, required: true },
    organisation: {
      type: String,
      required: true,
    },
    cognitoId: {
      type: String,
      unique: true,
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
      required: false,
      index: true,
    },
  },
  { timestamps: true }
);
const UserModel = model("User", userSchema);
export default UserModel;
