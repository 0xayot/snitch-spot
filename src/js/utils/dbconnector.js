import mongoose from "mongoose";

const db = {
  connect: async () => {
    try {
      await mongoose.connect(process.env.MONGODB_URI, {
        autoIndex: true,
        useNewUrlParser: true,
        useUnifiedTopology: true,
      });
      console.log("DB connection successful!");
    } catch (error) {
      throw new Error("Mongo connection Error", error);
    }
  },
  close: () => {
    mongoose.disconnect();
  },
};

export default db;
