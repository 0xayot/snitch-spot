import { createHash, randomBytes } from "crypto";

export function generateHash(inputString) {
  const hash = createHash("sha512");
  hash.update(inputString);
  return hash.digest("hex");
}

export function generateRandomString() {
  const randomByte = randomBytes(16); // Generate 12 random bytes

  // Convert the random bytes to a hexadecimal string
  const randomString = randomByte.toString("hex");

  return randomString;
}

// TODO
export function sendEmailNotification(emailAddress, message) {
  console.log(emailAddress, message);
  return true;
}
