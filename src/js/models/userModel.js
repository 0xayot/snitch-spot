import { Schema, model } from "mongoose";

const userSchema = new Schema(
  {
    organisation: {
      type: String,
      required: true,
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
    },
  },
  { timestamps: true }
);

export default model("User", userSchema);
