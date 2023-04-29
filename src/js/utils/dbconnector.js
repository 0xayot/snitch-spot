import mongoose from "mongoose";

const db = {
  connect: () => {
    mongoose
      .connect(process.env.MONGODB_URI, { autoIndex: true })
      .then(() => console.info("DB connection successful!"))
      .catch((e) => {
        throw new Error("Mongo connection Error", e);
      });
    return mongoose;
  },
  close: () => {
    mongoose.connection.close();
  },
};
process.on("SIGINT", () => {
  mongoose.connection.close(() => {
    console.log(
      "Mongoose default connection is disconnected due to application termination"
    );
    process.exit(0);
  });
});

export default db;
