import { Schema, model } from "mongoose";

const apiKeySchema = new Schema(
  {
    userId: {
      type: Schema.Types.ObjectId,
      ref: "User",
      index: true,
    },
    encryptedApiKey: {
      type: String,
      required: true,
      index: true,
    },
    active: {
      type: Boolean,
      required: true,
      default: false,
    },
    expiresAt: {
      type: false,
      required: true,
      default: false,
    },
  },
  { timestamps: true }
);

const ApiKeyModel = model("ApiKeyModel", apiKeySchema);
export default ApiKeyModel;
