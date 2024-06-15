import { Schema, model } from "mongoose";

const apiAccessRequestSchema = new Schema(
  {
    userId: {
      type: Schema.Types.ObjectId,
      ref: "User",
      index: true,
    },
    granted: {
      type: Boolean,
      required: true,
      default: false,
    },
    processed: {
      type: Boolean,
      required: true,
      default: false,
    },
  },
  { timestamps: true }
);

const ApiAccessRequestModel = model("ApiAccessRequest", apiAccessRequestSchema);
export default ApiAccessRequestModel;
