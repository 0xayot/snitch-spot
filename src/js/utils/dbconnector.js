import mongoose from "mongoose";

const db = {
  connect: () => {
    mongoose
      .connect(process.env.MONGODB_URI, {
        autoIndex: true,
        useNewUrlParser: true,
        useUnifiedTopology: true,
      })
      .then(() => console.info("DB connection successful!"))
      .catch((e) => {
        throw new Error("Mongo connection Error", e);
      });
    return mongoose.connection;
  },
  close: () => {
    mongoose.disconnect();
  },
};

export default db;
